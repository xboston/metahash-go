package metahash

const torUrl = "http://tor.net-main.metahashnetwork.com:5795"

var metahashClient RPCClient

func init() {
	metahashClient = NewClient(torUrl)
}

// FetchBalance gets the balance information of a given address
func FetchBalance(address string) (*Balance, error) {
	responseBalance, err := metahashClient.Call("fetch-balance", &BalanceArgs{Address: address})
	if err == nil {
		var resultBalance *Balance
		err = responseBalance.GetObject(&resultBalance)
		if err == nil {
			return resultBalance, nil
		}
		return nil, err
	}
	return nil, err
}

// FetchHistory returns all transactions history of a given address
func FetchHistory(address string) ([]*TransactionInfo, error) {
	responseHistory, err := metahashClient.Call("fetch-history", &HistoryArgs{Address: address})
	if err == nil {
		var resultHistory []*TransactionInfo
		err = responseHistory.GetObject(&resultHistory)
		if err == nil {
			return resultHistory, nil
		}
		return nil, err
	}
	return nil, err
}

//FetchHistoryRange returns list of transaction history from a given index
func FetchHistoryRange(address string, startIndex, numTrx int64) ([]*TransactionInfo, error) {
	responseHistory, err := metahashClient.Call("fetch-history", &HistoryArgs{Address: address, BeginTx: startIndex, CountTxs: numTrx})
	if err == nil {
		var resultHistory []*TransactionInfo
		err = responseHistory.GetObject(&resultHistory)
		if err == nil {
			return resultHistory, nil
		}
		return nil, err
	}
	return nil, err
}

/*
//This function is not working because of the metahash api error
//FetchHistoryFilter returns list of transaction history based on the provide filter
func FetchHistoryFilter(address string, filter *HistoryFilter) ([]*TransactionInfo, error) {
	responseHistory, err := metahashClient.Call("fetch-history-filter", &HistoryArgs{Address: address, Filters: *filter})
	if err == nil {
		pp.Println("response", responseHistory) //TODO: check ifthe history is not nil

		var resultHistory []*TransactionInfo
		err = responseHistory.GetObject(&resultHistory)
		if err == nil {
			return resultHistory, nil
		}
		return nil, err
	}
	return nil, err
}*/

// GetTransaction  returns the transaction details given the transaction hash
func GetTransaction(txHash string) (*Transaction, error) {
	responseTransaction, err := metahashClient.Call("get-tx", &TransactionArgs{Hash: txHash})
	if err == nil {
		var resultTransaction *Transaction
		err = responseTransaction.GetObject(&resultTransaction)
		if err == nil {
			return resultTransaction, nil
		}
		return nil, err
	}
	return nil, err
}

// GetLastTransactions returns the list of last transactions made
func GetLastTransactions() ([]*TransactionInfo, error) {
	responseLastTxs, err := metahashClient.Call("get-last-txs", &LastTxsArgs{})
	if err == nil {
		var resultLastTxs []*TransactionInfo
		err = responseLastTxs.GetObject(&resultLastTxs)
		if err == nil {
			return resultLastTxs, nil
		}
		return nil, err
	}

	return nil, err
}

//GetBlocks

func GetBlocks(startBlock, numBlocks int64) ([]*Block, error) {
	responseBlocks, err := metahashClient.Call("get-blocks", &BlocksArgs{CountBlocks: numBlocks, BeginBlock: startBlock})
	if err == nil {
		var resultBlocks []*Block
		err = responseBlocks.GetObject(&resultBlocks)
		if err == nil {
			return resultBlocks, nil
		}
		return nil, err
	}
	return nil, err
}

// GetBlockByNumber returns block info. Pass nil for startTx, numTx
// blockType should between 0 and 2: 0 = 0 or there isn’t - only block name, 1 = only hashes, 2 = full block dump,
//issue: block type 1 is not working at the moment

func GetBlockByNumber(blockNumber, blockType int64) (*Block, error) {

	if blockType > 2 || blockType == 1 { //skip block type 1 as it not working properly
		blockType = 0
	}

	blkArg := &BlockByNumberArgs{Number: blockNumber, Type: int8(blockType)}
	responseBlockByNumber, err := metahashClient.Call("get-block-by-number", blkArg)
	if err == nil {
		var resultBlockByNumber *Block
		err = responseBlockByNumber.GetObject(&resultBlockByNumber)
		if err == nil {
			return resultBlockByNumber, nil
		}
		return nil, err
	}
	return nil, err
}

func GetBlockByNumberFilter(blockNumber, startTx, numTx, blockType int64) (*Block, error) {

	if blockType > 2 || blockType == 1 { //skip block type 1 as it not working properly
		blockType = 0
	}

	blkArg := &BlockByNumberArgs{Number: blockNumber, BeginTx: startTx, CountTxs: numTx, Type: int8(blockType)}
	responseBlockByNumber, err := metahashClient.Call("get-block-by-number", blkArg)
	if err == nil {
		var resultBlockByNumber *Block
		err = responseBlockByNumber.GetObject(&resultBlockByNumber)
		if err == nil {
			return resultBlockByNumber, nil
		}
		return nil, err
	}
	return nil, err
}

func GetTotalBlocks() (int64, error) {
	responseCountBlocks, err := metahashClient.Call("get-count-blocks", &CountBlocksArgs{})
	if err == nil {
		var resultCountBlocks *CountBlocks
		err = responseCountBlocks.GetObject(&resultCountBlocks)
		if err == nil {
			return resultCountBlocks.CountBlocks, nil
		}
		return 0, err
	}
	return 0, err
}
