package blind_signature

import (
	"crypto/rand"
	"github.com/nik-gautam/major_project_algos/curve"
	"github.com/nik-gautam/major_project_algos/keys"
	"log"
	"math/big"
)

type signer struct {
	a      *keys.PrivateKey
	k1Dash *big.Int
	k2Dash *big.Int
	R1Dash *keys.PublicKey
	R2Dash *keys.PublicKey
	l1     *big.Int
	l2     *big.Int
}

type signerPublicData struct {
	Q      *keys.PublicKey
	R1Dash *keys.PublicKey
	R2Dash *keys.PublicKey
	l1     *big.Int
	l2     *big.Int
}

var defaultSigner signer

func generateSigner() (*signerPublicData, error) {
	var x, y *big.Int

	curve := curve.GetCurve()

	a, err := keys.GenerateKey()
	if err != nil {
		return nil, err
	}

	n, err := rand.Int(rand.Reader, big.NewInt(1<<32)) // might be wrong
	k1Dash, err := rand.Int(rand.Reader, big.NewInt(1<<32))
	k2Dash, err := rand.Int(rand.Reader, big.NewInt(1<<32))
	l1, err := rand.Int(rand.Reader, big.NewInt(1<<32))
	l2, err := rand.Int(rand.Reader, big.NewInt(1<<32))

	if err != nil {
		log.Fatalf("[signer] generateSigner: err while big.NewInt --> %v", err)
	}

	x, y = curve.ScalarBaseMult(k1Dash.Bytes())
	R1Dash := keys.PublicKey{
		X: x,
		Y: y,
	}

	x, y = curve.ScalarBaseMult(k2Dash.Bytes())
	R2Dash := keys.PublicKey{
		X: x,
		Y: y,
	}

	defaultSigner = signer{a, k1Dash, k2Dash, &R1Dash, &R2Dash, l1, l2}

	// r1Dash := R1Dash.X.Mod(R1Dash.X, n) // R1Dash.X % n
	// r2Dash := R2Dash.X.Mod(R2Dash.X, n) // R12Dash.X % n

	return &signerPublicData{
		Q: &keys.PublicKey{
			Curve: curve,
			X:     a.X,
			Y:     a.Y,
		},
		R1Dash: &R1Dash,
		R2Dash: &R2Dash,
		l1:     l1,
		l2:     l2,
	}, nil
}
