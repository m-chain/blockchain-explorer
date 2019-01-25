package schedule

import (
	"blockchain_explorer/block"
	"blockchain_explorer/constant"
	"blockchain_explorer/models"
	"blockchain_explorer/models/response"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego/toolbox"
)

// 启动定时器
func OpenTask() {
	blocksyncTask()
	peerInfoTask()
	gasRecordsTask()
	chainTask()
	transactionTask()
	syncEmptyBlockTask()
	toolbox.StartTask()
}

func blocksyncTask() {
	tk := toolbox.NewTask("blocksyncTask", "0/1 * * * * *", blocksyncGo)
	err := tk.Run()
	if err != nil {
		fmt.Println(err)
	}
	toolbox.AddTask("blocksyncTask", tk)

}

func peerInfoTask() {
	tk := toolbox.NewTask("peerInfoTask", "*/10 * * * *", peerInfo)
	err := tk.Run()
	if err != nil {
		fmt.Println(err)
	}
	toolbox.AddTask("peerInfoTask", tk)

}

func gasRecordsTask() {
	tk := toolbox.NewTask("gasRecordsTask", "0/1 * * * * *", gasRecords)
	err := tk.Run()
	if err != nil {
		fmt.Println(err)
	}
	toolbox.AddTask("gasRecordsTask", tk)

}

func chainTask() {
	tk := toolbox.NewTask("chainTask", "0/1 * * * * *", chainInfo)
	err := tk.Run()
	if err != nil {
		fmt.Println(err)
	}
	toolbox.AddTask("chainTask", tk)

}

func transactionTask() {
	tk := toolbox.NewTask("TransactionTask", "0/1 * * * * *", transactionStatus)
	err := tk.Run()
	if err != nil {
		fmt.Println(err)
	}
	toolbox.AddTask("TransactionTask", tk)
}

func syncEmptyBlockTask() {
	tk := toolbox.NewTask("SyncEmptyBlockTask", "0/10 * * * * *", saveBlockExt)
	err := tk.Run()
	if err != nil {
		fmt.Println(err)
	}
	toolbox.AddTask("SyncEmptyBlockTask", tk)
}

func blocksyncGo() error {
	go blocksync(1)
	// go blocksync(2)
	// go blocksync(3)
	// go blocksync(4)
	// go blocksync(5)
	return nil
}

func blocksync(id int) error {
	fmt.Println("========== blocksync ===========")
	var currentBlockNumber = block.GetBlockNumber()
	var blocknumber int
	var dataBlockNumer = models.QueryBlockNumber()
	if currentBlockNumber.Errorcode == constant.StatusCodeFail {
		return nil
	} else if dataBlockNumer == "" {
		blocknumber = 0
	} else {
		blocknumber, _ = strconv.Atoi(dataBlockNumer)
		blocknumber = blocknumber + id
	}

	// if blocknumber > currentBlockNumber.Result.(int) {
	// 	fmt.Printf("已经达到最新高度")
	// 	return nil
	// }
	var blocksResp, err = block.BlocksDetail(strconv.Itoa(blocknumber))
	var blocksExtResp, exterr = block.BlocksExtDetail(strconv.Itoa(blocknumber))
	if err != nil || exterr != nil {
		fmt.Println(err)
		fmt.Println(exterr)
		fmt.Println("区块获取不正常")
	} else {
		if blocksResp.Data.Blockid != "" {
			var blocktableid int64
			if blocksResp.Status == true {
				blocktableid, err = models.InsertOrUpdateBlockInfo(blocksResp.Data)
			}
			if blocksExtResp.Status == true {
				_, exterr = models.UpdateExtBlockInfo(blocktableid, blocksExtResp.Data)
			}
			if err != nil || exterr != nil {
				fmt.Println(err)
				fmt.Println(exterr)
				return nil
			}
			if len(blocksResp.Data.Transactions) <= 0 {
				return nil
			}
			//保存区块交易数据
			models.InsertOrUpdateBlockTransaction(blocksResp.Data.Blockid, blocksResp.Data.Transactions)
			for index := 0; index < len(blocksResp.Data.Transactions); index++ {
				blockTrans := blocksResp.Data.Transactions[index]
				if len(blockTrans.Trans) > 0 {
					models.InsertOrUpdateBlockTransactionDetail(blocksResp.Data.Blockid, blockTrans.Txhash, blockTrans.Trans)
				}
				if len(blockTrans.Tokens) > 0 {
					models.InsertOrUpdateTokenDetail(blockTrans.Txhash, blockTrans.Tokens, blockTrans.TokenMaster)
				}
				if len(blockTrans.Wallets) > 0 {
					models.InsertOrUpdateBlockWallet(blockTrans.Txhash, blockTrans.Wallets)
				}
				if blockTrans.Others != nil {
					models.InsertOrUpdateBlockOther(blockTrans.Txhash, blockTrans.Others)
				}
				if len(blockTrans.Contracts) > 0 {
					models.InsertOrUpdateContact(blocksResp.Data.Blockid, blockTrans.Contracts)
				}
				if len(blockTrans.GasReturns) > 0 {
					models.InsertOrUpdateGasReturns(blocksResp.Data.Blockid, blockTrans.GasReturns)
				}
			}
		}
		models.UpdateBlockInfomation()

	}
	return nil
}

