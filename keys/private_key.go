package keys

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/nik-gautam/major_project_algos/curve"
)

type PrivateKey struct {
	*PublicKey
	D *big.Int
}

func GenerateKey() (*PrivateKey, error) {

	curve := curve.GetCurve()

	p, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("cannot generate keys: %w", err)
	}

	d := big.NewInt(0).SetBytes(p)

	return &PrivateKey{
		&PublicKey{
			Curve: curve,
			X:     x,
			Y:     y,
		},
		d,
	}, nil

}

func (k *PrivateKey) Bytes() []byte {
	return k.D.Bytes()
}

func (k *PrivateKey) Hex() string {
	return hex.EncodeToString(k.Bytes())
}
