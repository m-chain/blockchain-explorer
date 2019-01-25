package models

import (
	"blockchain_explorer/models/response"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type InfoWalletDetail struct {
	Id           int64     `orm:"column(id);pk;auto" json:"id" description:"主键"` // 设置主键
	Txhash       string    `orm:"column(txhash)" json:"transationid" description:"交易hash"`
	Address      string    `json:"address" description:"地址"`
	Pubkey       string    `json:"pubkey" description:"公钥"`
	CreateTime   time.Time `orm:"column(create_time)" json:"createtime" description:"创建时间"`
	ReleasePlans string    `orm:"type(text)" json:"releaseplans" description:"释放计划JSON"`
	IsLocked     int64     `orm:"column(is_locked)" json:"islocked" description:"是否锁定"`
}

func init() {
	orm.RegisterModel(new(InfoWalletDetail))
}

func QueryWalletDetail(txhash string) []InfoWalletDetail {
	orm.Debug = true
	walletDetail := make([]InfoWalletDetail, 0)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" * ").From("info_wallet_detail").Where("txhash = '" + txhash + "'")

	sql := qb.String()

	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&walletDetail)
	fmt.Println(walletDetail)
	return walletDetail
}

func QueryAddressSearch(address string) InfoWalletDetail {
	orm.Debug = true
	var walletDetail InfoWalletDetail
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" * ").From("info_wallet_detail").Where("address = '" + address + "'")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRow(&walletDetail)
	fmt.Println(walletDetail)
	return walletDetail
}

func InsertOrUpdateBlockWallet(txhash string, blockWallet []response.BlockWallets) error {
	orm.Debug = true
	o := orm.NewOrm()
	err := o.Begin()
	for index := 0; index < len(blockWallet); index++ {
		wallet := blockWallet[index]
		infoWallet := new(InfoWalletDetail)
		infoWallet.Address = wallet.Address
		infoWallet.CreateTime = time.Now()
		infoWallet.IsLocked = wallet.IsLocked
		infoWallet.Pubkey = wallet.PubKey
		jsonstr, _ := json.Marshal(wallet.ReleasePlans)
		infoWallet.ReleasePlans = string(jsonstr)
		infoWallet.Txhash = txhash
		_, err = o.InsertOrUpdate(infoWallet)
		if err != nil {
			err = o.Rollback()
			return err
		}
	}
	err = o.Commit()
	return err
}
