package block

import (
	"blockchain_explorer/constant"
	"blockchain_explorer/models"
	"blockchain_explorer/models/response"
	"blockchain_explorer/service"
	"encoding/json"
	"fmt"
	"net/url"

	_ "github.com/astaxie/beego"
)

//================================= 调用接口直接获取链上数据 =========================//
func GetInfo() models.ResponseModel {
	body, err := service.HttpGet(constant.BlockChainURL + constant.URLGetBlockNumber)
	var resp models.ResponseModel
	var dat response.BlockNumber
	if err != nil {
		fmt.Println(err)
		resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
		if err != nil {
			fmt.Println(err)
			resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
		} else {
			if dat.Status == false {
				resp = models.GetResponse(dat.BlockNumber, false, dat.Msg, constant.StatusCodeFail)
			} else {
				resp = models.GetResponse(dat.BlockNumber, true, dat.Msg, constant.StatusCodeSuccess)
			}
		}
	}
	return resp
}

func GetBlockTransactionCountByNumber(blocknum string) models.ResponseModel {
	body, err := service.HttpGet(constant.BlockChainURL + constant.URLGetBlockTransactionCountByNumber + blocknum)
	var resp models.ResponseModel
	var dat response.BlockTransactionCountByNumber
	if err != nil {
		fmt.Println(err)
		resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
		if err != nil {
			fmt.Println(err)
			resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
		} else {
			if dat.Status == false {
				resp = models.GetResponse(dat.Count, false, dat.Msg, constant.StatusCodeFail)
			} else {
				resp = models.GetResponse(dat.Count, true, dat.Msg, constant.StatusCodeSuccess)
			}
		}
	}

	return resp
}

func QueryTokenInfo(tokenid string) (response.BlockTokenInfo, error) {
	body, err := service.HttpGet(constant.BlockChainURL + constant.URLQueryTokenInfo + tokenid)
	var dat response.BlockTokenInfo
	if err != nil {
		fmt.Println(err)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
	}
	return dat, err
}

func QueryBalance(address string) (response.WalletBalance, error) {
	body, err := service.HttpGet(constant.BlockChainURL + constant.URLQueryBalance + address)
	var dat response.WalletBalance
	if err != nil {
		fmt.Println(err)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
	}
	return dat, err
}

func BlocksDetail(blockid string) (response.BlockInfo, error) {
	body, err := service.HttpGet(constant.BlockChainURL + constant.URLBlocksDetail + blockid)
	var dat response.BlockInfo
	if err != nil {
		fmt.Println(err)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
	}

	return dat, err
}

func BlocksExtDetail(blockid string) (response.BlockExtInfo, error) {
	body, err := service.HttpGet(constant.BlockChainURL + constant.URLBlocksExtDetail + blockid)
	var dat response.BlockExtInfo
	if err != nil {
		fmt.Println(err)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
	}
	return dat, err
}

func BlockTransactionCountByHash(blockhash string) models.ResponseModel {
	var param = url.Values{}
	param.Add("blockHash", blockhash)
	body, err := service.HttpPost(constant.BlockChainURL+constant.URLBlockTransactionCountByHash, param)
	var resp models.ResponseModel
	var dat response.BlockTransactionCountByHash
	if err != nil {
		fmt.Println(err)
		resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
		if err != nil {
			fmt.Println(err)
			resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
		} else {
			if dat.Status == false {
				resp = models.GetResponse(dat.Count, false, dat.Msg, constant.StatusCodeFail)
			} else {
				resp = models.GetResponse(dat.Count, true, dat.Msg, constant.StatusCodeSuccess)
			}
		}
	}

	return resp
}

func TransactionByHash(hash string) models.ResponseModel {
	var param = url.Values{}
	param.Add("txId", hash)
	body, err := service.HttpPost(constant.BlockChainURL+constant.URLTransactionByHash, param)
	var resp models.ResponseModel
	var dat response.BaseReponse
	if err != nil {
		fmt.Println(err)
		resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
		if err != nil {
			fmt.Println(err)
			resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
		} else {
			if dat.Status == false {
				resp = models.GetResponse(nil, false, dat.Msg, constant.StatusCodeFail)
			} else {
				resp = models.GetResponse(dat, true, dat.Msg, constant.StatusCodeSuccess)
			}
		}
	}

	return resp
}

