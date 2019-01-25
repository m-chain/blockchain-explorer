package models

import (
	"blockchain_explorer/common"
	"blockchain_explorer/constant"
	"blockchain_explorer/models/response"
	"fmt"

	"github.com/astaxie/beego/orm"
)

type InfoTransactionDetail struct {
	Id              int64  `orm:"column(id);pk;auto"json:"id" description:"主键"` // 设置主键
	Blockid         string `json:"blockid" description:"区块id(区块高度)"`
	ParentTxhash    string `orm:"column(parenttxhash)" json:"parenttxhash" description:"区块总交易hash(非单笔交易hash)"`
	Txhash          string `orm:"column(txhash)" json:"txhash" description:"交易hash"`
	TokenId         string `json:"tokenid" description:"币种ID"`
	Tokensymbol     string `orm:"-" json:"tokensymbol" description:"币种单位"`
	FromAddress     string `json:"fromaddress" description:"转出地址(发起者)"`
	ToAddress       string `json:"toaddress" description:"转入地址(接收者)`
	Number          string `json:"number" description:"交易数量"`
	Fee             string `json:"fee" description:"手续费"`
	IsCost          int8   `json:"iscost" description:"是否手续费交易(1是0否)"`
	TransactionTime string `json:"transactiontime" description:"交易时间"`
	Nonce           string `json:"nonce" description:"交易下标"`
	State           int    `json:"state" description:"转账状态:0.待确认 1.成功 2.失败"`
	Msg             string `json:"msg" description:"转账错误信息"`
	Notes           string `json:"notes" description:"备注"`
	ContractAddress string `json:"contractaddress" description:"产生该笔交易所对应的合约"`
}

type InfoTransactionInfo struct {
	TransactionTime  int64
	TransactionTotal string
}

type InfoTransactionExt struct {
	TransactionCount string
	TransactionTotal string
}

func init() {
	orm.RegisterModel(new(InfoTransactionDetail))
}

func QueryTransactionDetail(txid string) InfoTransactionDetail {
	filter := "id = '" + txid + "'"
	return queryTransactionDetailWithFilter(filter)
}

func QueryTransactionDetailSearch(search string) InfoTransactionDetail {
	filter := "txhash = '" + search + "' and is_cost = 0 "
	return queryTransactionDetailWithFilter(filter)
}

func QueryTransactionDetailWait() []InfoTransactionDetail {
	orm.Debug = true
	var transaction []InfoTransactionDetail
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" id,txhash ").From("info_transaction_detail").Where("state = '0' and is_cost ='0' limit 100")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&transaction)
	// baseCoin := GetBaseCoin()
	// for index := 0; index < len(transaction); index++ {
	// 	transactionDetail := transaction[index]
	// 	infoToken := QueryTokenDetailID(transactionDetail.TokenId)
	// 	transactionDetail.Tokensymbol = infoToken.TokenSymbol
	// 	transactionDetail.Number = common.RealValue(transactionDetail.Number, infoToken.DecimalUnits)
	// 	transactionDetail.Fee = common.RealValue(transactionDetail.Fee, baseCoin.DecimalUnits)
	// 	fmt.Println(transactionDetail)
	// }
	return transaction
}

func queryTransactionDetailWithFilter(filter string) InfoTransactionDetail {
	orm.Debug = true
	var transactionDetail InfoTransactionDetail
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" * ").From("info_transaction_detail").Where(filter).Limit(1)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&transactionDetail)
	baseCoin := GetBaseCoin()
	infoToken := QueryTokenDetailID(transactionDetail.TokenId)
	transactionDetail.Tokensymbol = infoToken.TokenSymbol
	transactionDetail.Number = common.RealValue(transactionDetail.Number, infoToken.DecimalUnits)
	transactionDetail.Fee = common.RealValue(transactionDetail.Fee, baseCoin.DecimalUnits)
	fmt.Println(transactionDetail)
	return transactionDetail
}

