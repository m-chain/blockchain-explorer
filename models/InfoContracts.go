package models

import (
	"blockchain_explorer/common"
	"blockchain_explorer/models/response"
	"fmt"

	"github.com/astaxie/beego/orm"
)

type InfoContracts struct {
	Id              int64  `orm:"column(id);pk;auto"json:"id" description:"主键"` // 设置主键
	BlockId         string `json:"blockid" description:"区块ID"`
	Name            string `json:"name" description:"合约名称"`
	ContractAddress string `json:"contractAddress" description:"合约地址"`
	ContractSymbol  string `json:"contractSymbol" description:"合约名称简写(英文简称)"`
	Maddress        string `json:"mAddress" description:"拥有平台币地址"` //拥有平台币地址
	Version         string `json:"version" description:"合约版本号"`
	ContractPath    string `orm:"column(contract_path)" json:"contractPath" description:"合约路径"`
	Remark          string `json:"remark" description:"备注"`
	Status          string `json:"status" description:"合约状态(-1已删除 ,1待初始化 ,2正在运行 ,3余额不足)"`
	GasFee          string `json:"gasfee" description:"手续费"`
	GasReturnFee    string `json:"gasReturnFee" description:"已返还手续费"`
	Createtime      string `json:"createTime" description:"创建时间"`
	Updatetime      string `json:"updatetime" description:"更新时间"`
}

func init() {
	orm.RegisterModel(new(InfoContracts))
}

func QueryContractCount() int {
	orm.Debug = true
	var count int
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" count(id) ").From("info_contracts").Limit(1)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&count)
	return count
}

func QueryContractWithId(contractid string) InfoContracts {
	orm.Debug = true
	var contract InfoContracts
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" * ").From("info_contracts").Where("id = '" + contractid + "'")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&contract)
	contract.GasFee = QueryGasRecords(contract.ContractAddress)
	contract.GasReturnFee = QueryGasReturnFee(contract.ContractAddress)
	return contract
}

func QueryContractWithIdForPld(contractid string) InfoContracts {
	orm.Debug = true
	var contract InfoContracts
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" * ").From("info_contracts").Where("id = '" + contractid + "'")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&contract)

	var list orm.ParamsList
	num, err := o.Raw("SELECT SUM(fee) as gas_fee FROM info_transaction_detail WHERE contract_address=?", contract.ContractAddress).ValuesFlat(&list)
	if err == nil && num > 0 {
		baseCoin := GetBaseCoin()
		sumgas := common.RealValue(fmt.Sprintf("%v", list[0]), baseCoin.DecimalUnits)
		contract.GasFee = sumgas
	}
	contract.GasReturnFee = "0"
	return contract
}

func QueryContractSearch(search string) InfoContracts {
	orm.Debug = true
	var contract InfoContracts
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" * ").From("info_contracts").Where("name = '" + search + "' or contract_address = '" + search + "'").Limit(1)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&contract)
	contract.GasFee = QueryGasRecords(contract.ContractAddress)
	contract.GasReturnFee = QueryGasReturnFee(contract.ContractAddress)
	return contract
}

func QueryContractList(pageNo int, pageSize int, orderBy string) []InfoContracts {
	orm.Debug = true
	contracts := make([]InfoContracts, 0)
	o := orm.NewOrm()
	pageStart := (pageNo - 1) * pageSize
	o.Raw("select * from info_contracts  order by id "+orderBy+" limit ?, ? ", pageStart, pageSize).QueryRows(&contracts)
	fmt.Println(contracts)
	return contracts
}

func InsertOrUpdateContact(blockid string, blockContract []response.BlockContracts) error {
	orm.Debug = true
	o := orm.NewOrm()
	err := o.Begin()
	for index := 0; index < len(blockContract); index++ {
		contract := blockContract[index]
		InfoContracts := new(InfoContracts)
		InfoContracts.BlockId = blockid
		InfoContracts.Name = contract.Name
		InfoContracts.ContractAddress = contract.ContractAddress
		InfoContracts.ContractPath = contract.ContractPath
		InfoContracts.Status = contract.Status
		InfoContracts.Version = contract.Version
		InfoContracts.Maddress = contract.Maddress
		InfoContracts.Remark = contract.Remark
		InfoContracts.ContractSymbol = contract.ContractSymbol
		InfoContracts.Createtime = contract.CreateTime
		InfoContracts.Updatetime = contract.UpdateTime

		_, err = o.InsertOrUpdate(InfoContracts)
		if err != nil {
			err = o.Rollback()
			return err
		}
	}
	err = o.Commit()
	return err
}
