package curve

import (
	"crypto/elliptic"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func GetCurve() elliptic.Curve {
	return secp256k1.S256()
}
