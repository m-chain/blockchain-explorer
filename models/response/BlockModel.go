package response

import "time"

type BlockInfo struct {
	Status bool      `json:"status"`
	Data   BlockData `json:"data"`
	Msg    string    `json:"msg"`
}

type NodeInfo struct {
	Code int16      `json:"code"`
	Data []NodeData `json:"data"`
	Msg  string     `json:"msg"`
}

type BlockExtInfo struct {
	Status bool         `json:"status"`
	Data   BlockExtData `json:"data"`
	Msg    string       `json:"msg"`
}

type BlockData struct {
	Blockid      string              `json:"blockid"`
	PreviousHash string              `json:"previous_hash"`
	DataHash     string              `json:"data_hash"`
	Transactions []BlockTransactions `json:"transactions"`
}

type BlockExtData struct {
	Blockid      int64   `json:"blockId"`
	BlockSize    float64 `json:"blocksize"`
	PreviousHash string  `json:"previousHash"`
	CreateTime   string  `json:"createTime"`
	ChannelId    string  `json:"channelid"`
	BlockHash    string  `json:"currentHash"`
}

type BlockTransactions struct {
	Channelname      string            `json:"channelname"`
	Chaincodename    string            `json:"chaincodename"`
	Chaincodeversion string            `json:"chaincodeversion"`
	Txhash           string            `json:"txhash"`
	Createdt         string            `json:"createdt"`
	Wallets          []BlockWallets    `json:"wallets"`
	Trans            []BlockTrans      `json:"trans"`
	Tokens           []BlockTokens     `json:"tokens"`
	Contracts        []BlockContracts  `json:"contracts"`
	GasReturns       []BlockGasReturns `json:"gasReturns"`
	TokenMaster      map[string]string `json:"tokenMaster"`
	Others           interface{}       `json:"others"`
}

type BlockWallets struct {
	Address      string      `json:"address"`
	PubKey       string      `json:"pubKey"`
	CreateTime   string      `json:"createTime"`
	ReleasePlans interface{} `json:"releasePlans"`
	IsLocked     int64       `json:"isLocked"`
}

type WalletBalance struct {
	Status bool                `json:"status"`
	Data   WalletBalanceDetail `json:"data"`
	Msg    string              `json:"msg"`
}

type WalletBalanceDetail struct {
	Address     string              `json:"address" description:"地址"`
	PubKey      string              `json:"pubKey" description:"公钥"`
	CreateTime  string              `json:"createTime" description:"创建时间"`
	IsLocked    int64               `json:"isLocked" description:"是否锁定"`
	WalletRests []WalletRestsDetail `json:"walletRests"`
}

type WalletRestsDetail struct {
	FreezeNumber string `json:"freezeNumber" description:"冻结数量"`
	IsBaseCoin   bool   `json:"isBaseCoin" description:"是否主币"`
	RestNumber   string `json:"restNumber" description:"可用数量"`
	TokenID      string `json:"tokenID" description:"代币ID"`
	TokenSymbol  string `json:"tokenSymbol" description:"代币简称"`
}

type BlockTrans struct {
	TokenID     string `json:"tokenID"`
	TxId        string `json:"txId"`
	Blockid     string `json:"blockid"`
	FromAddress string `json:"fromAddress"`
	ToAddress   string `json:"toAddress"`
	Number      string `json:"number"`
	GasUsed     string `json:"gasUsed"`
	Time        string `json:"time"`
	Nonce       string `json:"nonce"`
	State       int    `json:"state"`
	Msg         string `json:"msg"`
	Notes       string `json:"notes"`
	Reason      string `json:"reason"`
}

type BlockTokenInfo struct {
	Status bool        `json:"status"`
	Data   BlockTokens `json:"data"`
	Msg    string      `json:"msg"`
}

type BlockTokens struct {
	TokenID      string      `json:"tokenID"`
	Name         string      `json:"name"`
	TokenSymbol  string      `json:"tokenSymbol"`
	IsBaseCoin   bool        `json:"isBaseCoin"`
	DecimalUnits int64       `json:"decimalUnits"`
	TotalNumber  string      `json:"totalNumber"`
	IssuePrice   string      `json:"issuePrice"`
	IssueTime    string      `json:"issueTime"`
	Status       int8        `json:"status"`
	RestNumber   string      `json:"restNumber"`
	IssueRecords interface{} `json:"issueRecords"`
}

