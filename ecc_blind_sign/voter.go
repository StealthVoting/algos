package ecc_blind_sign

import (
	"crypto/rand"
	"github.com/nik-gautam/major_project_algos/curve"
	"github.com/nik-gautam/major_project_algos/keys"
	"log"
	"math/big"
)

type voter struct {
	a *big.Int
	b *big.Int
	w *big.Int
	z *big.Int
	e *big.Int
	d *big.Int
}

type VoterPublicData struct {
	R1 *keys.PublicKey
	R2 *keys.PublicKey
	r1 *big.Int
	r2 *big.Int
}

var defaultVoterPvt voter
var defaultVoterPub VoterPublicData

func GenerateVoter() {
	var nHex = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141"
	n, success := big.NewInt(0).SetString(nHex, 16)
	if !success {
		panic("Panic at the Disco!")
	}

	a, err := rand.Int(rand.Reader, big.NewInt(1<<16))
	b, err := rand.Int(rand.Reader, big.NewInt(1<<16))
	w, err := rand.Int(rand.Reader, big.NewInt(1<<16))

	// generate coPrime
	z := big.NewInt(0)
	for i := big.NewInt(2); i.Cmp(w) == -1; i = i.Add(i, big.NewInt(1)) {
		if big.NewInt(1).Cmp(big.NewInt(0).GCD(nil, nil, w, i)) == 0 {
			z = i
			break
		}
	}

	// ew + dz = gcd(w, z) = 1
	e := big.NewInt(0)
	d := big.NewInt(0)

	gcd := big.NewInt(0).GCD(e, d, w, z)

	if gcd.Cmp(big.NewInt(1)) != 0 {
		println("gcd", gcd)
		panic("non 1 gcd")
	}

	defaultVoterPvt = voter{a: a, b: b, w: w, z: z, e: e, d: d}

	//helper := big.NewInt(0)
	//-------------------------------------------------------------------------------

	//R1 = R1' * w * a * l1 --> R1 = (xr1, yr1)
	//R2 = R2' * z * b * l2 --> R2 = (xr2, yr2)
	//
	//r1 = xr1 mod n
	//r2 = xr2 mod n

	signer, err := GenerateSigner()

	curveObj := curve.GetCurve()

	R1x, R1y := curveObj.ScalarMult(signer.R1Dash.X, signer.R1Dash.Y, BigIntMul(w, BigIntMul(a, signer.l1)).Bytes())
	R1 := keys.PublicKey{
		X: R1x,
		Y: R1y,
	}

	R2x, R2y := curveObj.ScalarMult(signer.R2Dash.X, signer.R2Dash.Y, BigIntMul(z, BigIntMul(b, signer.l2)).Bytes())
	R2 := keys.PublicKey{
		X: R2x,
		Y: R2y,
	}

	r1 := big.NewInt(0).Mod(R1.X, n)
	r2 := big.NewInt(0).Mod(R2.X, n)

	m := big.NewInt(16477298298399)

	// m1' = (e * m * r1' * 1/r1 * 1/r2 * 1/a) mod n
	//m1Dash := big.NewInt(0).Mod(
	//	big.NewInt(0).Div(
	//		big.NewInt(0).Mul(
	//			big.NewInt(0).Mul(e, m), signer.r1Dash), big.NewInt(0).Mul(r1, big.NewInt(0).Mul(r2, a))), n)
	m1Dash := BigIntMod(BigIntDiv(BigIntMul(e, BigIntMul(m, signer.r1Dash)), BigIntMul(r1, BigIntMul(r2, a))))

	// m2' = (d * m * r2' * 1/r1 * 1/r2 * 1/b) mod n
	//m2Dash := big.NewInt(0).Mod(big.NewInt(0).Div(big.NewInt(0).Mul(big.NewInt(0).Mul(d, m), signer.r2Dash), big.NewInt(0).Mul(r1, big.NewInt(0).Mul(r2, b))), n)
	m2Dash := BigIntMod(BigIntDiv(BigIntMul(d, BigIntMul(m, signer.r2Dash)), BigIntMul(r1, BigIntMul(r2, b))))

	//Signing
	s1Dash, s2Dash := RequestBlindSign(m1Dash, m2Dash, signer)

	// Extraction
	//s1 = (s1' * 1/r1' * r1 * r2 * w * a ) mod n
	//s2 = (s2' * 1/r2' * r1 * r2 * z * b ) mod n

	//s1 := big.NewInt(0).Mod(big.NewInt(0).Div(big.NewInt(0).Mul(s1Dash, big.NewInt(0).Mul(r1, big.NewInt(0).Mul(r2, big.NewInt(0).Mul(w, a)))), signer.r1Dash), n)
	//s2 := big.NewInt(0).Mod(big.NewInt(0).Div(big.NewInt(0).Mul(s2Dash, big.NewInt(0).Mul(r1, big.NewInt(0).Mul(r2, big.NewInt(0).Mul(z, b)))), signer.r2Dash), n)

	s1 := BigIntMod(BigIntDiv(BigIntMul(s1Dash, BigIntMul(r1, BigIntMul(r2, BigIntMul(w, a)))), signer.r1Dash))
	s2 := BigIntMod(BigIntDiv(BigIntMul(s2Dash, BigIntMul(r1, BigIntMul(r2, BigIntMul(z, b)))), signer.r2Dash))

	s := big.NewInt(0).Mod(big.NewInt(0).Add(s1, s2), n)
	Rx, Ry := curveObj.Add(R1.X, R1.Y, R2.X, R2.Y)
	R := &keys.PublicKey{
		X: Rx,
		Y: Ry,
	}
	r := BigIntMul(r1, r2)

	// Verification
	chalGya := VerifyBlindSign(s, R, r, m, signer)

	println("chalGya", chalGya)

	if chalGya {
		println("yayayayayyayayayayayyayayayay")
	} else {
		println("I give up")
	}

	if err != nil {
		log.Fatalf("[voter] generateVoter: err while big.NewInt --> %v", err)
	}
}
