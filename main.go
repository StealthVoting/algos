package main

import (
	"log"

	stealthaddress "github.com/nik-gautam/major_project_algos/stealth_address"
)

func main() {

	txn, err := stealthaddress.InitTransaction()
	if err != nil {
		log.Fatal("Phat gya")

		return
	}

	stealthaddress.ProcessTxn(txn)

}