func QueryTransactionInfo() InfoTransactionInfo {
	orm.Debug = true
	var infoTransactionInfo InfoTransactionInfo
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" sum(transaction_count) as transaction_time,sum(transaction_value) as transaction_total").From("info_blocks").Limit(1)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&infoTransactionInfo)
	baseCoin := GetBaseCoin()
	infoTransactionInfo.TransactionTotal = common.RealValue(infoTransactionInfo.TransactionTotal, baseCoin.DecimalUnits)
	fmt.Println(infoTransactionInfo)
	return infoTransactionInfo
}

func QueryTransactionDetailList(pageNo int, pageSize int, search string, searchType int, orderBy string, address string) ([]InfoTransactionDetail, int) {
	orm.Debug = true
	transaction := make([]InfoTransactionDetail, 0)
	o := orm.NewOrm()
	pageStart := (pageNo - 1) * pageSize
	var listSize int
	switch searchType {
	case constant.BlockId:
		o.Raw("select * from info_transaction_detail where blockid = ? and fee <> 0 order by id "+orderBy+" limit ?, ? ", search, pageStart, pageSize).QueryRows(&transaction)
		o.Raw("select count(id) from info_transaction_detail where blockid = ? and fee <> 0 ", search).QueryRow(&listSize)
		break
	case constant.Address:
		o.Raw("select * from info_transaction_detail where (from_address = ? or to_address = ?)   order by id "+orderBy+" limit ?, ? ", search, search, pageStart, pageSize).QueryRows(&transaction)
		o.Raw("select count(id) from info_transaction_detail where (from_address = ? or to_address = ?)    ", search, search).QueryRow(&listSize)
		break
	case constant.TokenId:
		o.Raw("select * from info_transaction_detail where token_id = ? and fee <> 0 order by id "+orderBy+" limit ?, ? ", search, pageStart, pageSize).QueryRows(&transaction)
		o.Raw("select count(id) from info_transaction_detail where token_id = ? and fee <> 0 ", search).QueryRow(&listSize)
		break
	case constant.TokenIdAndAddress:
		o.Raw("select * from info_transaction_detail where token_id = ? and (from_address = ? or to_address = ?)  order by id "+orderBy+" limit ?, ? ", search, address, address, pageStart, pageSize).QueryRows(&transaction)
		o.Raw("select count(id) from info_transaction_detail where token_id = ? and (from_address = ? or to_address = ?)  ", search, address, address).QueryRow(&listSize)
		break
	default:
		o.Raw("select * from info_transaction_detail where fee <> 0 order by id "+orderBy+" limit ?, ? ", pageStart, pageSize).QueryRows(&transaction)
		o.Raw("select count(id) from info_transaction_detail where fee <> 0 ").QueryRow(&listSize)
		break
	}
	baseCoin := GetBaseCoin()
	for index := 0; index < len(transaction); index++ {
		var transaction = &transaction[index]
		infoToken := QueryTokenDetailID(transaction.TokenId)
		transaction.Tokensymbol = infoToken.TokenSymbol
		transaction.Number = common.RealValue(transaction.Number, infoToken.DecimalUnits)
		transaction.Fee = common.RealValue(transaction.Fee, baseCoin.DecimalUnits)
	}
	return transaction, listSize
}

