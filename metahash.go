package metahash

type BalanceArgs struct {
	Address string `json:"address"`
}
type BalancesArgs struct {
	Addresses []string `json:"addresses"`
}

type Balance struct {
	Address           string `json:"address"`
	Received          int64  `json:"received"`
	Spent             int64  `json:"spent"`
	CountReceived     int64  `json:"count_received"`
	CountSpent        int64  `json:"count_spent"`
	CountTxs          int64  `json:"count_txs"`
	BlockNumber       int64  `json:"block_number"`
	CurrentBlock      int64  `json:"currentBlock"`
	Hash              string `json:"hash"`
	CountDelegatedOps int64  `json:"countDelegatedOps"`
	Delegate          int64  `json:"delegate"`
	Undelegate        int64  `json:"undelegate"`
	Delegated         int64  `json:"delegated"`
	Undelegated       int64  `json:"undelegated"`
	Reserved          int64  `json:"reserved"`
	CountForgedOps    int64  `json:"countForgedOps"`
	Forged            int64  `json:"forged"`
}

type TransactionArgs struct {
	Hash string `json:"hash"`
}

type Transaction struct {
	Transaction TransactionInfo `json:"transaction"`
	CountBlocks int64           `json:"countBlocks"`
	KnownBlocks int64           `json:"knownBlocks"`
}

type HistoryFilter struct {
	IsInput    bool `json:"isInput,omitempty"`    // isInput - Display only isInput transactions
	IsOutput   bool `json:"isOutput,omitempty"`   //isOutput - Display only isOutput transactions
	IsSuccess  bool `json:"isSuccess,omitempty"`  //isSuccess - Display only success transactions
	IsForging  bool `json:"isForging,omitempty"`  //isForging - Display only forging transactions
	IsTest     bool `json:"isTest,omitempty"`     //isTest - Display only test transactions
	IsDelegate bool `json:"isDelegate,omitempty"` //isDelegate - Display only delegation transactions
}
type HistoryArgs struct {
	Address  string        `json:"address"`
	BeginTx  int64         `json:"beginTx,omitempty"`
	CountTxs int64         `json:"countTxs,omitempty"`
	Filters  HistoryFilter `json:"filters,omitempty"`
}

type BlockByNumberArgs struct {
	Number   int64 `json:"number"`
	BeginTx  int64 `json:"beginTx,omitempty"`
	CountTxs int64 `json:"countTxs,omitempty"`
	Type     int8  `json:"type,omitempty"` // 0-4
	// 0 - простой тип с 7 подписями
	// 1 - 7 подписей и массив хешей всех транзакций
	// 2 - 7 подписей и все транзакции полностью
	// 3 - краткий формат блока, только прошлый и текущий хеш
	// 4 - краткий формат 3 + размер блока и расположение файла данных
}

type BlockByHashArgs struct {
	Hash     string `json:"hash"`
	BeginTx  int64  `json:"beginTx,omitempty"`
	CountTxs int64  `json:"countTxs,omitempty"`
	Type     int8   `json:"type,omitempty"` // 0-4
}

type BlocksArgs struct {
	CountBlocks int64 `json:"countBlocks,omitempty"`
	BeginBlock  int64 `json:"beginBlock,omitempty"`
}

type Block struct {
	Type       string `json:"type"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prev_hash"`
	TxHash     string `json:"tx_hash"`
	Number     int64  `json:"number"`
	TimeStamp  int64  `json:"timestamp"`
	CountTxs   int64  `json:"count_txs"`
	Sign       string `json:"sign"`
	Size       int64  `json:"size"`
	FileName   string `json:"fileName"`
	Signatures []struct {
		From        string `json:"from"`
		To          string `json:"to"`
		Value       int64  `json:"value"`
		Transaction string `json:"transaction"`
		Data        string `json:"data"`
		TimeStamp   int64  `json:"timestamp"`
		Type        string `json:"type"`
		BlockNumber int64  `json:"blockNumber"`
		Signature   string `json:"signature"`
		Publickey   string `json:"publickey"`
		Fee         int64  `json:"fee"`
		RealFee     int64  `json:"realFee"`
		Nonce       int64  `json:"nonce"`
		IntStatus   int64  `json:"intStatus"`
		Status      string `json:"status"`
	} `json:"signatures,omitempty"`
	Txs []*TransactionInfo `json:"txs"`
}

type LastTxsArgs struct {
}

