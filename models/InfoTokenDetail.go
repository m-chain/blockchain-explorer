package models

import (
	"blockchain_explorer/common"
	"blockchain_explorer/models/response"
	"blockchain_explorer/redis"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
	"github.com/shopspring/decimal"
)

type InfoTokenDetail struct {
	Id           int64  `orm:"column(id);pk;auto"json:"id" description:"主键"` // 设置主键
	Txhash       string `json:"txhash" description:"交易hash"`
	Name         string `json:"name" description:"币种名称"`
	TokenId      string `json:"tokenid" description:"币种ID"`
	TokenSymbol  string `json:"tokensymbol" description:"币种单位"`
	IsBaseCoin   int8   `json:"isbasecoin" description:"是否主币"`
	DecimalUnits string `json:"decimalunits" description:"精确小数位"`
	TotalNumber  string `json:"totalnumber" description:"发行总量"`
	IssuePrice   string `json:"issueprice" description:"发行价格"`
	IssueTime    string `json:"issuetime" description:"发行时间"`
	Status       int8   `json:"status" description:"状态(0失效1有效)"`
	TokenMaster  string `json:"tokenmaster" description:"资产存放地址"`
	RestNumber   string `json:"restNumber" description:"剩余量"`
	Circulation  string `json:"circulation" description:"流通量"`
}

type InfoTokenInfo struct {
	TokenNum int16
}

func init() {
	orm.RegisterModel(new(InfoTokenDetail))
}

func QueryTokenDetail(txhash string) []InfoTokenDetail {
	orm.Debug = true
	tokenDetail := make([]InfoTokenDetail, 0)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" * ").From("info_token_detail").Where("txhash = '" + txhash + "'")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&tokenDetail)
	return tokenDetail
}

func QueryTokenDetailID(tokenid string) InfoTokenDetail {
	var tokenDetail InfoTokenDetail
	tokenDetail = QueryTokenDetailWithID(tokenid)
	if tokenDetail.Id > 0 {
		tokenDetail.RestNumber = common.RealValue(tokenDetail.RestNumber, tokenDetail.DecimalUnits)
		tokenDetail.TotalNumber = common.RealValue(tokenDetail.TotalNumber, tokenDetail.DecimalUnits)
		tokenDetail.Circulation = common.RealValue(tokenDetail.Circulation, tokenDetail.DecimalUnits)
	}
	return tokenDetail
}

func QueryTokenDetailWithID(tokenid string) InfoTokenDetail {
	orm.Debug = true
	var tokenDetail InfoTokenDetail
	if tokenDetail.TokenId == "" {
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select(" * ").From("info_token_detail").Where("token_id = '" + tokenid + "'")
		sql := qb.String()
		o := orm.NewOrm()
		o.Raw(sql).QueryRow(&tokenDetail)
	}
	return tokenDetail
}

func QueryTokenSearch(search string) InfoTokenDetail {
	orm.Debug = true
	var tokenDetail InfoTokenDetail
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" * ").From("info_token_detail").Where("token_id = '" + search + "' or name = '" + search + "' or token_symbol = '" + search + "'").Limit(1)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&tokenDetail)
	fmt.Println(tokenDetail)
	tokenDetail.RestNumber = common.RealValue(tokenDetail.RestNumber, tokenDetail.DecimalUnits)
	tokenDetail.TotalNumber = common.RealValue(tokenDetail.TotalNumber, tokenDetail.DecimalUnits)
	tokenDetail.Circulation = common.RealValue(tokenDetail.Circulation, tokenDetail.DecimalUnits)
	return tokenDetail
}

func QueryTokenList(pageNo int, pageSize int, orderBy string) ([]InfoTokenDetail, int) {
	orm.Debug = true
	tokenDetail := make([]InfoTokenDetail, 0)
	o := orm.NewOrm()
	pageStart := (pageNo - 1) * pageSize
	var listSize int
	o.Raw("select * from info_token_detail where status = '1' order by id "+orderBy+" limit ?, ? ", pageStart, pageSize).QueryRows(&tokenDetail)
	o.Raw("select count(id) from info_token_detail where status = '1' ").QueryRow(&listSize)
	for index := 0; index < len(tokenDetail); index++ {
		var tokenDetailDec = &tokenDetail[index]
		tokenDetailDec.RestNumber = common.RealValue(tokenDetailDec.RestNumber, tokenDetailDec.DecimalUnits)
		tokenDetailDec.TotalNumber = common.RealValue(tokenDetailDec.TotalNumber, tokenDetailDec.DecimalUnits)
		tokenDetailDec.Circulation = common.RealValue(tokenDetailDec.Circulation, tokenDetailDec.DecimalUnits)
	}
	return tokenDetail, listSize
}

func QueryTokenInfo() InfoTokenInfo {
	orm.Debug = true
	var infoTokenInfo InfoTokenInfo
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" count(id) as  token_num").From("info_token_detail").Where("1 = 1").Limit(1)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&infoTokenInfo)
	return infoTokenInfo
}

