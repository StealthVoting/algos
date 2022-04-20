package stealthaddress

import "github.com/nik-gautam/major_project_algos/keys"

type PubSender struct {
	A *keys.PublicKey
	B *keys.PublicKey
}

type sender struct {
	a *keys.PrivateKey
	b *keys.PrivateKey
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

func InitTransaction() {

	sen, err := generateSender(); if err != nil {
		return
	}

	se = *sen

	rec := GetPublicAddress()

	

}