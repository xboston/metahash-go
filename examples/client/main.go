package main

import (
	"github.com/k0kubun/pp"
	"github.com/xboston/metahash-go"
)

//sample address
const address = "0x0062b734726a66f54e9ac72f63af9fd18bd5ca5583061e2c37"         //serengeti node
const address2 = "0x005b891007c2000fee08e085beb91494f1d3753eb8eee354f0"        //serengeti node
const trx = "361d5e113648366aa46010fa02ff1f66dcf8045abe4e30547ffaba0506edbdb2" //sample transaction
const blockHash = "49a840968be1a3c0449a6d7946fd021ce2746006ff5b855fdc0623669833f5e7"

func main() {

	//to test the api function, you can uncoment each function
	//fetch balance
	balance, err := metahash.FetchBalance(address)
	if err == nil {
		pp.Println("fetch-balance", balance)
	} else {
		pp.Println("error fetching balance")
	}

	//fetch balances
	/*balances, err := metahash.FetchBalances(address, address2)
	if err == nil {
		pp.Println("fetch-balances", balances)
	} else {
		pp.Println("error fetching balances")
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
		pp.Println("fetch-history", historyRange)
	} else {
		pp.Println("error fetching history")
	}*/

	//get transaction
	/*transactionDetails, err := metahash.GetTransaction(trx)
	if err == nil {
		pp.Println("get-tx", transactionDetails)
	} else {
		pp.Println("error getting transaction details")
	}*/

	//get transaction
	/*lastTx, err := metahash.GetLastTransactions()
	if err == nil {
		pp.Println("last transactions", lastTx)
	} else {
		pp.Println("error getting last transactions")
	}*/

	/*blocks, err := metahash.GetBlocks(500, 2)
	if err == nil {
		pp.Println("blocks", blocks)
	} else {
		pp.Println("error getting blocks")
	}*/

	/*block, err := metahash.GetBlockByNumber(500, 1)
	if err == nil {
		pp.Println("block-number", block)
	} else {
		pp.Println("error getting block by number transactions")
	}*/

	/*numBlocks, err := metahash.GetTotalBlocks()
	if err == nil {
		pp.Println("total number of blocks", numBlocks)
	} else {
		pp.Println("error getting blocks")
	}*/

	/*
		//get block by hash
		block, err := metahash.GetBlockByHash(blockHash, 0, 100, 1)
		if err == nil {
			pp.Println("block by hash", block)
		} else {
			pp.Println("error getting block by hash")
		}*/

	/*
		// dumpblock by hash
		block, err := metahash.GetDumpBlockByHash(blockHash, true)
		if err == nil {
			pp.Println("dump blokc by hash", block)
		} else {
			pp.Println("error getting dump block by hash")
		}*/

	/*block, err := metahash.GetDumpBlockByNumber(500, true)
	if err == nil {
		pp.Println("dump blokc by number", block)
	} else {
		pp.Println("error getting dump block by number")
	}*/

	/*stats, err := metahash.GetNodeStats(address)
	if err == nil {
		pp.Println("node stats", stats)
	} else {
		pp.Printf("error getting node stats: %v\n", err)
	}

	tests, err := stats.GetTest()
	if err == nil {
		pp.Println("node test data", tests)
	} else {
		pp.Printf("error getting node test data: %v\n", err)
	}*/

	/*trust, err := metahash.GetLastNodeStatTrust(address)
	if err == nil {
		pp.Println("node trust", trust)
	} else {
		pp.Printf("error getting node trust: %v\n", err)
	}

	trustData, err := trust.GetTrustData()
	if err == nil {
		pp.Println("node trust data", trustData)
	} else {
		pp.Printf("error getting node trust data: %v\n", err)
	}*/

	/*lastNodeCount, err := metahash.GetLastNodeCount(20)
	if err == nil {
		pp.Println("last node count", lastNodeCount)
	} else {
		pp.Printf("error getting last node count: %v\n", err)
	}*/

	/*raiting, err := metahash.GetNodeRaiting(address, 100)
	if err == nil {
		pp.Println("raiting", raiting)
	} else {
		pp.Printf("error getting raiting: %v\n", err)
	}*/

	/*addDelegations, err := metahash.GetAddressDelegations(address, 0, 10)
	if err == nil {
		pp.Println("address delegations", addDelegations)
	} else {
		pp.Printf("error getting address delegations: %v\n", err)
	}*/
	/*forgingSum, err := metahash.GetForgingSumAll()
	if err == nil {
		pp.Println("forging sum", forgingSum)
	} else {
		pp.Printf("error getting forging sum: %v\n", err)
	}*/

	/*forgingSum, err := metahash.GetForgingSum(3)
	if err == nil {
		pp.Println("forging sum", forgingSum)
	} else {
		pp.Printf("error getting forging sum: %v\n", err)
	}*/

	/*balance, err := metahash.GetCommonBalance()
	if err == nil {
		pp.Println("balance", balance)
	} else {
		pp.Printf("error getting balance: %v\n", err)
	}*/

}
