package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var clientOptions *options.ClientOptions

func main() {
	ctx := context.TODO()
	db := client.Database(os.Getenv("MONGO_ATLAS_DB_NAME"))

	// err := coingeckoData(ctx, db)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	err := coingeckoData(ctx, db)
	checkError(err)
}

const (
	cmcBaseUrl      = "https://pro-api.coinmarketcap.com"
	coingekoBaseUrl = "https://api.coingecko.com/api/v3"
	messariBaseUrl  = "https://data.messari.io/api"

	coingekoListBasicCollection = "coingeko_listbasic"
	coingekoListCollection      = "coingeko_list"
	coingekoCoinsCollection     = "coingeko_coins"

	cmcMapCollection      = "cmc_map"
	cmcMarketCollection   = "cmc_market"
	cmcMetadataCollection = "cmc_metadata"
)

func init() {
	connectDb()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
