package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type ResponseModel struct {
	Result    interface{} `json:"result"`
	Message   string      `json:"message"`
	Errorcode string      `json:"errorcode"`
	Status    bool        `json:"status"`
}

func GetResponse(blocks interface{}, isSuccess bool, message string, errorcode string) ResponseModel {

	var resp ResponseModel
	resp.Status = isSuccess
	resp.Message = message
	resp.Result = &blocks
	resp.Errorcode = errorcode
	return resp
}
