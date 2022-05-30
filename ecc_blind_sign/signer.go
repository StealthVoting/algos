package ecc_blind_sign

import (
	"crypto/rand"
	"github.com/nik-gautam/major_project_algos/curve"
	"github.com/nik-gautam/major_project_algos/keys"
	"log"
	"math/big"
)

type signer struct {
	db     *keys.PrivateKey // signer private key
	k1Dash *keys.PrivateKey
	k2Dash *keys.PrivateKey
}

type SignerPublicData struct {
	Q      *keys.PublicKey // signer public key
	R1Dash *keys.PublicKey
	R2Dash *keys.PublicKey
	r1Dash *big.Int
	r2Dash *big.Int
	l1     *big.Int
	l2     *big.Int
}

var defaultSignerPvt signer

func GenerateSigner() (*SignerPublicData, error) {
	var nHex = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141"
	n, success := big.NewInt(0).SetString(nHex, 16)
	if !success {
		panic("Panic at the Disco!")
	}

	//helper := big.NewInt(0)
	db, err := keys.GenerateKey()
	if err != nil {
		return nil, err
	}

	k1Dash, err := keys.GenerateKey()
	k2Dash, err := keys.GenerateKey()

	l1, err := rand.Int(rand.Reader, big.NewInt(1<<16))
	l2, err := rand.Int(rand.Reader, big.NewInt(1<<16))

	R1Dash := k1Dash.PublicKey
	R2Dash := k2Dash.PublicKey

	if err != nil {
		log.Fatalf("[signer] generateSigner: err while big.NewInt --> %v", err)
	}

	defaultSignerPvt = signer{db, k1Dash, k2Dash}

	//fmt.Println(R1Dash.X.String(), R2Dash.X.String())
	r1Dash := big.NewInt(0).Mod(R1Dash.X, n) // R1Dash.X % n
	r2Dash := big.NewInt(0).Mod(R2Dash.X, n) // R2Dash.X % n

	return &SignerPublicData{
		Q:      db.PublicKey,
		R1Dash: R1Dash,
		R2Dash: R2Dash,
		r1Dash: r1Dash,
		r2Dash: r2Dash,
		l1:     l1,
		l2:     l2,
	}, nil
}

func RequestBlindSign(m1Dash *big.Int, m2Dash *big.Int, signerPub *SignerPublicData) (*big.Int, *big.Int) {
	// s1Dash = ((ecPrivKey * m1Dash) % n) - ((r1Dash * k1Dash * l1) % n);
	// s2Dash = ((ecPrivKey * m2Dash) % n) - ((r2Dash * k2Dash * l2) % n);
	//var nHex = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141"
	//n, success := big.NewInt(0).SetString(nHex, 16)
	//if !success {
	//	panic("Panic at the Disco!")
	//}

	signerPvtKey := new(big.Int).SetBytes(defaultSignerPvt.db.Bytes())
	k1Dash := big.NewInt(0).SetBytes(defaultSignerPvt.k1Dash.Bytes())
	k2Dash := big.NewInt(0).SetBytes(defaultSignerPvt.k2Dash.Bytes())

	//helper := big.NewInt(0)

	s1Dash := big.NewInt(0).Sub(
		BigIntMul(signerPvtKey, m1Dash),
		BigIntMul(signerPub.r1Dash, BigIntMul(k1Dash, signerPub.l1)))

	s2Dash := big.NewInt(0).Sub(
		BigIntMul(signerPvtKey, m2Dash),
		BigIntMul(signerPub.r2Dash, BigIntMul(k2Dash, signerPub.l2)))

	return s1Dash, s2Dash
}

//function verifyBlindSignature(
//uint256 s,
//uint256 Rx,
//uint256 Ry,
//uint256 r,
//uint256 m
//)

func VerifyBlindSign(s *big.Int, R *keys.PublicKey, r *big.Int, m *big.Int, signerPub *SignerPublicData) bool {
	// mQ = sG + rR;
	println("s", s.String())

	curveObj := curve.GetCurve()

	mQx, mQy := curveObj.ScalarMult(signerPub.Q.X, signerPub.Q.Y, m.Bytes())

	sGx, sGy := curveObj.ScalarBaseMult(s.Bytes())

	rRx, rRy := curveObj.ScalarMult(R.X, R.Y, r.Bytes())

	sGrRx, sGrRy := curveObj.Add(sGx, sGy, rRx, rRy)

	if mQx == sGrRx && mQy == sGrRy {
		return true
	}

	return false
}
