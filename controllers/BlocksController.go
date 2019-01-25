package controllers

import (
	"blockchain_explorer/block"
	"blockchain_explorer/constant"
	"blockchain_explorer/models"
	"fmt"

	_ "github.com/astaxie/beego"
)

type BlocksController struct {
	baseController
}

//================================= 调用接口直接获取链上数据 =========================//
func (con *BlocksController) GetInfo() {
	var resp = block.GetInfo()
	con.Data["json"] = &resp
	con.ServeJSON()
}

func (con *BlocksController) GetBlockTransactionCountByNumber() {
	blocknum := con.GetString("blocknum")
	var resp = block.GetBlockTransactionCountByNumber(blocknum)
	con.Data["json"] = &resp
	con.ServeJSON()
}

func (con *BlocksController) QueryTokenInfo() {
	tokenid := con.GetString("tokenid")
	var resp = block.GetBlockTransactionCountByNumber(tokenid)
	con.Data["json"] = &resp
	con.ServeJSON()
}

func (con *BlocksController) BlocksDetail() {
	blockid := con.GetString("blockid")
	var data, err = block.BlocksDetail(blockid)
	var resp models.ResponseModel
	if err != nil {
		fmt.Println(err)
		resp = models.GetResponse(nil, false, "", constant.StatusCodeFail)
	} else {
		if data.Status == false {
			resp = models.GetResponse(data.Data, false, data.Msg, constant.StatusCodeFail)
		} else {
			resp = models.GetResponse(data.Data, true, data.Msg, constant.StatusCodeSuccess)
		}
	}

	con.Data["json"] = &resp
	con.ServeJSON()
}

func (con *BlocksController) BlockTransactionCountByHash() {
	blockhash := con.GetString("blockhash")
	var resp = block.BlockTransactionCountByHash(blockhash)
	con.Data["json"] = &resp
	con.ServeJSON()
}

func (con *BlocksController) TransactionByHash() {
	hash := con.GetString("hash")
	var resp = block.TransactionByHash(hash)
	con.Data["json"] = &resp
	con.ServeJSON()
}

func (con *BlocksController) ChainInfo() {
	var resp = block.ChainInfo()
	con.Data["json"] = &resp
	con.ServeJSON()
}

func (con *BlocksController) TransactionByBlockHashAndIndex() {
	blockHash := con.GetString("blockHash")
	index := con.GetString("index")
	var resp = block.TransactionByBlockHashAndIndex(blockHash, index)
	con.Data["json"] = &resp
	con.ServeJSON()
}

func (con *BlocksController) TransactionByBlockNumberAndIndex() {
	blockId := con.GetString("blockId")
	index := con.GetString("index")
	var resp = block.TransactionByBlockNumberAndIndex(blockId, index)
	con.Data["json"] = &resp
	con.ServeJSON()
}

func (con *BlocksController) QueryTxsByAddress() {
	address := con.GetString("address")
	var resp = block.QueryTxsByAddress(address)
	con.Data["json"] = &resp
	con.ServeJSON()
}

func (con *BlocksController) QueryMultiTxInfo() {
	txId := con.GetString("txId")
	var resp = block.QueryMultiTxInfo(txId)
	con.Data["json"] = &resp
	con.ServeJSON()
}

func (con *BlocksController) GetBlockNumber() {
	var resp = block.GetBlockNumber()
	con.Data["json"] = &resp
	con.ServeJSON()
}

func (con *BlocksController) QueryTransferInfos() {
	address := con.GetString("address")
	var resp = block.QueryTransferInfos(address)
	con.Data["json"] = &resp
	con.ServeJSON()
}

func (con *BlocksController) QueryTransferInfo() {
	txId := con.GetString("txId")
	var resp = block.QueryTransferInfo(txId)
	con.Data["json"] = &resp
	con.ServeJSON()
}

func (con *BlocksController) QueryTransferCount() {
	address := con.GetString("address")
	var resp = block.QueryTransferCount(address)
	con.Data["json"] = &resp
	con.ServeJSON()
}
