package models

import (
	"blockchain_explorer/common"
	"blockchain_explorer/models/response"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type InfoGasrecords struct {
	Id              int64     `orm:"column(id);pk;auto"json:"id" description:"主键"` // 设置主键
	Feeid           int64     `json:"feeid" description:"手续费序号"`
	Ccname          string    `json:"ccname" description:"合约名称"`
	Ccversion       string    `json:"ccversion" description:"合约版本"`
	ContractAddress string    `json:"contractaddress" description:"合约地址"`
	Fcnname         string    `json:"fcnname" description:"合约调用的参数"`
	Args            string    `json:"args" description:"交易数据"`
	Address         string    `json:"address" description:"地址"`
	Gasused         string    `json:"gasused" description:"所需的手续费"`
	Txid            string    `json:"txid" description:"执行合约方法时产生的交易ID"`
	ReturnTxid      string    `json:"returntxid" description:"对应返还记录的交易ID(为空时，表示还未进行手续费结算处理)"`
	Createtime      time.Time `json:"createtime" description:"创建时间(生成时间)"`
	Updatetime      time.Time `json:"updatetime" description:"更新时间"`
}

func init() {
	orm.RegisterModel(new(InfoGasrecords))
}

/*
根据ID获取gasrecord
*/
func QueryGasrecordsWithID(reordid string) InfoGasrecords {
	orm.Debug = true
	var gasrecords InfoGasrecords
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" * ").From("info_gasrecords").Where("id = " + reordid).Limit(1)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&gasrecords)
	baseCoin := GetBaseCoin()
	gasrecords.Gasused = common.RealValue(gasrecords.Gasused, baseCoin.DecimalUnits)
	return gasrecords
}

/*
获取最后条记录feeid
*/
func QueryLastGasrecordsID() string {
	orm.Debug = true
	var maxfeeid string
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" feeid ").From("info_gasrecords").OrderBy("feeid ").Desc().Limit(1)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&maxfeeid)
	return maxfeeid
}

func QueryGasRecords(contractAddress string) string {
	orm.Debug = true
	var sumgas string
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" sum(gasused) as sumgasused ").From("info_gasrecords").Where("contract_address = '" + contractAddress + "'")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&sumgas)
	baseCoin := GetBaseCoin()
	sumgas = common.RealValue(sumgas, baseCoin.DecimalUnits)
	return sumgas
}

func QueryGasRecordsList(id string, pageNo int, pageSize int, orderBy string) ([]InfoGasrecords, int) {
	orm.Debug = true
	gasRecords := make([]InfoGasrecords, 0)
	o := orm.NewOrm()
	pageStart := (pageNo - 1) * pageSize
	constact := QueryContractWithId(id)
	o.Raw("select * from info_gasrecords where contract_address = '"+constact.ContractAddress+"' order by id "+orderBy+" limit ?, ? ", pageStart, pageSize).QueryRows(&gasRecords)
	var count int
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" count(id) ").From("info_gasrecords where contract_address = '" + constact.ContractAddress + "'").Limit(1)
	sql := qb.String()
	o.Raw(sql).QueryRow(&count)
	baseCoin := GetBaseCoin()
	for index := 0; index < len(gasRecords); index++ {
		var gasRecordsDec = &gasRecords[index]
		gasRecordsDec.Gasused = common.RealValue(gasRecordsDec.Gasused, baseCoin.DecimalUnits)
	}
	return gasRecords, count
}

//查询合约下的所有手续费交易
func QueryContractTxList(contractAddress string, pageNo int, pageSize int, orderBy string) ([]InfoTransactionDetail, int) {
	orm.Debug = true
	gasRecords := make([]InfoTransactionDetail, 0)
	o := orm.NewOrm()
	pageStart := (pageNo - 1) * pageSize

	o.Raw("select * from info_transaction_detail where contract_address = '"+contractAddress+"' order by id "+orderBy+" limit ?, ? ", pageStart, pageSize).QueryRows(&gasRecords)
	var count int
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" count(id) ").From("info_transaction_detail where contract_address = '" + contractAddress + "'").Limit(1)
	sql := qb.String()
	o.Raw(sql).QueryRow(&count)
	baseCoin := GetBaseCoin()
	for index := 0; index < len(gasRecords); index++ {
		var gasRecordsDec = &gasRecords[index]
		gasRecordsDec.Fee = common.RealValue(gasRecordsDec.Fee, baseCoin.DecimalUnits)
	}
	return gasRecords, count
}

func InsertOrUpdateGasRecordsInfo(blockGasRecords []response.GasrecordsModel) error {
	orm.Debug = true
	o := orm.NewOrm()
	err := o.Begin()
	for index := 0; index < len(blockGasRecords); index++ {
		gasRecords := blockGasRecords[index]
		InfoGasrecords := new(InfoGasrecords)
		InfoGasrecords.Feeid = gasRecords.Id
		InfoGasrecords.ContractAddress = gasRecords.ContractAddr
		InfoGasrecords.Ccname = gasRecords.CcName
		InfoGasrecords.Ccversion = gasRecords.CcVersion
		InfoGasrecords.Fcnname = gasRecords.FcnName
		InfoGasrecords.Args = gasRecords.Args
		InfoGasrecords.Address = gasRecords.Address
		InfoGasrecords.Gasused = strconv.FormatInt(gasRecords.GasUsed, 10)

		InfoGasrecords.Txid = gasRecords.TxId
		InfoGasrecords.ReturnTxid = gasRecords.ReturnTxId
		InfoGasrecords.Updatetime = time.Now()
		gasRecordstime, _ := time.Parse("2006-01-02 15:04:05", gasRecords.CreateTime)
		InfoGasrecords.Createtime = gasRecordstime
		_, err = o.InsertOrUpdate(InfoGasrecords)
		if err != nil {
			err = o.Rollback()
			return err
		}
	}
	err = o.Commit()
	return err
}
