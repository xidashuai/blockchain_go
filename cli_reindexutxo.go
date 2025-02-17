package main

import "fmt"

func (cli *CLI) reindexUTXO(nodeID string) (int, bool){
	bc := NewBlockchain(nodeID)
	UTXOSet := UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
	return count,true
}
