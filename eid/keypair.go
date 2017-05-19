package eid

import (
	"golang.org/x/crypto/openpgp"
	"github.com/jchavannes/go-pgp/pgp"
)

type KeyPair struct {
	PublicKey  []byte
	PrivateKey []byte
}

func (p KeyPair) GetPgpEntity() (*openpgp.Entity, error) {
	return pgp.GetEntity(p.PublicKey, p.PrivateKey)
}
