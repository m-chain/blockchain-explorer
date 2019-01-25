package response

type BlockNumber struct {
	Status      bool   `json:"status"`
	Msg         string `json:"msg"`
	BlockNumber string `json:"blockNumber"`
}

type BlockTransactionCountByNumber struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Count  int64  `json:"count"`
}

type BlockTransactionCountByHash struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Count  int64  `json:"count"`
}

type BaseReponse struct {
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type BaseNodeReponse struct {
	Code int16       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
