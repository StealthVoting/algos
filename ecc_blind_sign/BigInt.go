package ecc_blind_sign

import (
	"crypto/rand"
	"github.com/nik-gautam/major_project_algos/curve"
	"math/big"
)

var Level = 32

var P = curve.GetCurve().Params().N

func BigIntMod(val *big.Int) *big.Int {
	return val.Mod(val, P)
}

func BigIntAdd(a *big.Int, b *big.Int) *big.Int {
	res := big.NewInt(0)

	return res.Add(res, a).Add(res, b)

	//return big.NewInt(0).Add(a, b).Mod()
}

func BigIntMul(a *big.Int, b *big.Int) *big.Int {
	res := big.NewInt(1)

	return res.Mul(res, a).Mul(res, b)

	//return big.NewInt(0).Mul(a, b)
}

func BigIntDiv(a *big.Int, b *big.Int) *big.Int {
	return big.NewInt(0).Div(a, b)
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
