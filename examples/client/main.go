package main

import (
	"github.com/k0kubun/pp"
	"github.com/xboston/metahash-go"
)

//sample address
const address = "0x0062b734726a66f54e9ac72f63af9fd18bd5ca5583061e2c37"         //serengeti node
const trx = "361d5e113648366aa46010fa02ff1f66dcf8045abe4e30547ffaba0506edbdb2" //sample transaction
func main() {
	//fetch balance
	/*balance, err := metahash.FetchBalance(address)
	if err == nil {
		pp.Println("fetch-balance", balance)
	} else {
		pp.Println("error fetching balance")
	}*/

	//fetch history
	/*history, err := metahash.FetchHistory(address)
	if err == nil {
		pp.Println("fetch-history", history[0]) //TODO: check ifthe history is not nil
	} else {
		pp.Println("error fetching history")
	}*/

	//fetch history
	/*historyRange, err := metahash.FetchHistoryRange(address, 1, 2)
	if err == nil {
		pp.Println("fetch-history", historyRange) //TODO: check ifthe history is not nil
	} else {
		pp.Println("error fetching history")
	}*/

	//get transaction
	/*transactionDetails, err := metahash.GetTransaction(trx)
	if err == nil {
		pp.Println("get-tx", transactionDetails) //TODO: check ifthe history is not nil
	} else {
		pp.Println("error getting transaction details")
	}*/

	//get transaction
	/*lastTx, err := metahash.GetLastTransactions()
	if err == nil {
		pp.Println("last transactions", lastTx) //TODO: check ifthe history is not nil
	} else {
		pp.Println("error getting last transactions")
	}*/

	/*blocks, err := metahash.GetBlocks(500, 2)
	if err == nil {
		pp.Println("blocks", blocks) //TODO: check ifthe history is not nil
	} else {
		pp.Println("error getting blocks")
	}*/

	/*block, err := metahash.GetBlockByNumber(500, 1)
	if err == nil {
		pp.Println("block-number", block) //TODO: check ifthe history is not nil
	} else {
		pp.Println("error getting block by number transactions")
	}*/

	numBlocks, err := metahash.GetTotalBlocks()
	if err == nil {
		pp.Println("total number of blocks", numBlocks) //TODO: check ifthe history is not nil
	} else {
		pp.Println("error getting blocks")
	}

}
