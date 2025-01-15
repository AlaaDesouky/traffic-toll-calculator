package main

import (
	"fmt"
	"log"
	"os"
	"traffic-toll-calculator/types"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	var (
		store = makeStore()
		_     = NewAggregatorService(store)
	)

	fmt.Println("aggregator Service")
}

func makeStore() Storer {
	storeType := os.Getenv("AGGREGATOR_STORE_TYPE")
	switch types.StoreType(storeType) {
	case types.MemeoryStoreType:
		return NewMemoryStore()
	default:
		log.Fatalf("invalid store type given %s\n", storeType)
		return nil
	}
}
