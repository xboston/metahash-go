package metahash

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const torUrl = "http://tor.net-main.metahashnetwork.com:5795"
const totalSupplyUrl = "https://app.metahash.io/api/stat/?method=supply"

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

// FetchBalances gets the balance information of list of addresses
func FetchBalances(addresses ...string) ([]*Balance, error) {
	responseBalance, err := metahashClient.Call("fetch-balances", &BalancesArgs{Addresses: addresses})
	if err == nil {
		var resultBalances []*Balance
		err = responseBalance.GetObject(&resultBalances)
		if err == nil {
			return resultBalances, nil
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
func FetchHistoryFiter(address string, startIndex, numTrx int64) ([]*TransactionInfo, error) {
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

// GetBlockByNumber returns block info. Pass nil for startTx, numTx
// blockType should between 0 and 2: 0 = 0 or there isn’t - only block name, 1 = only hashes, 2 = full block dump,
//issue: block type 1 is not working at the moment

func GetBlockByHash(blockHash string, startIndex, numTx, blockType int) (*Block, error) {
	if blockType > 2 || blockType == 1 { //skip block type 1 as it not working properly
		blockType = 0
	}
	blkArg := &BlockByHashArgs{Hash: blockHash, BeginTx: int64(startIndex), CountTxs: int64(numTx), Type: int8(blockType)}
	responseBlockByNumber, err := metahashClient.Call("get-block-by-hash", blkArg)
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
func GetDumpBlockByHash(blockHash string, isHex bool) (*DumpBlock, error) {
	arg := &DumpBlockByHashArgs{Hash: blockHash, IsHex: isHex}
	responseDumpBlockByHash, err := metahashClient.Call("get-dump-block-by-hash", arg)
	if err == nil {
		var dumpBlock *DumpBlock
		err = responseDumpBlockByHash.GetObject(&dumpBlock)
		if err == nil {
			return dumpBlock, nil
		}
		return nil, err
	}
	return nil, err
}

func GetDumpBlockByNumber(blockNumber int64, isHex bool) (*DumpBlock, error) {
	blockArg := &DumpBlockByNumberArgs{Number: blockNumber, IsHex: isHex}
	responseDumpBlockByNumber, err := metahashClient.Call("get-dump-block-by-number", blockArg)
	if err == nil {
		var dumbBlock *DumpBlock
		err = responseDumpBlockByNumber.GetObject(&dumbBlock)
		if err == nil {
			return dumbBlock, nil
		}
		return nil, err
	}
	return nil, err
}

func GetNodeStats(address string) (*NodeStats, error) {
	args := &NodeArgs{
		Address: address,
	}
	responsNodeStats, err := metahashClient.Call("get-last-node-stat-result", args)
	if err == nil {
		var nodeStats *NodeStats
		err = responsNodeStats.GetObject(&nodeStats)
		if err == nil {
			return nodeStats, nil
		}
		return nil, err
	}
	return nil, err
}

func (ns *NodeStats) GetTest() (*NodeStatsState, error) {
	var nodeTest *NodeStatsState

	str := fmt.Sprint(ns.State)
	err := json.Unmarshal([]byte(str), &nodeTest)
	if err == nil {
		return nodeTest, nil
	}
	return nil, errors.New("cannot parse state information")
}

func GetLastNodeStatTrust(address string) (*NodeTrust, error) {
	args := &NodeArgs{
		Address: address,
	}

	nodeTrustResp, err := metahashClient.Call("get-last-node-stat-trust", args)
	if err == nil {
		var nodeTrust *NodeTrust
		err = nodeTrustResp.GetObject(&nodeTrust)
		if err == nil {
			return nodeTrust, nil
		}
		return nil, err
	}
	return nil, err
}
func GetLastNodeCount(counts int) (*LastNodeCount, error) {
	arg := NodeArgs{CountTests: counts}

	nodeLastNodeCountResp, err := metahashClient.Call("get-all-last-nodes-count", arg)
	if err == nil {
		var nodeNodeCount *LastNodeCount
		err = nodeLastNodeCountResp.GetObject(&nodeNodeCount)
		if err == nil {
			return nodeNodeCount, nil
		}
		return nil, err
	}
	return nil, err
}

// GetTrustData returns trust and list of delete to and delegate from
func (nt *NodeTrust) GetTrustData() (*NodeTrustData, error) {
	var nodeTrustData *NodeTrustData

	str := fmt.Sprint(nt.Data)
	fmt.Print("node trust data: ", str)
	err := json.Unmarshal([]byte(str), &nodeTrustData)
	if err == nil {
		return nodeTrustData, nil
	}
	return nil, errors.New("cannot parse node trust data")
}

func GetNodeRaiting(address string, countTests int) (*NodeRaiting, error) {
	args := &NodeArgs{
		Address:    address,
		CountTests: countTests,
	}

	nodeRaitingResp, err := metahashClient.Call("get-nodes-raiting", args)
	if err == nil {
		var nodeRaiting *NodeRaiting
		err = nodeRaitingResp.GetObject(&nodeRaiting)
		if err == nil {
			return nodeRaiting, nil
		}
		return nil, err
	}
	return nil, err
}

func GetAddressDelegations(address string, startTx, countTx int64) (*AddressDelegations, error) {
	args := &NodeArgs{
		Address:  address,
		BeginTx:  startTx,
		CountTxs: countTx,
	}
	responsNodeStats, err := metahashClient.Call("get-address-delegations", args)
	if err == nil {
		var nodeStats *AddressDelegations
		err = responsNodeStats.GetObject(&nodeStats)
		if err == nil {
			return nodeStats, nil
		}
		return nil, err
	}
	return nil, err
}

// GetForgingSumAll returns sum all types of forging
//100; // forging transaction
//101; // wallet reward forging transaction
//102; // node reward forging transaction
//103; // coin reward forging transaction
//104; // random reward forging transaction
func GetForgingSumAll() (*ForgingSum, error) {
	responseLastTxs, err := metahashClient.Call("get-forging-sum-all", nil)
	if err == nil {
		var forgingSum *ForgingSum
		err = responseLastTxs.GetObject(&forgingSum)
		if err == nil {
			return forgingSum, nil
		}
		return nil, err
	}

	return nil, err
}

func GetForgingSum(blockIndent int) (*ForgingSum, error) {
	arg := &ForginSumArgs{BlockIndent: blockIndent}
	responseLastTxs, err := metahashClient.Call("get-forging-sum", arg)
	if err == nil {
		var forgingSum *ForgingSum
		err = responseLastTxs.GetObject(&forgingSum)
		if err == nil {
			return forgingSum, nil
		}
		return nil, err
	}
	return nil, err
}

func GetCommonBalance() (int64, error) {
	balanceRes, err := metahashClient.Call("get-common-balance", nil)

	if err == nil {
		var balance *CommonBalance
		err = balanceRes.GetObject(&balance)
		if err == nil {
			return balance.Balance, nil
		}
		return 0, err
	}

	return 0, err
}

func MetahashSupply() (*TotalSupply, error) {

	req, _ := http.NewRequest("GET", totalSupplyUrl, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {

		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {

		return nil, err

	}

	supply := &TotalSupply{}
	err = json.Unmarshal(body, supply)
	if err == nil {
		return supply, nil
	}
	return nil, err

}
