package ecc_blind_sign

import "math/big"

func BigIntMod(val *big.Int) *big.Int {
	var nHex = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141"
	n, success := big.NewInt(0).SetString(nHex, 16)
	if !success {
		panic("Panic at the Disco!")
	}

	return big.NewInt(0).Mod(val, n)
}

func BigIntMul(a *big.Int, b *big.Int) *big.Int {
	return BigIntMod(big.NewInt(0).Mul(a, b))
}

func BigIntDiv(a *big.Int, b *big.Int) *big.Int {
	return BigIntMod(big.NewInt(0).Div(a, b))
}