func ChainInfo() models.ResponseModel {
	body, err := service.HttpGet(constant.BlockChainURL + constant.URLChainInfo)
	var resp models.ResponseModel
	var dat response.BaseReponse
	if err != nil {
		fmt.Println(err)
		resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
		if err != nil {
			fmt.Println(err)
			resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
		} else {
			if dat.Status == false {
				resp = models.GetResponse(nil, false, dat.Msg, constant.StatusCodeFail)
			} else {
				resp = models.GetResponse(dat.Data, true, dat.Msg, constant.StatusCodeSuccess)
			}
		}
	}

	return resp
}

func TransactionByBlockHashAndIndex(blockHash string, index string) models.ResponseModel {
	var param = url.Values{}
	param.Add("blockHash", blockHash)
	param.Add("index", index)
	body, err := service.HttpPost(constant.BlockChainURL+constant.URLTransactionByBlockHashAndIndex, param)
	var resp models.ResponseModel
	var dat response.BaseReponse
	if err != nil {
		fmt.Println(err)
		resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
		if err != nil {
			fmt.Println(err)
			resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
		} else {
			if dat.Status == false {
				resp = models.GetResponse(nil, false, dat.Msg, constant.StatusCodeFail)
			} else {
				resp = models.GetResponse(dat.Data, true, dat.Msg, constant.StatusCodeSuccess)
			}
		}
	}

	return resp
}

func TransactionByBlockNumberAndIndex(blockId string, index string) models.ResponseModel {
	var param = url.Values{}
	param.Add("blockId", blockId)
	param.Add("index", index)
	body, err := service.HttpPost(constant.BlockChainURL+constant.URLTransactionByBlockNumberAndIndex, param)
	var resp models.ResponseModel
	var dat response.BaseReponse
	if err != nil {
		fmt.Println(err)
		resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
		if err != nil {
			fmt.Println(err)
			resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
		} else {
			if dat.Status == false {
				resp = models.GetResponse(nil, false, dat.Msg, constant.StatusCodeFail)
			} else {
				resp = models.GetResponse(dat.Data, true, dat.Msg, constant.StatusCodeSuccess)
			}
		}
	}

	return resp
}

func QueryTxsByAddress(address string) models.ResponseModel {
	body, err := service.HttpGet(constant.BlockChainURL + constant.URLQueryTxsByAddress + address)
	var resp models.ResponseModel
	var dat response.BaseReponse
	if err != nil {
		fmt.Println(err)
		resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
		if err != nil {
			fmt.Println(err)
			resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
		} else {
			if dat.Status == false {
				resp = models.GetResponse(nil, false, dat.Msg, constant.StatusCodeFail)
			} else {
				resp = models.GetResponse(dat.Data, true, dat.Msg, constant.StatusCodeSuccess)
			}
		}
	}

	return resp
}

func QueryMultiTxInfo(txId string) models.ResponseModel {
	var param = url.Values{}
	param.Add("txId", txId)
	body, err := service.HttpPost(constant.BlockChainURL+constant.URLQueryMultiTxInfo, param)
	var resp models.ResponseModel
	var dat response.BaseReponse
	if err != nil {
		fmt.Println(err)
		resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
		if err != nil {
			fmt.Println(err)
			resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
		} else {
			if dat.Status == false {
				resp = models.GetResponse(nil, false, dat.Msg, constant.StatusCodeFail)
			} else {
				resp = models.GetResponse(dat.Data, true, dat.Msg, constant.StatusCodeSuccess)
			}
		}
	}

	return resp
}

func GetBlockNumber() models.ResponseModel {
	body, err := service.HttpGet(constant.BlockChainURL + constant.URLGetBlockNumber)
	var resp models.ResponseModel
	var dat response.BlockNumber
	if err != nil {
		fmt.Println(err)
		resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
		if err != nil {
			fmt.Println(err)
			resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
		} else {
			if dat.Status == false {
				resp = models.GetResponse(nil, false, dat.Msg, constant.StatusCodeFail)
			} else {
				resp = models.GetResponse(dat.BlockNumber, true, dat.Msg, constant.StatusCodeSuccess)
			}
		}
	}

	return resp
}

