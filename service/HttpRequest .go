package service

import (
	"blockchain_explorer/constant"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
)

func HttpGet(urlString string) (string, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(urlString)
	if err != nil || resp.StatusCode != http.StatusOK {
		logs.SetLogger("console")
		logs.Error(err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.SetLogger("console")
		logs.Error(err)
		return "", err
	}
	return string(body), nil

}

func HttpPost(urlString string, param url.Values) (string, error) {
	paramString := param.Encode()

	req, err := http.NewRequest("POST", urlString, strings.NewReader(paramString))
	if err != nil {
		logs.SetLogger("console")
		logs.Error(err)
		return "", err
	}

	// 表单方式
	req.Header.Set("Content-Type", constant.ContentType)
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		logs.SetLogger("console")
		logs.Error(err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.SetLogger("console")
		logs.Error(err)
		return "", err
	}
	return string(body), nil

}
