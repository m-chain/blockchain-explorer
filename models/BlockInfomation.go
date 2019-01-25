package models

import (
	"blockchain_explorer/redis"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type BlockInfomation struct {
	Id               int64     `orm:"column(id);pk;"json:"id" description:"主键"` // 设置主键
	BlockHeight      int64     `json:"blockheight" description:"区块高度"`
	TransactionTimes int64     `json:"transactiontimes" description:"总交易次数"`
	TransactionValue float64   `json:"transactionvalue" description:"总交易额"`
	TokenNum         int16     `json:"tokennum"  description:"token总数"`
	ContractNum      int       `json:"contractnum"  description:"合约总数"`
	UpdateTime       time.Time `json:"updatetime"  description:"更新时间"`
}

func init() {
	orm.RegisterModel(new(BlockInfomation))
}

func QueryBlockInfomation() BlockInfomation {
	//判断是否有redis
	var blocks BlockInfomation
	blockstr := redis.GetCacheBlockInfo()
	var err error
	if blockstr != "" {
		err = json.Unmarshal([]byte(blockstr), &blocks)
		fmt.Println(blocks)
		fmt.Println(err)
	}

	if blocks.Id <= 0 || blockstr == "" || err != nil {
		orm.Debug = true
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select(" * ").From("block_infomation").Limit(1)
		sql := qb.String()
		o := orm.NewOrm()
		o.Raw(sql).QueryRow(&blocks)
		redis.SaveCacheBlockInfo(blocks)
	}

	fmt.Println(blocks)
	return blocks
}

func UpdateBlockInfomation() (int64, error) {
	orm.Debug = true
	o := orm.NewOrm()
	infoBlockInfomation := new(BlockInfomation)
	var blocknumber = QueryBlockNumber()
	var txtimes = QueryTransactionInfo()
	infoBlockInfomation.Id = 1
	infoBlockInfomation.BlockHeight, _ = strconv.ParseInt(blocknumber, 10, 64)
	infoBlockInfomation.TransactionTimes = txtimes.TransactionTime
	infoBlockInfomation.TransactionValue, _ = strconv.ParseFloat(txtimes.TransactionTotal, 64)
	infoBlockInfomation.ContractNum = QueryContractCount()
	infoBlockInfomation.TokenNum = QueryTokenInfo().TokenNum
	infoBlockInfomation.UpdateTime = time.Now()
	result, err := o.InsertOrUpdate(infoBlockInfomation)
	redis.SaveCacheBlockInfo(infoBlockInfomation)
	return result, err
}
