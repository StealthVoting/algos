package stealthaddress

import (
	"github.com/nik-gautam/major_project_algos/curve"
	"github.com/nik-gautam/major_project_algos/hash"
	"github.com/nik-gautam/major_project_algos/keys"
)

type PubReceiver struct {
	A *keys.PublicKey
	B *keys.PublicKey
}

type receiver struct {
	a *keys.PrivateKey
	b *keys.PrivateKey
}

var re receiver

func generateReceiver() (*receiver, error) {

	a, err := keys.GenerateKey()
	if err != nil {
		return nil, err
	}

	b, err := keys.GenerateKey()
	if err != nil {
		return nil, err
	}

	return &receiver{
		a: a,
		b: b,
	}, nil
}

func GetPublicAddress() *PubReceiver {

	res, err := generateReceiver()
	if err != nil {
		return nil
	}

	re = *res

	return &PubReceiver{
		A: re.a.PublicKey,
		B: re.b.PublicKey,
	}
}

func ProcessTxn(txn *Txn) {

	curve := curve.GetCurve()

	temp1x, temp1y := curve.ScalarMult(txn.R.X, txn.R.Y, re.a.Bytes())
	temp1 := keys.PublicKey{
		X: temp1x,
		Y: temp1y,
	}

	temp2x, temp2y := curve.ScalarBaseMult(hash.Hash(temp1.Bytes()))
	temp2 := keys.PublicKey{
		X: temp2x,
		Y: temp2y,
	}

	p_dash_x, p_dash_y := curve.Add(temp2.X, temp2.Y, re.b.PublicKey.X, re.b.PublicKey.Y)
	P_dash := keys.PublicKey{
		X: p_dash_x,
		Y: p_dash_y,
	}

	println("P ---->  ", txn.P.Hex())
	println("P' --->  ", P_dash.Hex())

	if txn.P.Hex() == P_dash.Hex() {
		println()
		println("Transaction Successfully Received....")
		println("Message ---> ", txn.M)
	}
}
