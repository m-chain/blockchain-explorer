package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetGasAddress",
            Router: `/GetGasAddress`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetBaseCoin",
            Router: `/getBaseCoin`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetBlockInfomation",
            Router: `/getBlockInfomation`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetBlocks",
            Router: `/getBlocks`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetBlocksList",
            Router: `/getBlocksList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetContractDetailList",
            Router: `/getContractDetailList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetContractDetailWithID",
            Router: `/getContractDetailWithID`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetContractGasFeeList",
            Router: `/getContractGasFeeList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetContractGasRecordList",
            Router: `/getContractGasRecordList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetContractGasRecordWithID",
            Router: `/getContractGasRecordWithID`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetContractTxList",
            Router: `/getContractTxList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetNodeInfomation",
            Router: `/getNodeInfomation`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetQueryBalance",
            Router: `/getQueryBalance`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetTokenDetail",
            Router: `/getTokenDetail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetTokenDetailList",
            Router: `/getTokenDetailList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetTokenDetailWithID",
            Router: `/getTokenDetailWithID`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetTransactionDetail",
            Router: `/getTransactionDetail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetTransactionDetailList",
            Router: `/getTransactionDetailList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetTransactionDetailWithHash",
            Router: `/getTransactionDetailWithHash`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetTransactions",
            Router: `/getTransactions`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "GetWalletDetail",
            Router: `/getWalletDetail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"] = append(beego.GlobalControllerRouter["blockchain_explorer/controllers:APIController"],
        beego.ControllerComments{
            Method: "SearchInformation",
            Router: `/searchInformation`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
