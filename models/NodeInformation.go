package models

import (
	"blockchain_explorer/models/response"
	"blockchain_explorer/redis"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type NodeInformation struct {
	Id          int64     `orm:"column(id);pk;auto"json:"id"  description:"主键"` // 设置主键
	NodeName    string    `json:"nodename"  description:"节点名称"`
	NodeIp      string    `json:"nodeip"  description:"节点IP"`
	NodeAddress string    `json:"nodeaddress"  description:"节点位置"`
	ContainerId string    `json:"containerid"  description:"节点ID"`
	RunStatus   string    `json:"runstatus"  description:"运行状态"`
	BlockHeight int64     `json:"blockheight"  description:"区块高度"`
	Longitude   string    `json:"lon"  description:"经度"`
	Latitude    string    `json:"lat"  description:"纬度"`
	UpdateTime  time.Time `json:"updatetime"  description:"更新时间"`
}

func init() {
	orm.RegisterModel(new(NodeInformation))
}

func QueryNodeInfomation() []NodeInformation {

	//判断是否有redis
	node := make([]NodeInformation, 0)
	nodestr := redis.GetCacheNodeInfo()
	var err error
	if nodestr != "" {
		err = json.Unmarshal([]byte(nodestr), &node)
		fmt.Println(node)
		fmt.Println(err)
	}

	if node == nil || err != nil {
		orm.Debug = true
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select(" * ").From("node_information")

		sql := qb.String()

		o := orm.NewOrm()
		o.Raw(sql).QueryRows(&node)
		fmt.Println(node)
	}
	return node
}

func InsertOrUpdateNodeInfomation(nodeInfo response.NodeInfo) error {
	orm.Debug = true
	o := orm.NewOrm()
	err := o.Begin()
	if len(nodeInfo.Data) <= 0 {
		return errors.New("data is empty")
	} else {
		var nodes []*NodeInformation = make([]*NodeInformation, len(nodeInfo.Data))
		for index := 0; index < len(nodeInfo.Data); index++ {
			nodeInfo := nodeInfo.Data[index]
			nodeInformation := new(NodeInformation)
			nodeInformation.NodeAddress = nodeInfo.PeerLocation
			nodeInformation.NodeIp = nodeInfo.PeerIp
			nodeInformation.NodeName = nodeInfo.PeerName
			nodeInformation.RunStatus = nodeInfo.PeerStatus
			nodeInformation.UpdateTime = time.Now()
			nodeInformation.ContainerId = nodeInfo.ContainerId
			nodeInformation.Latitude = nodeInfo.PeerLatitude
			nodeInformation.Longitude = nodeInfo.PeerLongitude
			nodeInformation.BlockHeight = QueryBlockInfomation().BlockHeight
			_, err = o.InsertOrUpdate(nodeInformation)
			if err != nil {
				err = o.Rollback()
				return err
			}
			nodes[index] = nodeInformation
		}
		err = o.Commit()
		redis.SaveCacheNodeInfo(nodes)
	}
	return err
}
