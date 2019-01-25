package controllers

import (
	"blockchain_explorer/block"
	"blockchain_explorer/constant"
	"blockchain_explorer/models"
	"blockchain_explorer/models/response"
	"blockchain_explorer/redis"
	"encoding/json"
	"fmt"
	"strconv"

	_ "github.com/astaxie/beego"
	"github.com/shopspring/decimal"
)

type APIController struct {
	baseController
}

// @Title getBlocks
// @Description 根据区块id获取区块数据
// @Param   blockId     query    string  true        "区块id"
// @Success 200 {object} []models.InfoBlocks {"result":[object],"message": "","errorcode": "1000","status": true}
// @router /getBlocks [post]
func (con *APIController) GetBlocks() {
	var blocks []models.InfoBlocks
	blockId := con.GetString("blockId")
	blocks = models.QueryBlocks(blockId)
	var resp models.ResponseModel
	if len(blocks) == 0 {
		resp = models.GetResponse(&blocks, false, con.Tr("last_block"), constant.StatusCodeSuccess)
	} else {
		resp = models.GetResponse(&blocks, true, "", constant.StatusCodeSuccess)
	}

	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getBlocksList
// @Description 分页获取区块数据
// @Param   pageNo     query    int  true        "页码"
// @Param   pageSize     query    int  true        "每页条数"
// @Param   orderBy     query    string  false        "排序(asc正序desc倒序)"
// @Success 200 {object} models.InfoBlocks {"result":{ "totalnum": 21,"totalpage": 3,"pageno": 1,"data":[object]},"message": "","errorcode": "1000","status": true}
// @router /getBlocksList [post]
func (con *APIController) GetBlocksList() {
	var blocks []models.InfoBlocks
	pageNo, _ := con.GetInt("pageNo")
	pageSize, _ := con.GetInt("pageSize")
	orderBy := con.GetString("orderBy")
	blocks, blockCount := models.QueryBlocksList(pageNo, pageSize, orderBy)
	var totalPage int
	if blockCount > 0 {
		totalPage = (blockCount + pageSize - 1) / pageSize
	}

	var blockPage response.InfoPageModel
	blockPage.TotalPage = totalPage
	blockPage.TotalNum = blockCount
	blockPage.Data = &blocks
	blockPage.PageNo = pageNo

	var resp models.ResponseModel
	resp = models.GetResponse(&blockPage, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getBlockInfomation
// @Description 区块数据(区块高度,总交易次数,总交易额,Token总数,合约总数)
// @Success 200 {object} models.BlockInfomation {"result":[object],"message": "","errorcode": "1000","status": true}
// @router /getBlockInfomation [post]
func (con *APIController) GetBlockInfomation() {
	var blocks models.BlockInfomation
	blocks = models.QueryBlockInfomation()
	var resp models.ResponseModel
	resp = models.GetResponse(&blocks, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getNodeInfomation
// @Description 获取节点数据
// @Success 200 {object} []models.NodeInformation {"result":[object],"message": "","errorcode": "1000","status": true}
// @router /getNodeInfomation [post]
func (con *APIController) GetNodeInfomation() {
	var node []models.NodeInformation
	node = models.QueryNodeInfomation()
	var resp models.ResponseModel
	resp = models.GetResponse(&node, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getTransactions
// @Description 根据区块id获取交易数据
// @Param   blockId     query    string  true        "区块id"
// @Success 200 {object} []models.InfoTransaction {"result":[object],"message": "","errorcode": "1000","status": true}
// @router /getTransactions [post]
func (con *APIController) GetTransactions() {
	blockID := con.GetString("blockId")
	var transaction []models.InfoTransaction
	transaction = models.QueryTransaction(blockID)
	var resp models.ResponseModel
	resp = models.GetResponse(&transaction, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getWalletDetail
// @Description 根据交易hash获取钱包数据
// @Param   txhash     query    string  true        "交易hash"
// @Success 200 {object} []models.InfoWalletDetail {"result":[object],"message": "","errorcode": "1000","status": true}
// @router /getWalletDetail [post]
func (con *APIController) GetWalletDetail() {
	txhash := con.GetString("txhash")
	var walletDetail []models.InfoWalletDetail
	walletDetail = models.QueryWalletDetail(txhash)
	var resp models.ResponseModel
	resp = models.GetResponse(&walletDetail, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getTokenDetail
// @Description 根据交易hash获取代币详情
// @Param   txhash     query    string  true        "交易hash"
// @Success 200 {object} []models.InfoTokenDetail {"result":[object],"message": "","errorcode": "1000","status": true}
// @router /getTokenDetail [post]
func (con *APIController) GetTokenDetail() {
	txhash := con.GetString("txhash")
	var tokenDetail []models.InfoTokenDetail
	tokenDetail = models.QueryTokenDetail(txhash)
	var resp models.ResponseModel
	resp = models.GetResponse(&tokenDetail, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getTokenDetailWithID
// @Description 根据ID获取代币详情
// @Param   tokenid     query    string  true        "代币id"
// @Success 200 {object} models.InfoTokenDetail {"result":[object],"message": "","errorcode": "1000","status": true}
// @router /getTokenDetailWithID [post]
func (con *APIController) GetTokenDetailWithID() {
	tokenid := con.GetString("tokenid")
	var tokenDetail models.InfoTokenDetail
	tokenstr := redis.GetCacheTokenInfo(tokenid)
	var err error
	if tokenstr != "" {
		err = json.Unmarshal([]byte(tokenstr), &tokenDetail)
		fmt.Println(tokenDetail)
		fmt.Println(err)
	} else {
		//如果不存在则获取区块链数据,并且保存到缓存和数据库中
		blockToken, _ := block.QueryTokenInfo(tokenid)
		models.UpdateTokenDetail(blockToken.Data)
		tokenDetail = models.QueryTokenDetailID(tokenid)
	}
	var resp models.ResponseModel
	resp = models.GetResponse(&tokenDetail, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getTransactionDetail
// @Description 根据交易id获取交易数据
// @Param   txid     query    string  false        "交易id"
// @Success 200 {object} []models.InfoTransactionDetail {"result":[object],"message": "","errorcode": "1000","status": true}
// @router /getTransactionDetail [post]
func (con *APIController) GetTransactionDetail() {
	txid := con.GetString("txid")
	var transactionDetail models.InfoTransactionDetail
	transactionDetail = models.QueryTransactionDetail(txid)
	var resp models.ResponseModel
	resp = models.GetResponse(&transactionDetail, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getTransactionDetailList
// @Description 获取交易详情列表数据
// @Param   pageNo     query    int  true        "页码"
// @Param   pageSize     query    int  true        "每页条数"
// @Param   search     query    string  false        "查询条件"
// @Param   searchType     query    int  false        "查询类型(Empty = 0 BlockId = 1 Address = 2 TokenId = 3, TokenIdAndAddress = 6"
// @Param   address     query    string  false        "地址(当联合查询TokenID和地址时使用)"
// @Param   orderBy     query    string  false        "排序(asc正序desc倒序)"
// @Success 200 {object} []models.InfoTransactionDetail {"result":{ "totalnum": 21,"totalpage": 3,"pageno": 1,"data":[object]},"message": "","errorcode": "1000","status": true}
// @router /getTransactionDetailList [post]
func (con *APIController) GetTransactionDetailList() {
	var transactionDetail []models.InfoTransactionDetail
	pageNo, _ := con.GetInt("pageNo")
	pageSize, _ := con.GetInt("pageSize")
	search := con.GetString("search")
	address := con.GetString("address")
	searchType, _ := con.GetInt("searchType")
	orderBy := con.GetString("orderBy")
	transactionDetail, transactionDetailCount := models.QueryTransactionDetailList(pageNo, pageSize, search, searchType, orderBy, address)
	var transactionPage response.InfoPageModel
	var resp models.ResponseModel
	var totalPage int
	if transactionDetailCount > 0 {
		totalPage = (transactionDetailCount + pageSize - 1) / pageSize
	}

	transactionPage.TotalPage = totalPage
	transactionPage.TotalNum = transactionDetailCount
	transactionPage.Data = &transactionDetail
	transactionPage.PageNo = pageNo
	resp = models.GetResponse(&transactionPage, true, "", constant.StatusCodeSuccess)

	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getTokenDetailList
// @Description 获取代币列表数据
// @Param   pageNo     query    int  true        "页码"
// @Param   pageSize     query    int  true        "每页条数"
// @Param   orderBy     query    string  false        "排序(asc正序desc倒序)"
// @Success 200 {object} []models.InfoTokenDetail {"result":{ "totalnum": 21,"totalpage": 3,"pageno": 1,"data":[object]},"message": "","errorcode": "1000","status": true}
// @router /getTokenDetailList [post]
func (con *APIController) GetTokenDetailList() {
	var tokenDetail []models.InfoTokenDetail
	pageNo, _ := con.GetInt("pageNo")
	pageSize, _ := con.GetInt("pageSize")
	orderBy := con.GetString("orderBy")
	tokenDetail, tokenCount := models.QueryTokenList(pageNo, pageSize, orderBy)
	totalPage := (tokenCount + pageSize - 1) / pageSize

	var tokenPage response.InfoPageModel
	tokenPage.TotalPage = totalPage
	tokenPage.TotalNum = tokenCount
	tokenPage.Data = &tokenDetail
	tokenPage.PageNo = pageNo

	var resp models.ResponseModel
	resp = models.GetResponse(&tokenPage, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getContractDetailWithID
// @Description 根据合约ID获取合约详情
// @Param   contractid     query    string  true        "合约id"
// @Success 200 {object} models.InfoContracts {"result":[object],"message": "","errorcode": "1000","status": true}
// @router /getContractDetailWithID [post]
func (con *APIController) GetContractDetailWithID() {
	contractid := con.GetString("contractid")
	var contract models.InfoContracts
	contract = models.QueryContractWithId(contractid)
	var resp models.ResponseModel
	resp = models.GetResponse(&contract, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getContractDetailWithIDForPld
// @Description 根据合约ID获取合约详情（房产链）
// @Param   contractid     query    string  true        "合约id"
// @Success 200 {object} models.InfoContracts {"result":[object],"message": "","errorcode": "1000","status": true}
// @router /getContractDetailWithIDForPld [post]
func (con *APIController) GetContractDetailWithIDForPld() {
	contractid := con.GetString("contractid")
	var contract models.InfoContracts
	contract = models.QueryContractWithIdForPld(contractid)
	var resp models.ResponseModel
	resp = models.GetResponse(&contract, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getContractDetailList
// @Description 获取合约列表数据
// @Param   pageNo     query    int  true        "页码"
// @Param   pageSize     query    int  true        "每页条数"
// @Param   orderBy     query    string  false        "排序(asc正序desc倒序)"
// @Success 200 {object} []models.InfoContracts {"result":{ "totalnum": 21,"totalpage": 3,"pageno": 1,"data":[object]},"message": "","errorcode": "1000","status": true}
// @router /getContractDetailList [post]
func (con *APIController) GetContractDetailList() {
	var contract []models.InfoContracts
	pageNo, _ := con.GetInt("pageNo")
	pageSize, _ := con.GetInt("pageSize")
	orderBy := con.GetString("orderBy")
	contract = models.QueryContractList(pageNo, pageSize, orderBy)
	contractCount := models.QueryContractCount()
	totalPage := (contractCount + pageSize - 1) / pageSize

	var contractPage response.InfoPageModel
	contractPage.TotalPage = totalPage
	contractPage.TotalNum = contractCount
	contractPage.Data = &contract
	contractPage.PageNo = pageNo

	var resp models.ResponseModel
	resp = models.GetResponse(&contractPage, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getContractGasFeeList
// @Description 获取费用返还列表数据
// @Param   pageNo     query    int  true        "页码"
// @Param   pageSize     query    int  true        "每页条数"
// @Param   id     query    string  true        "合约id""
// @Param   orderBy     query    string  false        "排序(asc正序desc倒序)"
// @Success 200 {object} []models.InfoGasreturn {"result":{ "totalnum": 21,"totalpage": 3,"pageno": 1,"data":[object]},"message": "","errorcode": "1000","status": true}
// @router /getContractGasFeeList [post]
func (con *APIController) GetContractGasFeeList() {
	var gasreturn []models.InfoGasreturn
	pageNo, _ := con.GetInt("pageNo")
	pageSize, _ := con.GetInt("pageSize")
	orderBy := con.GetString("orderBy")
	id := con.GetString("id")
	gasreturn, gasreturnCount := models.QueryGasReturnsList(id, pageNo, pageSize, orderBy)
	totalPage := (gasreturnCount + pageSize - 1) / pageSize

	var gasreturnPage response.InfoPageModel
	gasreturnPage.TotalPage = totalPage
	gasreturnPage.TotalNum = gasreturnCount
	gasreturnPage.Data = &gasreturn
	gasreturnPage.PageNo = pageNo

	var resp models.ResponseModel
	resp = models.GetResponse(&gasreturnPage, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getContractGasRecordList
// @Description 获取费用交易列表数据
// @Param   pageNo     query    int  true        "页码"
// @Param   pageSize     query    int  true        "每页条数"
// @Param   id     query    string  true        "合约id""
// @Param   orderBy     query    string  false        "排序(asc正序desc倒序)"
// @Success 200 {object} []models.InfoGasrecords {"result":{ "totalnum": 21,"totalpage": 3,"pageno": 1,"data":[object]},"message": "","errorcode": "1000","status": true}
// @router /getContractGasRecordList [post]
func (con *APIController) GetContractGasRecordList() {
	var gasrecord []models.InfoGasrecords
	pageNo, _ := con.GetInt("pageNo")
	pageSize, _ := con.GetInt("pageSize")
	orderBy := con.GetString("orderBy")
	id := con.GetString("id")
	gasrecord, gasrecordCount := models.QueryGasRecordsList(id, pageNo, pageSize, orderBy)
	totalPage := (gasrecordCount + pageSize - 1) / pageSize

	var gasrecordPage response.InfoPageModel
	gasrecordPage.TotalPage = totalPage
	gasrecordPage.TotalNum = gasrecordCount
	gasrecordPage.Data = &gasrecord
	gasrecordPage.PageNo = pageNo

	var resp models.ResponseModel
	resp = models.GetResponse(&gasrecordPage, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getContractTxList
// @Description 获取合约下产生的手续费交易
// @Param   pageNo     query    int  true        "页码"
// @Param   pageSize     query    int  true        "每页条数"
// @Param   contractAddress     query    string  true        "合约地址 42210100081711e98188b503d53e9ebb"
// @Param   orderBy     query    string  false        "排序(asc正序desc倒序)"
// @Success 200 {object} []models.InfoTransactionDetail {"result":{ "totalnum": 21,"totalpage": 3,"pageno": 1,"data":[object]},"message": "","errorcode": "1000","status": true}
// @router /getContractTxList [post]
func (con *APIController) GetContractTxList() {
	pageNo, _ := con.GetInt("pageNo")
	pageSize, _ := con.GetInt("pageSize")
	orderBy := con.GetString("orderBy")
	contractAddress := con.GetString("contractAddress")
	gasrecord, gasrecordCount := models.QueryContractTxList(contractAddress, pageNo, pageSize, orderBy)
	totalPage := (gasrecordCount + pageSize - 1) / pageSize

	var gasrecordPage response.InfoPageModel
	gasrecordPage.TotalPage = totalPage
	gasrecordPage.TotalNum = gasrecordCount
	gasrecordPage.Data = &gasrecord
	gasrecordPage.PageNo = pageNo

	var resp models.ResponseModel
	resp = models.GetResponse(&gasrecordPage, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getContractGasRecordWithID
// @Description 根据ID获取费用交易
// @Param   id     query    string  true        "费用记录id"
// @Success 200 {object} models.InfoGasrecords {"result":[object],"message": "","errorcode": "1000","status": true}
// @router /getContractGasRecordWithID [post]
func (con *APIController) GetContractGasRecordWithID() {
	contractid := con.GetString("id")
	var gasrecord models.InfoGasrecords
	gasrecord = models.QueryGasrecordsWithID(contractid)
	var resp models.ResponseModel
	resp = models.GetResponse(&gasrecord, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title searchInformation
// @Description 模糊查询
// @Param   search     query    string  false        "查询条件"
// @Success 200 {object} response.SearchInformationModel {"result":[object],"message": "","errorcode": "1000","status": true}
// @router /searchInformation [post]
func (con *APIController) SearchInformation() {
	search := con.GetString("search")
	var searchInfo response.SearchInformationModel
	var resp models.ResponseModel
	infoBlock, _ := models.QueryBlocksSearch(search)
	if infoBlock.Id > 0 {
		searchInfo.ResultContent = strconv.FormatInt(infoBlock.BlockId, 10)
		searchInfo.ResultType = constant.BlockId
	} else {
		infoContract := models.QueryContractSearch(search)
		if infoContract.Id > 0 {
			searchInfo.ResultContent = strconv.FormatInt(infoContract.Id, 10)
			searchInfo.ResultType = constant.Contract
		} else {
			infoTransactionDetail := models.QueryTransactionDetailSearch(search)
			if infoTransactionDetail.Id > 0 {
				searchInfo.ResultContent = strconv.FormatInt(infoTransactionDetail.Id, 10)
				searchInfo.ResultType = constant.Transaction
			} else {
				infoAddress := models.QueryAddressSearch(search)
				if infoAddress.Id > 0 {
					searchInfo.ResultContent = infoAddress.Address
					searchInfo.ResultType = constant.Address
				} else {
					infoToken := models.QueryTokenSearch(search)
					if infoToken.Id > 0 {
						searchInfo.ResultContent = infoToken.TokenId
						searchInfo.ResultType = constant.TokenId
					}
				}
			}
		}
	}
	if searchInfo.ResultType == 0 {
		resp = models.GetResponse(searchInfo, false, con.Tr("no_found"), constant.StatusCodeSuccess)
	} else {
		resp = models.GetResponse(searchInfo, true, "", constant.StatusCodeSuccess)
	}

	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getQueryBalance
// @Description 根据地址获取地址账户信息
// @Param   address     query    string  true        "地址"
// @Success 200 {object} response.WalletBalanceDetail {"result":[object],"message": "","errorcode": "1000","status": true}
// @router /getQueryBalance [post]
func (con *APIController) GetQueryBalance() {
	address := con.GetString("address")
	var walletBalanceDetail response.WalletBalanceDetail
	var walletBalance response.WalletBalance
	walletBalance, err := block.QueryBalance(address)
	for index := 0; index < len(walletBalance.Data.WalletRests); index++ {
		var walletRest = &walletBalance.Data.WalletRests[index]
		// walletRest.RestNumber = models.GetTokenRealValue(walletRest.RestNumber, walletRest.TokenID)
		// walletRest.FreezeNumber = models.GetTokenRealValue(walletRest.FreezeNumber, walletRest.TokenID)
		//计算大数字字符串
		var freezeDec, restDec, resultDec decimal.Decimal
		freezeDec, _ = decimal.NewFromString(walletRest.FreezeNumber)
		restDec, _ = decimal.NewFromString(walletRest.RestNumber)
		resultDec = freezeDec.Add(restDec)
		fmt.Print(string(resultDec.String()))
		walletRest.RestNumber = models.GetTokenRealValue(string(resultDec.String()), walletRest.TokenID)
		walletRest.FreezeNumber = models.GetTokenRealValue(walletRest.FreezeNumber, walletRest.TokenID)
	}
	walletBalanceDetail = walletBalance.Data
	var resp models.ResponseModel
	if err == nil {
		resp = models.GetResponse(&walletBalanceDetail, true, "", constant.StatusCodeSuccess)
	} else {
		resp = models.GetResponse(&walletBalanceDetail, false, con.Tr("no_address"), constant.StatusCodeSuccess)
	}

	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getBaseCoin
// @Description 获取基础币种
// @Success 200 {object} models.InfoTokenDetail {"result":[object],"message": "","errorcode": "1000","status": true}
// @router /getBaseCoin [post]
func (con *APIController) GetBaseCoin() {
	var tokenDetial models.InfoTokenDetail
	tokenDetial = models.GetBaseCoin()
	var resp models.ResponseModel
	resp = models.GetResponse(&tokenDetial, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title GetGasAddress
// @Description 获取手续费收取地址
// @Success 200 string
// @router /GetGasAddress [post]
func (con *APIController) GetGasAddress() {
	addressDetail, _ := block.GetGasAddressInfo()
	address, _ := addressDetail.Data.(string)
	var resp models.ResponseModel
	resp = models.GetResponse(address, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}

// @Title getTransactionDetailWithHash
// @Description 根据交易hash获取交易数据
// @Param   hash     query    string  true        "交易hash"
// @Success 200 {object} models.InfoTransactionDetail {"result":[object],"message": "","errorcode": "1000","status": true}
// @router /getTransactionDetailWithHash [post]
func (con *APIController) GetTransactionDetailWithHash() {
	hash := con.GetString("txhash")
	var resp models.ResponseModel
	infoTransactionDetail := models.QueryTransactionDetailSearch(hash)
	resp = models.GetResponse(infoTransactionDetail, true, "", constant.StatusCodeSuccess)
	con.Data["json"] = &resp
	con.ServeJSON()
}
