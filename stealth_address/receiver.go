package stealthaddress

import "github.com/nik-gautam/major_project_algos/keys"

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
