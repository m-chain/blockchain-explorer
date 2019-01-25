package models

import (
	"blockchain_explorer/common"
	"blockchain_explorer/models/response"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type InfoBlocks struct {
	Id               int64     `orm:"column(id);pk;auto"json:"id" description:"主键"` // 设置主键
	BlockId          int64     `json:"blockid" description:"区块ID(区块高度)"`
	PreviousHash     string    `json:"previoushash" description:"上一个区块的hash"`
	NextHash         string    `json:"nexthash" description:"下一个区块的hash"`
	DataHash         string    `json:"datahash" description:"区块数据hash"`
	TransactionCount string    `json:"transactioncount" description:"交易笔数(交易次数)"`
	TransactionValue string    `json:"transactionvalue" description:"交易数量"`
	Createtime       time.Time `json:"createtime" description:"创建时间(生成时间)"`
	Updatetime       time.Time `json:"updatetime" description:"更新时间"`
	Content          string    `json:"content" description:"备注"`
	BlockSize        float64   `json:"blocksize" description:"区块数据大小(KB)(区块大小)"`
	BlockHash        string    `json:"blockhash" description:"区块Hash"`
}

func init() {
	orm.RegisterModel(new(InfoBlocks))
}

func QueryBlocks(blockId string) []InfoBlocks {
	orm.Debug = true
	blocks := make([]InfoBlocks, 0)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" * ").From("info_blocks").Where("block_id = " + blockId)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&blocks)
	baseCoin := GetBaseCoin()
	for index := 0; index < len(blocks); index++ {
		var block = &blocks[index]
		block.TransactionValue = common.RealValue(block.TransactionValue, baseCoin.DecimalUnits)
	}

	return blocks
}

func QueryBlocksSearch(search string) (InfoBlocks, error) {
	orm.Debug = true
	var blocks InfoBlocks
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" * ").From("info_blocks").Where("concat(block_id) = '" + search + "' or block_hash = '" + search + "'").Limit(1)
	sql := qb.String()
	o := orm.NewOrm()
	err := o.Raw(sql).QueryRow(&blocks)
	baseCoin := GetBaseCoin()
	blocks.TransactionValue = common.RealValue(blocks.TransactionValue, baseCoin.DecimalUnits)
	return blocks, err
}

func QueryBlocksList(pageNo int, pageSize int, orderBy string) ([]InfoBlocks, int) {
	orm.Debug = true
	blocks := make([]InfoBlocks, 0)
	o := orm.NewOrm()
	pageStart := (pageNo - 1) * pageSize
	var listSize int
	o.Raw("select * from info_blocks  order by id "+orderBy+" limit ?, ? ", pageStart, pageSize).QueryRows(&blocks)
	o.Raw("select count(id) from info_blocks ").QueryRow(&listSize)
	baseCoin := GetBaseCoin()
	for index := 0; index < len(blocks); index++ {
		var block = &blocks[index]
		block.TransactionValue = common.RealValue(block.TransactionValue, baseCoin.DecimalUnits)
	}
	return blocks, listSize
}

func QueryBlockCount() int {
	orm.Debug = true
	var count int
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" count(block_id) ").From("info_blocks").Limit(1)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&count)
	return count
}

func QueryBlockNumber() string {
	orm.Debug = true
	var blocknumber string
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" block_id ").From("info_blocks").OrderBy("block_id desc").Limit(1)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&blocknumber)
	return blocknumber
}

/**
查询数据为空的区块id
**/
func QueryEmptyDataBlockNumber() []InfoBlocks {
	orm.Debug = true
	var blocks []InfoBlocks
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" id,block_id ").From("info_blocks").Where(" block_hash = '' or  block_size = 0 ").OrderBy(" block_id")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&blocks)
	return blocks
}

func InsertOrUpdateBlockInfo(blockData response.BlockData) (int64, error) {
	orm.Debug = true
	o := orm.NewOrm()
	infoBlock := new(InfoBlocks)
	blockid, err := strconv.ParseInt(blockData.Blockid, 10, 64)
	infoBlock.BlockId = blockid
	infoBlock.PreviousHash = blockData.PreviousHash
	infoBlock.DataHash = blockData.DataHash
	infoBlock.TransactionCount = "0"
	infoBlock.Updatetime = time.Now()
	infoBlock.TransactionValue = "0"
	result, err := o.InsertOrUpdate(infoBlock)
	return result, err
}

func UpdateExtBlockInfo(id int64, blockData response.BlockExtData) (int64, error) {
	orm.Debug = true
	o := orm.NewOrm()
	infoBlockExt := new(InfoBlocks)
	infoBlockExt.Id = id
	infoBlockExt.BlockId = blockData.Blockid
	infoBlockExt.BlockSize = blockData.BlockSize
	infoBlockExt.BlockHash = blockData.BlockHash
	infoBlockExt.PreviousHash = blockData.PreviousHash
	blocktime, _ := time.Parse("2006-01-02 15:04:05", blockData.CreateTime)
	infoBlockExt.Createtime = blocktime
	o.Raw("UPDATE `info_blocks` SET `next_hash` = ? WHERE `block_id` = ?", blockData.BlockHash, blockData.Blockid-1).Exec()
	result, err := o.Update(infoBlockExt, "createtime", "blocksize", "blockid", "block_hash")

	return result, err
}
