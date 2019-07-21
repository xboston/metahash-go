package metahash

type BalanceArgs struct {
	Address string `json:"address"`
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

type HistoryArgs struct {
	Address  string `json:"address"`
	BeginTx  int64  `json:"beginTx,omitempty"`
	CountTxs int64  `json:"countTxs,omitempty"`
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