func peerInfo() error {
	fmt.Println("========== peerInfo ===========")
	var nodePeerInfo, _ = block.GetNodePeerInfo()
	var err error
	if nodePeerInfo.Code == 200 {
		models.InsertOrUpdateNodeInfomation(nodePeerInfo)
	}

	return err
}

func updateToken(blockTrans []response.BlockTrans) error {
	var err error
	for index := 0; index < len(blockTrans); index++ {
		trans := blockTrans[index]
		var walletData response.WalletBalance
		if models.IsTokenHolder(trans.FromAddress) {
			walletData, err = block.QueryBalance(trans.FromAddress)
		} else if models.IsTokenHolder(trans.ToAddress) {
			walletData, err = block.QueryBalance(trans.ToAddress)
		}
		for j := 0; j < len(walletData.Data.WalletRests); j++ {
			rest := walletData.Data.WalletRests[j]
			if rest.TokenID == trans.TokenID {
				_, err = models.UpdateTokenExt(rest)
			}
		}
	}
	return err
}

func gasRecords() error {
	fmt.Println("========== gasRecords ===========")
	feeid := models.QueryLastGasrecordsID()
	if feeid == "" {
		feeid = "0"
	}
	var gasRecordsInfo, _ = block.GetGasRecordsInfo(feeid)
	var err error
	if gasRecordsInfo.Status == true {
		models.InsertOrUpdateGasRecordsInfo(gasRecordsInfo.Data)
	}
	return err
}

func chainInfo() error {
	fmt.Println("========== chainInfo ===========")
	chainInfo, err := block.GetChainInfo()
	if chainInfo.Data.CurrentHash != "" {
		models.InsertOrUpdateChainInfo(chainInfo.Data)
	}
	return err
}

func transactionStatus() error {
	transaction := models.QueryTransactionDetailWait()
	var txhashArr []string
	for index := 0; index < len(transaction); index++ {
		tx := transaction[index]
		txhashArr = append(txhashArr, "\""+tx.Txhash+"\"")
	}
	txhashStr := strings.Join(txhashArr, ",")
	txhashStr = "[" + txhashStr + "]"
	transactionStatus, _ := block.GetQueryMultiTransferStatus(txhashStr)
	if transactionStatus.Status == true {
		for index := 0; index < len(transactionStatus.Data); index++ {
			transaction := transactionStatus.Data[index]
			models.UpdateTransactionStatus(transaction)
		}

	}

	return nil
}

/**
区块数据补录方法
**/
func saveBlockExt() error {
	emptyBlock := models.QueryEmptyDataBlockNumber()
	for index := 0; index < len(emptyBlock); index++ {
		blockEmpty := emptyBlock[index]
		blockNumberStr := blockEmpty.BlockId
		blocksExtResp, _ := block.BlocksExtDetail(strconv.FormatInt(blockNumberStr, 10))
		if blocksExtResp.Status == true {
			models.UpdateExtBlockInfo(blockEmpty.Id, blocksExtResp.Data)
		}
	}

	return nil
}
