package ecc_blind_sign

import (
	"crypto/rand"
	"github.com/nik-gautam/major_project_algos/curve"
	"log"
	"math/big"
)

func BigIntMod(val *big.Int) *big.Int {
	var nHex = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141"
	n, success := big.NewInt(0).SetString(nHex, 16)
	if !success {
		log.Fatalf("Error creating n")
	}

	return big.NewInt(0).Mod(val, n)
}

var P = curve.GetCurve().Params().P

func BigIntAdd(a *big.Int, b *big.Int) *big.Int {
	res := big.NewInt(0)

	return res.Add(res, a).Add(res, b).Mod(res, P)

	//return big.NewInt(0).Add(a, b).Mod()
}

func BigIntMul(a *big.Int, b *big.Int) *big.Int {
	res := big.NewInt(1)

	return res.Mul(res, a).Mul(res, b).Mod(res, P)

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
