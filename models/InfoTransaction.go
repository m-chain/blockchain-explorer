package models

import (
	"blockchain_explorer/models/response"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type InfoTransaction struct {
	Id               int64     `orm:"column(id);pk;auto"json:"id"  description:"主键"` // 设置主键
	Blockid          string    `json:"blockid"  description:"区块id"`
	Txhash           string    `json:"txhash"  description:"交易hash"`
	Channelname      string    `json:"channelname"  description:"通道名称"`
	Chaincodename    string    `json:"chaincodename"  description:"链码名称"`
	Chaincodeversion string    `json:"chaincodeversion"  description:"链码版本号"`
	Createdt         time.Time `json:"createdt"  description:"创建时间"`
}

func init() {
	orm.RegisterModel(new(InfoTransaction))
}

func QueryTransaction(blockId string) []InfoTransaction {
	orm.Debug = true
	transaction := make([]InfoTransaction, 0)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" * ").From("info_transaction").Where("blockid = " + blockId)

	sql := qb.String()

	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&transaction)
	fmt.Println(transaction)
	return transaction
}

func InsertOrUpdateBlockTransaction(blockid string, blockTransaction []response.BlockTransactions) error {
	orm.Debug = true
	o := orm.NewOrm()
	err := o.Begin()
	for index := 0; index < len(blockTransaction); index++ {
		transaction := blockTransaction[index]
		infoTransaction := new(InfoTransaction)
		infoTransaction.Blockid = blockid
		infoTransaction.Chaincodename = transaction.Chaincodename
		infoTransaction.Chaincodeversion = transaction.Chaincodeversion
		infoTransaction.Channelname = transaction.Channelname
		infoTransaction.Createdt = time.Now()
		infoTransaction.Txhash = transaction.Txhash
		_, err = o.InsertOrUpdate(infoTransaction)
		if err != nil {
			err = o.Rollback()
			return err
		}
	}
	err = o.Commit()
	return err
}