func InsertOrUpdateBlockTransactionDetail(blockid string, transactionHash string, blockTrans []response.BlockTrans) error {
	orm.Debug = true
	o := orm.NewOrm()
	err := o.Begin()
	for index := 0; index < len(blockTrans); index++ {
		transaction := blockTrans[index]
		infoTransactionDetail := new(InfoTransactionDetail)
		infoTransactionDetail.Blockid = blockid
		infoTransactionDetail.FromAddress = transaction.FromAddress
		infoTransactionDetail.ToAddress = transaction.ToAddress
		infoTransactionDetail.TokenId = transaction.TokenID
		infoTransactionDetail.Txhash = transaction.TxId
		infoTransactionDetail.ParentTxhash = transactionHash
		infoTransactionDetail.State = transaction.State
		infoTransactionDetail.Number = transaction.Number
		infoTransactionDetail.Nonce = transaction.Nonce
		infoTransactionDetail.Msg = transaction.Msg
		infoTransactionDetail.Fee = transaction.GasUsed
		infoTransactionDetail.TransactionTime = transaction.Time
		infoTransactionDetail.Notes = transaction.Notes
		infoTransactionDetail.ContractAddress = transaction.Reason

		if transaction.GasUsed != "0" {
			infoTransactionDetail.IsCost = 0
		} else {
			infoTransactionDetail.IsCost = 1
		}
		_, err = o.InsertOrUpdate(infoTransactionDetail)
		if err != nil {
			err = o.Rollback()
			return err
		}

	}

	UpdateTransactionCount(blockid, o)
	err = o.Commit()
	return err
}

func UpdateTransactionCount(blockid string, o orm.Ormer) error {
	orm.Debug = true
	var transactionCount, transactionTotal string
	basecoin := GetBaseCoin()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" count(id) as transaction_count").From("info_transaction_detail").Where(" blockid = " + blockid + " and is_cost = 0 ").Limit(1)
	sql := qb.String()
	err := o.Raw(sql).QueryRow(&transactionCount)

	qbtotal, _ := orm.NewQueryBuilder("mysql")
	qbtotal.Select(" sum(number) as transaction_total").From("info_transaction_detail").Where(" blockid = " + blockid + " and is_cost = 0 and token_id = '" + basecoin.TokenId + "'").Limit(1)
	sql = qbtotal.String()
	err = o.Raw(sql).QueryRow(&transactionTotal)
	// 更新数据库Blockhash
	if transactionCount == "" {
		transactionCount = "0"
	}
	if transactionTotal == "" {
		transactionTotal = "0"
	}

	o.Raw("UPDATE `info_blocks` SET `transaction_count` = ? , `transaction_value` = ? WHERE `block_id` = ? ", transactionCount, transactionTotal, blockid).Exec()
	err = o.Commit()
	return err
}

func UpdateTransactionStatus(transactionInfo response.TxStatusData) error {
	orm.Debug = true
	o := orm.NewOrm()
	o.Begin()
	// infoTransactionDetail := new(InfoTransactionDetail)
	// infoTransactionDetail.Txhash = transactionInfo.TxHash
	// infoTransactionDetail.State, _ = strconv.Atoi(transactionInfo.TxStatus)
	// var transSuccess []string
	// var transFail []string
	// for index := 0; index < len(transactionInfo.Data); index++ {
	// 	transactionDetail := transactionInfo.Data[index]
	// 	if transactionDetail.TxStatus == "1" {
	// 		transSuccess = append(transSuccess, "'"+transactionDetail.TxHash+"'")
	// 	}
	// 	if transactionDetail.TxStatus == "2" {
	// 		transFail = append(transFail, "'"+transactionDetail.TxHash+"'")
	// 	}
	// }
	// transSuccessStr := strings.Join(transSuccess, ",")
	// transFailStr := strings.Join(transFail, ",")
	// if transSuccessStr != "" {
	var id string
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" id ").From("info_transaction_detail").Where("txhash = '" + transactionInfo.TxHash + "'").ForUpdate()
	sql := qb.String()
	o.Raw(sql).QueryRow(&id)
	result, err := o.Raw("UPDATE `info_transaction_detail` SET `state` = ?  WHERE `txhash` =  ?", transactionInfo.TxStatus, transactionInfo.TxHash).Exec()
	if err != nil {
		fmt.Printf("%s", result.RowsAffected)
	}
	// }
	// if transFailStr != "" {
	// 	o.Raw("UPDATE `info_transaction_detail` SET `state` = 2  WHERE `txhash` in (" + transFailStr + ") ").Exec()
	// }
	// o.Update(&infoTransactionDetail, "state")
	o.Commit()
	return err
}
