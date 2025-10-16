package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/andrearcaina/wudb/api"
	"github.com/andrearcaina/wudb/kvstore"
)

func main() {
	addr, persistence, finalAOFPath := ParseFlags()

	kvStore := kvstore.NewKVStore(persistence, finalAOFPath)

	server := api.NewServer(kvStore)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Fatal(server.Start(":" + addr))
	}()

	<-sigChan
	log.Println("Shutting down server...")

	if err := kvStore.Close(); err != nil {
		log.Printf("Error closing kvstore: %v", err)
	}
}
