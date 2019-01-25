package redis

import (
	"blockchain_explorer/constant"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/gomodule/redigo/redis"
)

func SaveCacheBlockInfo(blocks interface{}) {
	bm, err := cache.NewCache("redis", constant.RedisSource)
	if err != nil {
		fmt.Println(err)
	} else {
		blocksString, _ := json.Marshal(blocks)
		err = bm.Put("BlockInfomationCache", string(blocksString), 3*time.Second)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func GetCacheBlockInfo() string {
	bm, err := cache.NewCache("redis", constant.RedisSource)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if bm.IsExist("BlockInfomationCache") == false {
		fmt.Println("GetCacheBlockInfo数据不存在")
		return ""
	} else {
		blocks := bm.Get("BlockInfomationCache")
		blockstr, _ := redis.String(blocks, err)
		return blockstr
	}
}

func SaveCacheNodeInfo(nodes interface{}) {
	bm, err := cache.NewCache("redis", constant.RedisSource)
	if err != nil {
		fmt.Println(err)
	} else {
		nodesString, _ := json.Marshal(nodes)
		err = bm.Put("NodeInfoCache", string(nodesString), 3600*time.Second)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func GetCacheNodeInfo() string {
	bm, err := cache.NewCache("redis", constant.RedisSource)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if bm.IsExist("NodeInfoCache") == false {
		fmt.Println("GetCacheNodeInfo数据不存在")
		return ""
	} else {
		nodes := bm.Get("NodeInfoCache")
		nodestr, _ := redis.String(nodes, err)
		return nodestr
	}
}

func SaveCacheTokenInfo(tokenid string, token interface{}) {
	bm, err := cache.NewCache("redis", constant.RedisSource)
	if err != nil {
		fmt.Println(err)
	} else {
		tokenString, _ := json.Marshal(token)
		name := "TokenInfoCache_" + tokenid
		err = bm.Put(name, string(tokenString), 5*time.Second)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func GetCacheTokenInfo(tokenid string) string {
	bm, err := cache.NewCache("redis", constant.RedisSource)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	name := "TokenInfoCache_" + tokenid
	if bm.IsExist(name) == false {
		fmt.Println("GetCacheTokenInfo数据不存在")
		return ""
	} else {
		token := bm.Get(name)
		tokenstr, _ := redis.String(token, err)
		return tokenstr
	}
}

func SaveCacheBaseCoin(token interface{}) error {
	bm, err := cache.NewCache("redis", constant.RedisSource)
	if err != nil {
		return err
	} else {
		tokenString, _ := json.Marshal(token)
		name := "BaseCoinCache"
		err = bm.Put(name, string(tokenString), 36000*time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetCacheBaseCoinInfo() string {
	bm, err := cache.NewCache("redis", constant.RedisSource)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	name := "BaseCoinCache"
	if bm.IsExist(name) == false {
		fmt.Println("GetCacheBaseCoinInfo数据不存在")
		return ""
	} else {
		token := bm.Get(name)
		tokenstr, _ := redis.String(token, err)
		return tokenstr
	}
}
