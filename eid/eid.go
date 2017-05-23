package eid

import (
	"encoding/json"
	"github.com/jchavannes/go-pgp/pgp"
	"fmt"
	"regexp"
)

func JsonMarshalSignAndEncrypt(data interface{}, recipientPublicKey KeyPair, senderPrivateKey KeyPair) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return []byte{}, fmt.Errorf("Unable to json marshal data: %s", err)
	}

	senderPrivateEntity, err := senderPrivateKey.GetPgpEntity()
	if err != nil {
		return []byte{}, fmt.Errorf("Unable to create private pgp entity: %s", err)
	}

	signature, err := pgp.Sign(senderPrivateEntity, jsonData)
	if err != nil {
		return []byte{}, fmt.Errorf("Unable to sign message: %s", err)
	}

	recipientPublicEntity, err := recipientPublicKey.GetPgpEntity()
	if err != nil {
		return []byte{}, fmt.Errorf("Unable to create public pgp entity: %s", err)
	}

	message := append(jsonData, signature...)
	encrypted, err := pgp.Encrypt(recipientPublicEntity, message)
	if err != nil {
		return []byte{}, fmt.Errorf("Unable to encrypt data: %s", err)
	}
	return encrypted, nil
}

func DecryptVerifyAndUnmarshal(encrypted []byte, recipientPrivateKey KeyPair, senderPublicKey KeyPair, unmarshal interface{}) error {
	pgpEntity, err := recipientPrivateKey.GetPgpEntity()
	if err != nil {
		return fmt.Errorf("Error getting recipient entity: %s", err)
	}

	decrypted, err := pgp.Decrypt(pgpEntity, encrypted)
	if err != nil {
		return fmt.Errorf("Error decrypting message: %s", err)
	}

	reg := regexp.MustCompile(`-----BEGIN PGP SIGNATURE-----[A-Za-z0-9/=+\s]+-----END PGP SIGNATURE-----$`)
	signature := reg.FindString(string(decrypted))
	jsonString := reg.ReplaceAllString(string(decrypted), "")

	senderPublicEntity, err := senderPublicKey.GetPgpEntity()
	if err != nil {
		return fmt.Errorf("Error getting sender public entity: %s", err)
	}

	err = pgp.Verify(senderPublicEntity, []byte(jsonString), []byte(signature))
	if err != nil {
		return fmt.Errorf("Error verifying message: %s", err)
	}

	return json.Unmarshal([]byte(jsonString), unmarshal)
}
