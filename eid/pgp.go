package eid

import (
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"bytes"
)

type PGPKeyPair struct {
	PublicKey  string
	PrivateKey string
}

func GenerateKeyPair(fullname string, comment string, email string) (PGPKeyPair, error) {
	var e *openpgp.Entity
	e, err := openpgp.NewEntity(fullname, comment, email, nil)
	if err != nil {
		return PGPKeyPair{}, err
	}

	for _, id := range e.Identities {
		err := id.SelfSignature.SignUserId(id.UserId.Id, e.PrimaryKey, e.PrivateKey, nil)
		if err != nil {
			return PGPKeyPair{}, err
		}
	}

	buf := new(bytes.Buffer)
	w, err := armor.Encode(buf, openpgp.PublicKeyType, nil)
	if err != nil {
		return PGPKeyPair{}, err
	}
	e.Serialize(w)
	w.Close()
	pubKey := buf.String()

	buf = new(bytes.Buffer)
	w, err = armor.Encode(buf, openpgp.PrivateKeyType, nil)
	if err != nil {
		return PGPKeyPair{}, err
	}
	e.SerializePrivate(w, nil)
	w.Close()
	privateKey := buf.String()

	return PGPKeyPair{
		PublicKey: pubKey,
		PrivateKey: privateKey,
	}, nil
}

const TestPrivateKey = `-----BEGIN PGP PRIVATE KEY BLOCK-----

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
Yt0sfezumXBvc6mvLNiI
=vIFa
-----END PGP PRIVATE KEY BLOCK-----`

const TestPublicKey = `-----BEGIN PGP PUBLIC KEY BLOCK-----

xsBNBFkJyfUBCADF07OhaDV97RHGQVR6kUnYwdXY0GFArlRmloYqO2ALEe4zdunM
cNMxXLIqY/A0p23jrClbylo+KDnBpR1Y4m+5g3DwdK0jOgmc5QPr5F/s56vW52LQ
wpYHdaWlchCNhJIpCPnbcdf3Og7pmD7M6UBC03JbHepirSgR9OTi5axyKH/EZ1Li
2yLc9Nx/WkZPftW6FPOaWB0f8v2vMhg1kNlZu9Cn4d1e2NYTEl8YleBxDmQElnBe
nUe1xcINoIuk9OdcTfnwR7tufh/NVNn3PsRdjdP6YjcpdNtpKCwyQ86TogVHNu2j
e7l/sjvUzXA0SS7UN9I521MMo/EYeY4wyhJJABEBAAHNIFRlc3QgS2V5ICh0ZXN0
KSA8dGVzdEBqYXNvbmMubWU+wsBiBBMBCAAWBQJZCcn1CRBLJ+J6kOitRAIbAwIZ
AQAA6QYIAJxNDFKFuiaO4SVX5kQ7J1nsbFOJVz53RlXqNABp2xm1i1+/x7uHcwmI
nOgFJGEdLO82rwYVfQDveYclVGRQBe3q2ui06O28g+/7ZNw3GxfJkzg4Tn4aXINM
1ABgv6+IrE+8iGkOUhzPq4jwIqcLL8SZe9HAhSepOI6vkWgcOYTVkPqFs37/XmY/
kTtMMMKAntIB75johv5qSh0CKh4h+QeD1te76G0Qit82FeXM+LXmd9edSXjV7nyz
vFi2raerhBd6bwbaFBqKaoB84h/idrTwvweyrSBEBB+CvtC9uyQhpLg/eP4G4BTG
F0ZpwvWDIXg9SxqkKnHjdablbtEP2T7OwE0EWQnJ9QEIANwvG0HVPXcPiW1ujBiU
bR4ELgr6s+25UB2tDzUrzP4IXTtCqW+JecinyjfBVgwWHkmDKKoByo9V0TIkqxhU
5JDk+w980bEWRuCD8XRQnvsBo4xJLPRqqLjYd/V4m6mxjQzWdNnpkeJcpWRqSP1p
zTh4jjcuG8NnYzgpSYr9gF5Z1RE7XglowHgJ/SACZWcrQPCZHq7Quq1TIsADMKNb
IfFepOufMFsT17+rOn8jyNPTObe7T5rlH8+C164FigqdRRgEBaTUyRVdWzCu8aIr
SRr7DgIF6C8E4D9uzRh/lJclVxNQhM1kyEePhY0E1d6oz7Y2hlKu0l6ra7LeP5g5
+ukAEQEAAQ==
=fAsQ
-----END PGP PUBLIC KEY BLOCK-----`
