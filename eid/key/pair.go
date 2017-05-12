package key

import (
	"golang.org/x/crypto/openpgp"
	"github.com/jchavannes/go-pgp/pgp"
)

type Pair struct {
	PublicKey  []byte
	PrivateKey []byte
}

func (p Pair) GetPgpEntity() (*openpgp.Entity, error) {
	return pgp.GetEntity(p.PublicKey, p.PrivateKey)
}
