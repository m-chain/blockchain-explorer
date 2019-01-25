package constant

import "github.com/astaxie/beego"

const (
	URLGetBlockNumber                   = "v1/wallet/getBlockNumber"                      //获取区块高度
	URLGetBlockTransactionCountByNumber = "v1/wallet/getBlockTransactionCountByNumber/"   //根据区块号获取交易数
	URLQueryTokenInfo                   = "v1/wallet/queryTokenInfo/"                     //获取token信息
	URLQueryBalance                     = "v1/wallet/queryBalance/"                       //查询余额
	URLBlocksDetail                     = "v1/wallet/blocks/"                             //查询区块信息
	URLBlocksExtDetail                  = "v1/wallet/queryBlockExtInfo/"                  //查询区块扩展信息
	URLBlockTransactionCountByHash      = "v1/wallet/getBlockTransactionCountByHash"      //根据区块hash获取交易数
	URLTransactionByHash                = "v1/wallet/getTransactionByHash"                //根据交易hash获取交易详情
	URLChainInfo                        = "v1/wallet/getChainInfo"                        //获取链信息
	URLTransactionByBlockHashAndIndex   = "v1/wallet/getTransactionByBlockHashAndIndex"   //根据区块hash和交易在区块中的位置查询交易信息
	URLTransactionByBlockNumberAndIndex = "v1/wallet/getTransactionByBlockNumberAndIndex" //根据区块号和交易在区块中的位置查询交易信息
	URLQueryTxsByAddress                = "v1/wallet/queryTxsByAddress/"                  //查询钱包地址下信息变动交易
	URLQueryMultiTxInfo                 = "v1/wallet/queryMultiTxInfo"                    //查询多地址转账信息
	URLQueryTransferInfos               = "v1/wallet/queryTransferInfos/"                 //查询钱包地址下所有转账记录
	URLQueryTransferInfo                = "v1/wallet/queryTransferInfo"                   //查询转账详情记录
	URLQueryTransferCount               = "v1/wallet/queryTransferCount/"                 //查询钱包地址下所有转账次数
	URLGetChainInfo                     = "v1/wallet/getChainInfo"                        //获取链信息
	URLNodePeerInfo                     = "v1/peer/getPeerInfo"                           //获取节点数据
	URLQueryGasRecordsInfo              = "v1/wallet/queryGasRecords/"                    //查询手续费记录列表
	URLGetGasAddressInfo                = "v1/wallet/queryGasAddress"                     //查询手续费地址
	URLQueryTransferStatus              = "v1/wallet/queryTransferStatus/"                //查询转账状态
	URLQueryMultiTransferStatus         = "v1/wallet/queryMultiTransferStatus/"           //批量查询转账状态

	URLLocation = "https://api.map.baidu.com/location/ip" //百度根据IP获取经纬度

	StatusCodeSuccess = "1000" //正常状态码
	StatusCodeFail    = "0000" //错误状态码

	ContentType = "application/x-www-form-urlencoded"
)

const (
	Empty = iota
	BlockId
	Address
	TokenId
	Contract
	Transaction
	TokenIdAndAddress
)

var BlockChainURL string = beego.AppConfig.String("blockchainurl")
var NodePeerURL string = beego.AppConfig.String("peernodeurl")
var RedisSource string = beego.AppConfig.String("RedisSource")

// var CoinUnit = "6"
