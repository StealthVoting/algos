package keys

import (
	"bytes"
	"crypto/elliptic"
	"encoding/hex"
	"math/big"
)

type PublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}

func (K *PublicKey) Bytes() []byte {
	x := K.X.Bytes()

	if len(x) < 32 {
		for i := 0; i < 32-len(x); i++ {
			x = append([]byte{0}, x...)
		}
	}

	y := K.Y.Bytes()

	if len(y) < 32 {
		for i := 0; i < 32-len(y); i++ {
			y = append([]byte{0}, y...)
		}
	}

	return bytes.Join([][]byte{{0x04}, x, y}, nil)
}

func (K *PublicKey) Hex() string {
	return hex.EncodeToString(K.Bytes())
}
