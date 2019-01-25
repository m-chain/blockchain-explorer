package models

import (
	"blockchain_explorer/common"
	"blockchain_explorer/models/response"
	"math/big"
	"strconv"
	"time"

	"github.com/shopspring/decimal"

	"github.com/astaxie/beego/orm"
)

type InfoGasreturn struct {
	Id                int64     `orm:"column(id);pk;auto"json:"id"  description:"主键"` // 设置主键
	Blockid           string    `json:"blockid" description:"区块ID"`
	TokenId           string    `json:"tokenID" description:"代币ID"`
	StartTime         time.Time `json:"startTime" description:"开始时间(生效时间)"`
	InitReleaseamount string    `json:"initReleaseAmount" description:"初次释放数量(立即释放)"`
	Amount            string    `json:"amount" description:"交易数量(总量)"`
	Interval          int64     `json:"interval" description:"时间间隔(秒)"`
	Periods           int64     `json:"periods" description:"总期数"`
	Withdrawed        string    `json:"withdrawed" description:"已经返还数量"`
	Txid              string    `json:"txID" description:"交易ID"`
	Reason            string    `json:"reason" description:"原因"`
	Createtime        string    `json:"createTime" description:"创建时间"`
	Address           string    `json:"address" description:"地址"`
	Gasfeeid          string    `json:"gasfeeid" description:"返还手续费ID"`
	ContractAddress   string    `json:"contractaddress" description:"合约地址"`
	GasReturnhash     string    `json:"gasreturnhash" description:"返还手续费hash"`
}

func init() {
	orm.RegisterModel(new(InfoGasreturn))
}

func QueryGasReturnsList(id string, pageNo int, pageSize int, orderBy string) ([]InfoGasreturn, int) {
	orm.Debug = true
	gasReturns := make([]InfoGasreturn, 0)
	o := orm.NewOrm()
	pageStart := (pageNo - 1) * pageSize
	constact := QueryContractWithId(id)
	o.Raw("select * from info_gasreturn where contract_address = '"+constact.ContractAddress+"' order by id "+orderBy+" limit ?, ? ", pageStart, pageSize).QueryRows(&gasReturns)
	var count int
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" count(id) ").From("info_gasreturn where contract_address = '" + constact.ContractAddress + "'").Limit(1)
	sql := qb.String()
	o.Raw(sql).QueryRow(&count)
	baseCoin := GetBaseCoin()
	for index := 0; index < len(gasReturns); index++ {
		var gasReturnsDec = &gasReturns[index]
		initReleaseAmount, _ := new(big.Int).SetString(gasReturnsDec.InitReleaseamount, 10)
		amount, _ := new(big.Int).SetString(gasReturnsDec.Amount, 10)
		withdrawed := vestingFunc(time.Now().Unix(), gasReturnsDec.StartTime.Unix(), initReleaseAmount, amount, gasReturnsDec.Interval, gasReturnsDec.Periods)
		gasReturnsDec.Withdrawed = common.RealValue(withdrawed.String(), baseCoin.DecimalUnits)
		gasReturnsDec.Amount = common.RealValue(gasReturnsDec.Amount, baseCoin.DecimalUnits)
		gasReturnsDec.InitReleaseamount = common.RealValue(gasReturnsDec.InitReleaseamount, baseCoin.DecimalUnits)
	}
	return gasReturns, count
}

func QueryGasReturnFee(contractAddress string) string {
	orm.Debug = true
	var sumgas string
	orm.Debug = true
	gasReturns := make([]InfoGasreturn, 0)
	o := orm.NewOrm()
	o.Raw("select * from info_gasreturn where contract_address = '" + contractAddress + "' ").QueryRows(&gasReturns)
	var resultDec decimal.Decimal
	resultDec, _ = decimal.NewFromString("0")
	baseCoin := GetBaseCoin()
	for index := 0; index < len(gasReturns); index++ {
		var gasReturnsDec = &gasReturns[index]
		initReleaseAmount, _ := new(big.Int).SetString(gasReturnsDec.InitReleaseamount, 10)
		amount, _ := new(big.Int).SetString(gasReturnsDec.Amount, 10)
		withdrawed := vestingFunc(time.Now().Unix(), gasReturnsDec.StartTime.Unix(), initReleaseAmount, amount, gasReturnsDec.Interval, gasReturnsDec.Periods)

		var withdrawedDec decimal.Decimal
		withdrawedDec, _ = decimal.NewFromString(withdrawed.String())
		resultDec = resultDec.Add(withdrawedDec)
	}
	sumgas = common.RealValue(string(resultDec.String()), baseCoin.DecimalUnits)
	return sumgas
}

func InsertOrUpdateGasReturns(blockid string, blockGasReturns []response.BlockGasReturns) error {
	orm.Debug = true
	o := orm.NewOrm()
	err := o.Begin()
	for index := 0; index < len(blockGasReturns); index++ {
		gasReturns := blockGasReturns[index]
		InfoGasreturns := new(InfoGasreturn)
		InfoGasreturns.Blockid = blockid
		InfoGasreturns.InitReleaseamount = common.HexToString(gasReturns.InitReleaseAmount)
		InfoGasreturns.Amount = common.HexToString(gasReturns.Amount)
		InfoGasreturns.Interval = gasReturns.Interval
		InfoGasreturns.Periods = gasReturns.Periods
		InfoGasreturns.Reason = gasReturns.Reason
		formatTimeStr := time.Unix(gasReturns.StartTime, 0).Format("2006-01-02 15:04:05")
		formatTime, err := time.Parse("2006-01-02 15:04:05", formatTimeStr)
		InfoGasreturns.StartTime = formatTime
		InfoGasreturns.TokenId = gasReturns.TokenID
		InfoGasreturns.Txid = gasReturns.TxID
		InfoGasreturns.Withdrawed = common.HexToString(gasReturns.Withdrawed)
		InfoGasreturns.Createtime = gasReturns.CreateTime
		InfoGasreturns.Address = gasReturns.Address
		InfoGasreturns.Gasfeeid = strconv.FormatInt(gasReturns.Id, 10)
		InfoGasreturns.ContractAddress = gasReturns.ContractAddress
		InfoGasreturns.GasReturnhash = gasReturns.GasReturnhash
		_, err = o.InsertOrUpdate(InfoGasreturns)
		if err != nil {
			err = o.Rollback()
			return err
		}
	}
	err = o.Commit()
	return err
}

func vestingFunc(currentTime, startTime int64, initReleaseAmount, amount *big.Int, interval, periods int64) *big.Int {
	if currentTime < startTime {
		return big.NewInt(0)
	}
	minus := currentTime - startTime
	end := periods * interval
	if minus >= end {
		return amount
	}

	ip := big.NewInt(0)
	iAmount := ip.Sub(amount, initReleaseAmount).Div(ip, big.NewInt(periods))
	i := minus / interval
	ip = big.NewInt(0)
	return ip.Mul(iAmount, big.NewInt(i)).Add(ip, initReleaseAmount)
}
