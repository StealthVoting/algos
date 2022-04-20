package stealthaddress

import (
	"crypto/rand"
	"math/big"

	"github.com/nik-gautam/major_project_algos/hash"

	"github.com/nik-gautam/major_project_algos/curve"
	"github.com/nik-gautam/major_project_algos/keys"

	"fmt"
)

type PubSender struct {
	A *keys.PublicKey
	B *keys.PublicKey
}

type sender struct {
	a *keys.PrivateKey
	b *keys.PrivateKey
}

type Txn struct {
	P *keys.PublicKey
	R *keys.PublicKey
	M string
}

var se sender

func generateSender() (*sender, error) {

	a, err := keys.GenerateKey()
	if err != nil {
		return nil, err
	}

	b, err := keys.GenerateKey()
	if err != nil {
		return nil, err
	}

	return &sender{
		a: a,
		b: b,
	}, nil
}

func InitTransaction() (*Txn, error) {
	sen, err := generateSender()
	if err != nil {
		return nil, fmt.Errorf("cannot generate txn: %w", err)
	}

	se = *sen

	rec := GetPublicAddress() // receiver public address

	curve := curve.GetCurve()

	r, err := rand.Int(rand.Reader, big.NewInt(1<<32)) // sender private key
	if err != nil {
		return nil, fmt.Errorf("cannot generate txn: %w", err)
	}

	x, y := curve.ScalarBaseMult(r.Bytes())

	R := keys.PublicKey{ // sender public key
		X: x,
		Y: y,
	}

	temp1x, temp1y := curve.ScalarMult(rec.A.X, rec.A.Y, r.Bytes())
	temp1 := keys.PublicKey{
		X: temp1x,
		Y: temp1y,
	}

	temp2x, temp2y := curve.ScalarBaseMult(hash.Hash(temp1.Bytes()))
	temp2 := keys.PublicKey{
		X: temp2x,
		Y: temp2y,
	}

	px, py := curve.Add(temp2.X, temp2.Y, rec.B.X, rec.B.Y)

	P := keys.PublicKey{
		X: px,
		Y: py,
	}

	// println("R ---> ", R.Hex())
	// println("P ---> ", P.Hex())

	return &Txn{
		P: &P,
		R: &R,
		M: "Hello Receiver",
	}, nil
}
