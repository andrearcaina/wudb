package main

import (
	"github.com/andrearcaina/wudb/api"
	"github.com/andrearcaina/wudb/kvstore"
)

func main() {
	kvStore := kvstore.NewKVStore()

	server := api.NewServer(kvStore)

	err := server.Start(":8080")
	if err != nil {
		return
	}
}
