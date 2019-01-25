package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type InfoOtherDetail struct {
	Id          int64     `orm:"column(id);pk;auto"json:"id" description:"主键"` // 设置主键
	Txhash      string    `json:"txhash" description:"交易hash"`
	OtherDetail string    `json:"otherdetail" description:"其他数据详情"`
	CreateTime  time.Time `json:"createtime" description:"创建时间"`
}

func init() {
	orm.RegisterModel(new(InfoOtherDetail))
}

func QueryOtherDetail(txhash string) []InfoOtherDetail {
	orm.Debug = true
	transaction := make([]InfoOtherDetail, 0)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" * ").From("info_other_detail").Where("txhash = '" + txhash + "'")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&transaction)
	fmt.Println(transaction)
	return transaction
}

func InsertOrUpdateBlockOther(txhash string, otherdetail interface{}) (int64, error) {
	orm.Debug = true
	o := orm.NewOrm()
	infoOhter := new(InfoOtherDetail)
	infoOhter.Txhash = txhash
	jsonStr, _ := json.Marshal(otherdetail)
	infoOhter.OtherDetail = string(jsonStr)
	infoOhter.CreateTime = time.Now()
	result, err := o.InsertOrUpdate(infoOhter)
	return result, err
}
