package main

import (
	"github.com/k0kubun/pp"
	"github.com/xboston/metahash-go/metahash"
)

var addr = "http://tor.net-main.metahashnetwork.com:5795"
var wall = "0x00fa2a5279f8f0fd2f0f9d3280ad70403f01f9d62f52373833"
var trxHash = "ec81dec38082d238bf88af3c39128b0f52a2bbb3038a277a75015684e5f93ca9"
var blockHash = "0b3dafdd27595922daa304ebaf5f0a3b190a62a0af47de83ed359a303c1d4765"

func main() {
	metahashClient := metahash.NewClient(addr)

	responseBalance, err := metahashClient.Call("fetch-balance", &metahash.BalanceArgs{Address: wall})
	if err == nil {
		var resultBalance *metahash.Balance
		err = responseBalance.GetObject(&resultBalance)
		if err == nil {
			pp.Println("fetch-balance", resultBalance)
		}
	}

	responseHistory, err := metahashClient.Call("fetch-history", &metahash.HistoryArgs{Address: wall})
	if err == nil {
		var resultHistory []*metahash.TransactionInfo
		err = responseHistory.GetObject(&resultHistory)
		if err == nil {
			pp.Println("fetch-history", resultHistory)
		}
	} else {
		pp.Println("err", err.Error())
	}

	responseTransaction, err := metahashClient.Call("get-tx", &metahash.TransactionArgs{Hash: trxHash})
	if err == nil {
		var resultTransaction *metahash.Transaction
		err = responseTransaction.GetObject(&resultTransaction)
		if err == nil {
			pp.Println("get-tx", resultTransaction)
		}
	} else {
		pp.Println("err", err.Error())
	}

	responseBlockByNumber, err := metahashClient.Call("get-block-by-number", &metahash.BlockByNumberArgs{Number: 1})
	if err == nil {
		var resultBlockByNumber *metahash.Block
		err = responseBlockByNumber.GetObject(&resultBlockByNumber)
		if err == nil {
			pp.Println("get-block-by-number", resultBlockByNumber)
		}
	} else {
		pp.Println("err", err.Error())
	}

	responseBlockByHash, err := metahashClient.Call("get-block-by-hash", &metahash.BlockByHashArgs{Hash: blockHash})
	if err == nil {
		var resultBlockByHash *metahash.Block
		err = responseBlockByHash.GetObject(&resultBlockByHash)
		if err == nil {
			pp.Println("get-block-by-hash", resultBlockByHash)
		}
	} else {
		pp.Println("err", err.Error())
	}

	responseLastTxs, err := metahashClient.Call("get-last-txs", &metahash.LastTxsArgs{})
	if err == nil {
		var resultLastTxs []*metahash.TransactionInfo
		err = responseLastTxs.GetObject(&resultLastTxs)
		if err == nil {
			pp.Println("get-last-txs", resultLastTxs)
		}
	} else {
		pp.Println("err", err.Error())
	}

	responseBlocks, err := metahashClient.Call("get-blocks", &metahash.BlocksArgs{CountBlocks: 2})
	if err == nil {
		var resultBlocks []*metahash.Block
		err = responseBlocks.GetObject(&resultBlocks)
		if err == nil {
			pp.Println("get-blocks", resultBlocks)
		}
	} else {
		pp.Println("err", err.Error())
	}

	responseDumpBlockByNumber, err := metahashClient.Call("get-dump-block-by-number", &metahash.DumpBlockByNumberArgs{Number: 1, IsHex: true})
	if err == nil {
		var resultDumpBlockByNumber *metahash.DumpBlock
		err = responseDumpBlockByNumber.GetObject(&resultDumpBlockByNumber)
		if err == nil {
			pp.Println("get-dump-block-by-number", resultDumpBlockByNumber)
		}
	} else {
		pp.Println("err", err.Error())
	}

	responseDumpBlockByHash, err := metahashClient.Call("get-dump-block-by-hash", &metahash.DumpBlockByHashArgs{Hash: blockHash, IsHex: true})
	if err == nil {
		var resultDumpBlockByHash *metahash.DumpBlock
		err = responseDumpBlockByHash.GetObject(&resultDumpBlockByHash)
		if err == nil {
			pp.Println("get-dump-block-by-hash", resultDumpBlockByHash)
		}
	} else {
		pp.Println("err", err.Error())
	}

	responseCountBlocks, err := metahashClient.Call("get-count-blocks", &metahash.CountBlocksArgs{})
	if err == nil {
		var resultCountBlocks *metahash.CountBlocks
		err = responseCountBlocks.GetObject(&resultCountBlocks)
		if err == nil {
			pp.Println("get-count-blocks", resultCountBlocks)
		}
	} else {
		pp.Println("err", err.Error())
	}
}