func QueryTransferInfos(address string) models.ResponseModel {
	body, err := service.HttpGet(constant.BlockChainURL + constant.URLQueryTransferInfos + address)
	var resp models.ResponseModel
	var dat response.BaseReponse
	if err != nil {
		fmt.Println(err)
		resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
		if err != nil {
			fmt.Println(err)
			resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
		} else {
			if dat.Status == false {
				resp = models.GetResponse(nil, false, dat.Msg, constant.StatusCodeFail)
			} else {
				resp = models.GetResponse(dat.Data, true, dat.Msg, constant.StatusCodeSuccess)
			}
		}
	}

	return resp
}

func QueryTransferInfo(txId string) models.ResponseModel {
	var param = url.Values{}
	param.Add("txId", txId)
	body, err := service.HttpPost(constant.BlockChainURL+constant.URLQueryTransferInfo, param)
	var resp models.ResponseModel
	var dat response.BaseReponse
	if err != nil {
		fmt.Println(err)
		resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
		if err != nil {
			fmt.Println(err)
			resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
		} else {
			if dat.Status == false {
				resp = models.GetResponse(nil, false, dat.Msg, constant.StatusCodeFail)
			} else {
				resp = models.GetResponse(dat.Data, true, dat.Msg, constant.StatusCodeSuccess)
			}
		}
	}

	return resp
}

func QueryTransferCount(address string) models.ResponseModel {
	body, err := service.HttpGet(constant.BlockChainURL + constant.URLQueryTransferCount + address)
	var resp models.ResponseModel
	var dat response.BlockTransactionCountByNumber
	if err != nil {
		fmt.Println(err)
		resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
		if err != nil {
			fmt.Println(err)
			resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
		} else {
			if dat.Status == false {
				resp = models.GetResponse(nil, false, dat.Msg, constant.StatusCodeFail)
			} else {
				resp = models.GetResponse(dat.Count, true, dat.Msg, constant.StatusCodeSuccess)
			}
		}
	}

	return resp
}

func GetChainInfo() (response.ChainInfoData, error) {
	body, err := service.HttpGet(constant.BlockChainURL + constant.URLGetChainInfo)
	var dat response.ChainInfoData
	if err != nil {
		fmt.Println(err)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
	}
	return dat, err
}

func GetNodePeerInfo() (response.NodeInfo, error) {
	body, err := service.HttpGet(constant.NodePeerURL + constant.URLNodePeerInfo)
	var dat response.NodeInfo
	if err != nil {
		fmt.Println(err)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
	}
	return dat, err
}

func GetGasRecordsInfo(lastfeeid string) (response.Gasrecords, error) {
	body, err := service.HttpGet(constant.BlockChainURL + constant.URLQueryGasRecordsInfo + lastfeeid + "/100")
	var dat response.Gasrecords
	if err != nil {
		fmt.Println(err)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
	}
	return dat, err
}

func GetGasAddressInfo() (response.BaseReponse, error) {
	body, err := service.HttpGet(constant.BlockChainURL + constant.URLGetGasAddressInfo)
	var dat response.BaseReponse
	if err != nil {
		fmt.Println(err)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
	}
	return dat, err
}

func GetQueryTransferStatus(txid string) (response.BaseReponse, error) {
	body, err := service.HttpGet(constant.BlockChainURL + constant.URLQueryTransferStatus + txid)
	var dat response.BaseReponse
	if err != nil {
		fmt.Println(err)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
	}
	return dat, err
}

/**
* 批量获取交易状态
 */
func GetQueryMultiTransferStatus(data string) (response.TxStatusInfoData, error) {
	var param = url.Values{}
	param.Add("data", data)
	body, err := service.HttpPost(constant.BlockChainURL+constant.URLQueryMultiTransferStatus, param)
	var dat response.TxStatusInfoData
	if err != nil {
		fmt.Println(err)
	} else {
		err = json.Unmarshal([]byte(body), &dat)
	}
	return dat, err
}