type TransactionInfo struct {
	From         string `json:"from" db:"fromA"`
	To           string `json:"to" db:"toA"`
	Value        int64  `json:"value"`
	Transaction  string `json:"transaction"`
	Data         string `json:"data"`
	TimeStamp    int64  `json:"timestamp" db:"timestamp,int64"`
	Type         string `json:"type" db:"typeTx"`
	BlockNumber  int64  `json:"blockNumber" db:"blockNumber"`
	Signature    string `json:"signature"`
	PublicKey    string `json:"publickey"`
	Fee          int64  `json:"fee"`
	RealFee      int64  `json:"realFee" db:"realFee"`
	Nonce        int64  `json:"nonce"`
	IntStatus    int64  `json:"intStatus" db:"intStatus"`
	Status       string `json:"status"`
	IsDelegate   bool   `json:"isDelegate,omitempty" db:"isDelegate"`
	DelegateInfo struct {
		IsDelegate   bool   `json:"isDelegate"`
		Delegate     int64  `json:"delegate,omitempty"`
		DelegateHash string `json:"delegateHash,omitempty"`
	} `json:"delegate_info,omitempty" db:"-"`
	Delegate     int64  `json:"delegate,omitempty"`
	DelegateHash string `json:"delegateHash,omitempty" db:"delegateHash"`
}

type DumpBlockByNumberArgs struct {
	Number int64 `json:"number"`
	IsHex  bool  `json:"isHex,omitempty"`
}

type DumpBlockByHashArgs struct {
	Hash  string `json:"hash"`
	IsHex bool   `json:"isHex,omitempty"`
}

type DumpBlock struct {
	Dump string `json:"dump"`
}

type CountBlocksArgs struct {
}

type CountBlocks struct {
	CountBlocks int64 `json:"count_blocks"`
	BeginBlock  int64 `json:"beginBlock"`
}

// внутренние структуры-методы
type NodeRegistration struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Host string `json:"host"`
		Name string `json:"name"`
		Type string `json:"type,omitempty"`
	} `json:"params"`
}

type AbstractMethod struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type NodeArgs struct {
	Address  string `json:"address"`
	BeginTx  int64  `json:"beginTx,omitempty"`
	CountTxs int64  `json:"countTxs,omitempty"`
}

type NodeTest struct {
	TimeStamp int64  `json:"timestamp"`
	Method    string `json:"method"`
	Params    struct {
		Mhaddr   string `json:"mhaddr"`
		IP       string `json:"ip"`
		QPS      int64  `json:"qps"`
		Rps      int64  `json:"rps"`
		Closed   string `json:"closed"`
		Timeouts string `json:"timeouts"`
		Ver      string `json:"ver"`
		Success  string `json:"success"`
	} `json:"params"`
}

type NodeStats struct {
	Address            string `json:"address"`
	Type               string `json:"type"`
	AvgRps             string `json:"avgRps"`
	IsLatency          bool   `json:"isLatency"`
	IP                 string `json:"ip"`
	Geo                string `json:"geo"`
	State              string `json:"state"`
	Timestamp          int    `json:"timestamp"`
	Success            bool   `json:"success"`
	LastBlockTimestamp int    `json:"lastBlockTimestamp"`
}

type NodeStatsState struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Type              string `json:"type"`
		Ver               string `json:"ver"`
		Address           string `json:"address"`
		Host              string `json:"host"`
		Blockheightcheck  string `json:"blockHeightCheck"`
		Requestsperminute string `json:"requestsPerMinute"`
		Latency           string `json:"latency"`
		Geo               string `json:"geo"`
		Success           string `json:"success"`
	} `json:"params"`
}

type NodeTrust struct {
	Address            string `json:"address"`
	Trust              int64  `json:"trust"`
	Data               string `json:"data"`
	Timestamp          int64  `json:"timestamp"`
	Lastblocktimestamp int64  `json:"lastBlockTimestamp"`
}

type NodeTrustData struct {
	State      int `json:"state"`
	Trust      int `json:"trust"`
	DelegateTo []struct {
		A string `json:"a"`
		V int64  `json:"v"`
	} `json:"delegate_to"`
	DelegatedFrom []struct {
		A string `json:"a"`
		V int64  `json:"v"`
	} `json:"delegated_from"`
}

type AddressDelegations struct {
	Address string `json:"address"`
	States  []struct {
		To    string `json:"to"`
		Value int64  `json:"value"`
		Tx    string `json:"tx"`
	} `json:"states"`
}

type ForgingSum struct {
	Blocknumber int `json:"blockNumber"`
	Sums        []struct {
		Type  int   `json:"type"`
		Value int64 `json:"value"`
	} `json:"sums"`
}