type BlockGasReturns struct {
	Id                int64  `json:"id"`
	TokenID           string `json:"tokenID"`
	StartTime         int64  `json:"startTime"`
	InitReleaseAmount string `json:"initReleaseAmount"`
	Amount            string `json:"amount"`
	Interval          int64  `json:"interval"`
	Periods           int64  `json:"periods"`
	Withdrawed        string `json:"withdrawed"`
	TxID              string `json:"txID"`
	Reason            string `json:"reason"`
	CreateTime        string `json:"createTime"`
	Address           string `json:"address"`
	ContractAddress   string `json:"contractAddress"`
	GasReturnhash     string `json:"gasReturnHash"`
}

type BlockContracts struct {
	Name            string `json:"name"`
	ContractAddress string `json:"contractAddress"`
	ContractSymbol  string `json:"contractSymbol"`
	Maddress        string `json:"mAddress"` //拥有平台币地址
	Version         string `json:"version"`
	ContractPath    string `json:"ccPath"`
	Remark          string `json:"remark"`
	Status          string `json:"status"`
	CreateTime      string `json:"createTime"`
	UpdateTime      string `json:"updatetime"`
}

type NodeDetail struct {
	Id          int64     `orm:"column(id);pk;"json:"id"` // 设置主键
	NodeName    string    `json:"nodename"`
	NodeIp      string    `json:"nodeip"`
	NodeAddress string    `json:"nodeaddress"`
	RunStatus   int       `json:"runstatus"`
	BlockHeight int       `json:"blockheight"`
	UpdateTime  time.Time `json:"updatetime"`
}

type NodeData struct {
	PeerIp        string `json:"peer_ip"`
	PeerLocation  string `json:"peer_location"`
	PeerName      string `json:"peer_name"`
	PeerStatus    string `json:"peer_status"`
	ContainerId   string `json:"container_id"`
	PeerLatitude  string `json:"peer_latitude"`
	PeerLongitude string `json:"peer_longitude"`
}

type InfoPageModel struct {
	TotalNum  int         `json:"totalnum"`
	TotalPage int         `json:"totalpage"`
	PageNo    int         `json:"pageno"`
	Data      interface{} `json:"data"`
}

type SearchInformationModel struct {
	ResultContent string `json:"resultcontent" `
	ResultType    int    `json:"resulttype" description:"结果类型(Empty = 0 BlockId = 1 Address = 2 TokenId = 3  Contract = 4  Transaction = 5)"`
}

type Gasrecords struct {
	Status bool              `json:"status"`
	Data   []GasrecordsModel `json:"data"`
	Msg    string            `json:"msg"`
}

type GasrecordsModel struct {
	Id           int64  `json:"id"`
	CcName       string `json:"ccName"`
	CcVersion    string `json:"ccVersion"`
	FcnName      string `json:"fcnName"`
	Args         string `json:"args"`
	ContractAddr string `json:"contractAddr"`
	Address      string `json:"address"`
	GasUsed      int64  `json:"gasUsed"`
	TxId         string `json:"txId"`
	ReturnTxId   string `json:"returnTxId"`
	CreateTime   string `json:"createTime"`
}

type ChainInfoData struct {
	Status bool      `json:"status"`
	Data   ChainData `json:"data"`
	Msg    string    `json:"msg"`
}

type ChainData struct {
	High         int64  `json:"high"`
	Low          int64  `json:"low""`
	PreviousHash string `json:"previousBlockHash"`
	CurrentHash  string `json:"currentBlockHash"`
}

type TxStatusInfoData struct {
	Status bool           `json:"status"`
	Data   []TxStatusData `json:"data"`
	Msg    string         `json:"msg"`
}

type TxStatusData struct {
	TxHash   string `json:"txID"`
	TxStatus string `json:"state"`
}