func InsertOrUpdateTokenDetail(txhash string, blockTokens []response.BlockTokens, tokenMaster map[string]string) error {
	orm.Debug = true
	o := orm.NewOrm()
	err := o.Begin()
	for index := 0; index < len(blockTokens); index++ {
		token := blockTokens[index]
		infoToken := new(InfoTokenDetail)
		if token.IsBaseCoin == true {
			infoToken.IsBaseCoin = 1
		} else {
			infoToken.IsBaseCoin = 0
		}
		infoToken.Txhash = txhash
		infoToken.Name = token.Name
		infoToken.IssuePrice = token.IssuePrice
		infoToken.IssueTime = token.IssueTime
		infoToken.Status = token.Status
		infoToken.TokenId = token.TokenID
		infoToken.TokenSymbol = token.TokenSymbol
		infoToken.TotalNumber = common.HexToString(token.TotalNumber)
		infoToken.DecimalUnits = strconv.FormatInt(token.DecimalUnits, 10)
		if infoToken.RestNumber == "" {
			infoToken.RestNumber = "0"
			infoToken.Circulation = common.HexToString(token.TotalNumber)
		} else {
			infoToken.RestNumber = infoToken.RestNumber
		}
		if tokenMaster[token.TokenID] != "" {
			infoToken.TokenMaster = tokenMaster[token.TokenID]
			_, err = o.InsertOrUpdate(infoToken)
		} else {
			_, err = o.Update(infoToken, "total_number", "status", "token_master")
		}
		if err != nil {
			err = o.Rollback()
			return err
		}
	}
	err = o.Commit()
	return nil
}

func UpdateTokenDetail(token response.BlockTokens) error {
	orm.Debug = true
	o := orm.NewOrm()
	infoToken := new(InfoTokenDetail)
	if token.IsBaseCoin == true {
		infoToken.IsBaseCoin = 1
	} else {
		infoToken.IsBaseCoin = 0
	}

	infoToken.Name = token.Name
	infoToken.IssuePrice = token.IssuePrice
	infoToken.IssueTime = token.IssueTime
	infoToken.Status = token.Status
	infoToken.TokenId = token.TokenID
	infoToken.TokenSymbol = token.TokenSymbol
	infoToken.TotalNumber = common.HexToString(token.TotalNumber)
	infoToken.DecimalUnits = strconv.FormatInt(token.DecimalUnits, 10)
	infoToken.RestNumber = common.HexToString(token.RestNumber)

	//计算大数字字符串
	var totalNumberDec, restNumberDec, resultDec decimal.Decimal
	totalNumberDec, _ = decimal.NewFromString(infoToken.TotalNumber)
	restNumberDec, _ = decimal.NewFromString(infoToken.RestNumber)
	resultDec = totalNumberDec.Sub(restNumberDec)
	fmt.Print(string(resultDec.String()))
	circulation := string(resultDec.String())
	infoToken.Circulation = circulation
	_, err := o.InsertOrUpdate(infoToken)
	infoTokenResult := QueryTokenDetailWithID(token.TokenID)
	redis.SaveCacheTokenInfo(token.TokenID, infoTokenResult)
	return err
}

func GetBaseCoin() InfoTokenDetail {
	orm.Debug = true
	var tokenDetail InfoTokenDetail

	tokenstr := redis.GetCacheBaseCoinInfo()
	var err error
	if tokenstr != "" {
		err = json.Unmarshal([]byte(tokenstr), &tokenDetail)
		fmt.Println(tokenDetail)
		fmt.Println(err)
	} else {
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select(" * ").From("info_token_detail").Where("is_base_coin = 1").Limit(1)
		sql := qb.String()
		o := orm.NewOrm()
		o.Raw(sql).QueryRow(&tokenDetail)
		if tokenDetail.Id > 0 {
			redis.SaveCacheBaseCoin(tokenDetail)
		}
	}

	fmt.Println(tokenDetail)
	return tokenDetail
}

func IsTokenHolder(tokenMaster string) bool {
	orm.Debug = true
	var isHolder int
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" count(*) ").From("info_token_detail").Where("token_master = '" + tokenMaster + "'").Limit(1)
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&isHolder)
	if isHolder > 0 {
		return true
	}
	return false
}

func UpdateTokenExt(wallet response.WalletRestsDetail) (int64, error) {
	orm.Debug = true
	o := orm.NewOrm()
	infoTokenResult := new(InfoTokenDetail)
	infoToken := QueryTokenDetailWithID(wallet.TokenID)
	infoTokenResult.TokenId = infoToken.TokenId
	infoTokenResult.Id = infoToken.Id
	if infoToken.RestNumber == "" {
		infoTokenResult.RestNumber = "0"
	} else {
		infoTokenResult.RestNumber = wallet.RestNumber
	}

	//计算大数字字符串
	var totalNumberDec, restNumberDec, resultDec decimal.Decimal
	totalNumberDec, _ = decimal.NewFromString(infoToken.TotalNumber)
	restNumberDec, _ = decimal.NewFromString(wallet.RestNumber)
	resultDec = totalNumberDec.Sub(restNumberDec)
	fmt.Print(string(resultDec.String()))
	circulation := string(resultDec.String())
	infoTokenResult.Circulation = circulation
	result, err := o.Update(infoTokenResult, "circulation", "rest_number")
	redis.SaveCacheTokenInfo(wallet.TokenID, infoTokenResult)
	return result, err
}

func GetTokenRealValue(tokenvalue string, tokenid string) string {
	infoToken := QueryTokenDetailWithID(tokenid)
	return common.RealValue(tokenvalue, infoToken.DecimalUnits)
}
