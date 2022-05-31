package main

import "github.com/nik-gautam/major_project_algos/ecc_blind_sign"

func main() {

	//txn, err := stealthaddress.InitTransaction()
	//if err != nil {
	//	log.Fatal("Phat gya")
	//
	//	return
	//}
	//
	//stealthaddress.ProcessTxn(txn)

	ecc_blind_sign.GenerateVoter()
}
