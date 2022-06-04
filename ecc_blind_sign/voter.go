package ecc_blind_sign

import (
	"crypto/rand"
	"crypto/sha1"
	"github.com/nik-gautam/major_project_algos/curve"
	"github.com/nik-gautam/major_project_algos/keys"
	"math"
	"math/big"
)

type voter struct {
	a *big.Int
	b *big.Int
	w *big.Int
}

type VoterPub struct {
	u1 *big.Int
	u2 *big.Int
	M  *keys.PublicKey
	K  *keys.PublicKey
	P  *keys.PublicKey // P = aY
	Q  *keys.PublicKey // Q = bY
}

var defaultVoter voter

func GenerateVoter() {
	curve := curve.GetCurve()

	signer := GenerateSigner()

	hasher := sha1.New()

	a, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt32)))
	if err != nil {
		panic("[Voter] Error generating a")
	}

	b, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt32)))
	if err != nil {
		panic("[Voter] Error generating b")
	}

	w, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt32)))
	if err != nil {
		panic("[Voter] Error generating w")
	}

	defaultVoter = voter{
		a: a,
		b: b,
		w: w,
	}

	A := &keys.PublicKey{}
	A.X, A.Y = curve.ScalarBaseMult(a.Bytes())

	B := &keys.PublicKey{}
	B.X, B.Y = curve.ScalarBaseMult(b.Bytes())

	P := &keys.PublicKey{}
	P.X, P.Y = curve.ScalarMult(signer.Y.X, signer.Y.Y, a.Bytes())

	Q := &keys.PublicKey{}
	Q.X, Q.Y = curve.ScalarMult(signer.Y.X, signer.Y.Y, b.Bytes())

	K := &keys.PublicKey{}
	K.X, K.Y = curve.ScalarBaseMult(w.Bytes())

	// Signing Phase starts here
	m := big.NewInt(1021)

	hasher.Write(A.Bytes())
	hasher.Write(B.Bytes())
	hasher.Write(m.Bytes())

	u1 := big.NewInt(0).SetBytes(hasher.Sum(nil))

	u2 := BigIntAdd(u1, b)

	HQ := &keys.PublicKey{}
	HQ.X, HQ.Y = curve.Add(signer.H.X, signer.H.Y, Q.X, Q.Y)

	M := &keys.PublicKey{}
	M.X, M.Y = curve.ScalarMult(HQ.X, HQ.Y, a.Bytes())

	voterPub := &VoterPub{
		u1: u1,
		u2: u2,
		M:  M,
		K:  K,
		P:  P,
		Q:  Q,
	}

	z := SignMessage(voterPub)
	// Signing Phase ends here

	// Extraction Phase starts here
	temp1 := BigIntAdd(BigIntMul(z, a), w)

	//println(temp1.String())

	Zdash := &keys.PublicKey{}
	Zdash.X, Zdash.Y = curve.ScalarBaseMult(temp1.Bytes())
	// Extraction Phase ends here

	// Verification starts here
	isValid := VerifySign(Zdash, K, M, u1, P)

	if isValid {
		println("Valid Sign")
	} else {
		println("InValid Sign")
	}
}
