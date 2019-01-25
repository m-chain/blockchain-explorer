package models

import (
	"blockchain_explorer/models/response"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/shopspring/decimal"
	// "github.com/pingcap/tidb/types"
)

type InfoChain struct {
	Id           int64     `orm:"column(id);pk;auto"json:"id" description:"主键"` // 设置主键
	BlockId      string    `json:"blockid" description:"区块ID"`
	High         string    `json:"high" description:"区块ID(高位)"`
	Low          string    `json:"low" description:"区块ID(低位)"`
	PreviousHash string    `json:"previousBlockHash" description:"上一个区块的hash"`
	CurrentHash  string    `json:"currentBlockHash" description:"区块Hash"`
	UpdateTime   time.Time `orm:"column(updatetime)" json:"updatetime"  description:"更新时间"`
}

func init() {
	orm.RegisterModel(new(InfoChain))
}

func QueryChain(blockId string) InfoChain {
	orm.Debug = true
	var chain InfoChain
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" * ").From("info_chain").Where("block_id = " + blockId).OrderBy("id").Desc().Limit(1)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&chain)
	return chain
}

func InsertOrUpdateChainInfo(chainData response.ChainData) (int64, error) {
	orm.Debug = true
	o := orm.NewOrm()
	infoChain := new(InfoChain)
	var highDec, lowDec, resultDec, lastDec decimal.Decimal
	highStr := strconv.FormatInt(chainData.High, 10)
	logStr := strconv.FormatInt(chainData.Low, 10)
	highDec, _ = decimal.NewFromString(highStr)
	lowDec, _ = decimal.NewFromString(logStr)
	lastDec, _ = decimal.NewFromString("1")
	resultDec = highDec.Add(lowDec)
	resultDec = resultDec.Sub(lastDec)
	infoChain.BlockId = string(resultDec.String())
	resultDec = resultDec.Sub(lastDec)
	nextBlockId := string(resultDec.String())
	infoChain.PreviousHash = chainData.PreviousHash
	infoChain.CurrentHash = chainData.CurrentHash
	infoChain.UpdateTime = time.Now()

	infoChain.Low = logStr
	infoChain.High = highStr
	result, err := o.InsertOrUpdate(infoChain)
	//更新数据库Blockhash
	o.Raw("UPDATE `info_blocks` SET `block_hash` = ? WHERE `block_id` = ? and `block_hash` = ''", chainData.CurrentHash, infoChain.BlockId).Exec()
	o.Raw("UPDATE `info_blocks` SET `next_hash` = ? WHERE `block_id` = ? and `next_hash` = ''", chainData.CurrentHash, nextBlockId).Exec()
	return result, err
}
