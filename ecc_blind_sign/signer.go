package ecc_blind_sign

import (
	"crypto/rand"
	"github.com/nik-gautam/major_project_algos/curve"
	"github.com/nik-gautam/major_project_algos/keys"
	"math"
	"math/big"
)

type signer struct {
	x *big.Int
	r *big.Int
}

type SignerPub struct {
	Y *keys.PublicKey
	H *keys.PublicKey
}

var defaultSigner signer

func GenerateSigner() *SignerPub {
	curve := curve.GetCurve()

	x, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt32)))
	if err != nil {
		panic("[Signer] Error generating x")
	}

	r, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt32)))
	if err != nil {
		panic("[Signer] Error generating r")
	}

	defaultSigner = signer{
		x: x,
		r: r,
	}

	Yx, Yy := curve.ScalarBaseMult(x.Bytes())
	Y := &keys.PublicKey{
		X: Yx,
		Y: Yy,
	}

	Hx, Hy := curve.ScalarBaseMult(r.Bytes())
	H := &keys.PublicKey{
		X: Hx,
		Y: Hy,
	}

	return &SignerPub{
		Y: Y,
		H: H,
	}
}

func SignMessage(voter *VoterPub) *big.Int {
	z := BigIntAdd(BigIntMul(voter.u2, defaultSigner.x), defaultSigner.r)
	return z
}

func VerifySign(Zdash *keys.PublicKey, K *keys.PublicKey, M *keys.PublicKey, u1 *big.Int, P *keys.PublicKey) bool {

	curve := curve.GetCurve()

	//tempKMx, tempKMy := curve.Add(K.X, K.Y, M.X, M.Y)
	//
	//negTempKMy := BigIntMul(tempKMy, big.NewInt(-1))
	//
	//lhsX, lhsY := curve.Add(Zdash.X, Zdash.Y, tempKMx, negTempKMy)
	//lhs := keys.PublicKey{
	//	X: lhsX,
	//	Y: lhsY,
	//}

	tempLshX, tempLshY := curve.Add(Zdash.X, Zdash.Y, K.X, BigIntMul(K.Y, big.NewInt(-1)))

	temp2LshX, temp2LshY := curve.Add(tempLshX, tempLshY, M.X, BigIntMul(M.Y, big.NewInt(-1)))

	lsh := keys.PublicKey{
		X: temp2LshX,
		Y: temp2LshY,
	}

	rhsX, rhsY := curve.ScalarMult(P.X, P.Y, u1.Bytes())
	rhs := keys.PublicKey{
		X: rhsX,
		Y: rhsY,
	}

	println("LHS:- ", lsh.Hex())
	println("RHS:- ", rhs.Hex())

	if lsh.Hex() == rhs.Hex() {
		return true
	} else {
		return false
	}
}
