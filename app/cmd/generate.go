package cmd

import "github.com/jchavannes/go-pgp/pgp"

func CmdGenerate() error {
	keyPair, err := pgp.GenerateKeyPair("Test Key", "test", "test@jasonc.me")
	if err != nil {
		return err
	}
	println(keyPair.PrivateKey)
	println(keyPair.PublicKey)
	return nil
}
