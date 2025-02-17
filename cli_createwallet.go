package main

import "fmt"

func (cli *CLI) createWallet(nodeID string) string{
	wallets, _ := NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	fmt.Printf("Your new address: %s\n", address)
	return address

}
