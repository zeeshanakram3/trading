package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"sync"
	"time"

	"github.com/zeeshanakram3/trading/coingeko"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDb() {
	dbName := os.Getenv("MONGO_ATLAS_DB_NAME")
	dbPasswd := os.Getenv("MONGO_ATLAS_DB_PASSWORD")
	if dbName == "" || dbPasswd == "" {
		panic("dbName or dbPasswd does not exits")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	mongoURI := `mongodb+srv://zeeshan:` + dbPasswd +
		`@cluster0.5uqoj.mongodb.net/` + dbName + `?retryWrites=true&w=majority`
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
}

func pointerBool(b bool) *bool {
	return &b
}

func makeInterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}
	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}
	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)

	// Skip this function, and fetch the PC and file for its parent.
	pc, _, _, _ := runtime.Caller(1)

	// Retrieve a function object this functions parent.
	funcObj := runtime.FuncForPC(pc)

	// Regex to extract just the function name (and not the module path).
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

	fmt.Println(fmt.Sprintf("%s took %s", name, elapsed))
}

func aggregateUpdateResults() func(result ...*mongo.UpdateResult) mongo.UpdateResult {
	aggregated := mongo.UpdateResult{
		MatchedCount:  0,
		ModifiedCount: 0,
		UpsertedCount: 0,
	}
	var mu sync.Mutex
	return func(result ...*mongo.UpdateResult) mongo.UpdateResult {
		if len(result) > 0 && result[0] != nil {
			mu.Lock()
			aggregated.MatchedCount += result[0].MatchedCount
			aggregated.ModifiedCount += result[0].ModifiedCount
			aggregated.UpsertedCount += result[0].UpsertedCount
			mu.Unlock()
		}
		return aggregated
	}
}

func printAggregatedResults(res mongo.UpdateResult) {
	fmt.Printf("%d matched, %d modified, %d upserted \n",
		res.MatchedCount, res.ModifiedCount, res.UpsertedCount)
}

type ByTicker []coingeko.Tickers

func (a ByTicker) Len() int           { return len(a) }
func (a ByTicker) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTicker) Less(i, j int) bool { return a[i].Last < a[j].Last }

func findArbitrageOpportunities() func(coin *coingeko.CoinInfo) {
	file, err := os.Create("arbitrage.csv")
	checkError(err)
	writer := csv.NewWriter(file)
	// defer writer.Flush()

	return func(coin *coingeko.CoinInfo) {
		tickerPrice := coin.Tickers
		sort.Sort(ByTicker(tickerPrice))
		var maxPriceDiffPercentage float64
		for i := 0; i < len(tickerPrice)-1; i++ {
			if tickerPrice[i].Target == "USDT" && tickerPrice[i+1].Target == "USDT" {
				maxPriceDiffPercentage = math.Abs(tickerPrice[i].Last-tickerPrice[i+1].Last) /
					tickerPrice[i].Last
			}
			if maxPriceDiffPercentage != math.Inf(0) && maxPriceDiffPercentage >= 0.05 {
				err := writer.Write([]string{coin.Symbol, fmt.Sprintf("%f", maxPriceDiffPercentage), fmt.Sprintf("%f", tickerPrice[i].Last), tickerPrice[i].Market.Name, fmt.Sprintf("%f", tickerPrice[i+1].Last), tickerPrice[i+1].Market.Name})
				checkError(err)
				// fmt.Println(coin.Symbol, maxPriceDiffPercentage, tickerPrice[i].TradeURL, tickerPrice[i+1].TradeURL)
			}
		}
	}
}
