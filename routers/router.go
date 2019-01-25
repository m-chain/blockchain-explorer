// @APIVersion 1.0
// @Title 区块浏览器接口文档
// @Description 区块浏览器相关API接口
// @Contact	telong@m-chain.com
package routers

import (
	"blockchain_explorer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//========== 直接调用链上接口 ===== //
	beego.Router("/blocks/getInfo", &controllers.BlocksController{}, "get,post:GetInfo")
	beego.Router("/blocks/getBlockTransactionCountByNumber/:blocknum", &controllers.BlocksController{}, "get,post:GetBlockTransactionCountByNumber")
	beego.Router("/blocks/queryTokenInfo/:tokenid", &controllers.BlocksController{}, "get,post:QueryTokenInfo")
	beego.Router("/blocks/blocksDetail", &controllers.BlocksController{}, "get,post:BlocksDetail")
	beego.Router("/blocks/blockTransactionCountByHash/:blockhash", &controllers.BlocksController{}, "get,post:BlockTransactionCountByHash")
	beego.Router("/blocks/getTransactionByHash/:hash", &controllers.BlocksController{}, "get,post:TransactionByHash")
	beego.Router("/blocks/getChainInfo", &controllers.BlocksController{}, "get,post:ChainInfo")
	beego.Router("/blocks/getTransactionByBlockHashAndIndex", &controllers.BlocksController{}, "get,post:TransactionByBlockHashAndIndex")
	beego.Router("/blocks/getTransactionByBlockNumberAndIndex", &controllers.BlocksController{}, "get,post:TransactionByBlockNumberAndIndex")
	beego.Router("/blocks/queryTxsByAddress", &controllers.BlocksController{}, "post:QueryTxsByAddress")
	beego.Router("/blocks/queryMultiTxInfo", &controllers.BlocksController{}, "post:QueryMultiTxInfo")
	beego.Router("/blocks/getBlockNumber", &controllers.BlocksController{}, "post:GetBlockNumber")
	beego.Router("/blocks/queryTransferInfos", &controllers.BlocksController{}, "post:QueryTransferInfos")
	beego.Router("/blocks/queryTransferInfo", &controllers.BlocksController{}, "post:QueryTransferInfo")
	beego.Router("/blocks/queryTransferCount", &controllers.BlocksController{}, "post:QueryTransferCount")

	ns :=
		beego.NewNamespace("/blocks",
			beego.NSNamespace("/browser",
				beego.NSInclude(
					&controllers.APIController{},
				),
			),
		)
	beego.AddNamespace(ns)

}
