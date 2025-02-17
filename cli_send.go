package main

import (
	"fmt"
	"log"
)

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) bool {
	if !ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
		return false
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
		return false
	}

	bc := NewBlockchain(nodeID)
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()

	wallets, err := NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
		return false
	}
	wallet := wallets.GetWallet(from)

	tx := NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := NewCoinbaseTX(from, "")
		txs := []*Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		fmt.Println(knownNodes[0])
		sendTx(knownNodes[0], tx)
	}

	fmt.Println("Success!")
	return true
}
