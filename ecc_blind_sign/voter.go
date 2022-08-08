package ecc_blind_sign

import (
	"crypto/rand"
	"crypto/sha1"
	"github.com/nik-gautam/major_project_algos/curve"
	"github.com/nik-gautam/major_project_algos/keys"
	"math/big"
)

type voter struct {
	a *big.Int // A = aG
	b *big.Int // B = bG
	w *big.Int
}

type VoterPub struct {
	u1 *big.Int
	u2 *big.Int
	M  *keys.PublicKey // M = a(H + Q) --> signer dependent
	K  *keys.PublicKey // k = wG
	P  *keys.PublicKey // P = aY --> signer dependent
	Q  *keys.PublicKey // Q = bY --> signer dependent
}

var defaultVoter voter

func GenerateVoter() {
	curve := curve.GetCurve()

	signer := GenerateSigner()

	hasher := sha1.New()

	byt, _ := GenerateRandomBytes(16)

	a, err := rand.Int(rand.Reader, new(big.Int).SetBytes(byt))
	if err != nil {
		panic("[Voter] Error generating a")
	}

	println("a " + a.String())

	byt, _ = GenerateRandomBytes(16)

	b, err := rand.Int(rand.Reader, new(big.Int).SetBytes(byt))
	if err != nil {
		panic("[Voter] Error generating b")
	}

	println("b " + b.String())

	byt, _ = GenerateRandomBytes(16)

	w, err := rand.Int(rand.Reader, new(big.Int).SetBytes(byt))
	if err != nil {
		panic("[Voter] Error generating w")
	}

	println("w " + w.String())

	defaultVoter = voter{
		a: a,
		b: b,
		w: w,
	}

	A := &keys.PublicKey{}
	A.X, A.Y = curve.ScalarBaseMult(a.Bytes())

	println("A:- " + A.Hex())

	B := &keys.PublicKey{}
	B.X, B.Y = curve.ScalarBaseMult(b.Bytes())
	println("B:- " + B.Hex())

	P := &keys.PublicKey{}
	P.X, P.Y = curve.ScalarMult(signer.Y.X, signer.Y.Y, a.Bytes())
	println("P:- " + P.Hex())

	Q := &keys.PublicKey{}
	Q.X, Q.Y = curve.ScalarMult(signer.Y.X, signer.Y.Y, b.Bytes())
	println("Q:- " + Q.Hex())

	K := &keys.PublicKey{}
	K.X, K.Y = curve.ScalarBaseMult(w.Bytes())
	println("K:- " + K.Hex())

	// Signing Phase starts here
	m := big.NewInt(1021)

	hasher.Write(A.Bytes())
	hasher.Write(B.Bytes())
	hasher.Write(m.Bytes())

	// u1 = hash(aG || bG || m)

	u1 := big.NewInt(0).SetBytes(hasher.Sum(nil))

	u2 := BigIntAdd(u1, b) // u2 = u1 + b

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
	// Sign = {Zdash, u1, K}
	// Zdash = (z * a + w)G
	temp1 := BigIntAdd(BigIntMul(z, a), w)

	//println(temp1.String())

	Zdash := &keys.PublicKey{}
	Zdash.X, Zdash.Y = curve.ScalarBaseMult(temp1.Bytes())
	// Extraction Phase ends here

	// Verification starts here
	// Zdash - (M + K) = u1*P
	isValid := VerifySign(Zdash, K, M, u1, P)

	if isValid {
		println("Valid Sign")
	} else {
		println("InValid Sign")
	}
}
