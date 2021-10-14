package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/url"
// 	"sync"
// 	"time"

// 	"github.com/zeeshanakram3/trading/coingeko"
// 	"github.com/zeeshanakram3/trading/coinmarketcap"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func ProcessMapData(context context.Context, collection *mongo.Collection) error {
// 	defer TimeTrack(time.Now())
// 	cmcClient := coinmarketcap.New(cmcBaseUrl, c)
// 	coins, err := cmcClient.MapData(nil)
// 	if err != nil {
// 		return err
// 	}

// 	var wg sync.WaitGroup
// 	wg.Add(len(coins))
// 	aggregator := aggregateUpdateResults()

// 	for _, coin := range coins {
// 		go func(c coingeko.CoinsListBasic) {
// 			defer wg.Done()
// 			filter := bson.M{
// 				"symbol": bson.M{
// 					"$eq": c.Symbol,
// 				},
// 			}
// 			update := bson.M{
// 				"$set": c,
// 			}
// 			res, err := collection.UpdateOne(context, filter, update,
// 				&options.UpdateOptions{Upsert: pointerBool(true)})
// 			if err != nil {
// 				// log.Println(c.ID, err)
// 			}
// 			aggregator(res)
// 		}(coin)
// 	}
// 	wg.Wait()
// 	printAggregatedResults(aggregator())
// 	return err
// }

// func ProcessMarketData(context context.Context, collection *mongo.Collection) error {
// 	defer TimeTrack(time.Now())
// 	cgClient := coingeko.New(coingekoBaseUrl)
// 	q := url.Values{
// 		"vs_currency": []string{"usd"},
// 		"page":        []string{"1"},
// 	}
// 	coins, err := cgClient.CoinsList(q)
// 	if err != nil {
// 		return err
// 	}

// 	var wg sync.WaitGroup
// 	wg.Add(len(coins))
// 	for _, coin := range coins {
// 		go func(c coingeko.CoinsList) {
// 			defer wg.Done()
// 			filter := bson.M{
// 				"symbol": bson.M{
// 					"$eq": c.Symbol,
// 				},
// 			}
// 			update := bson.M{
// 				"$set": c,
// 			}
// 			// TODO should we ignore `res, err`
// 			collection.UpdateOne(context, filter, update,
// 				&options.UpdateOptions{Upsert: pointerBool(true)})
// 		}(coin)
// 	}
// 	wg.Wait()
// 	return nil
// }

// func ProcessMetadata(context context.Context, collection *mongo.Collection) error {
// 	defer TimeTrack(time.Now())
// 	cgClient := coingeko.New(coingekoBaseUrl)
// 	basicCoins, err := cgClient.CoinsListBasic(nil)
// 	if err != nil {
// 		return err
// 	}

// 	q := url.Values{
// 		"localization": []string{"false"},
// 	}

// 	var wg sync.WaitGroup
// 	aggregator := aggregateUpdateResults()
// 	for wgIndex := 0; wgIndex < len(basicCoins)/coingekoConcurrentConnLimit; wgIndex++ {
// 		//ratelimit coins
// 		var rtBasicCoins []coingeko.CoinsListBasic
// 		if len(basicCoins) < (wgIndex+1)*coingekoConcurrentConnLimit {
// 			rtBasicCoins = basicCoins[wgIndex*coingekoConcurrentConnLimit : len(basicCoins)-1]
// 		} else {
// 			rtBasicCoins = basicCoins[wgIndex*coingekoConcurrentConnLimit : (wgIndex+1)*coingekoConcurrentConnLimit]
// 		}

// 		wg.Add(len(rtBasicCoins))
// 		delayBy := 0
// 		for _, basicCoin := range rtBasicCoins {
// 			go func(id string) {
// 				defer wg.Done()
// 				filter := bson.M{
// 					"id": bson.M{
// 						"$eq": id,
// 					},
// 				}
// 				coin, err := findCoinInfo(context, q, id, filter, collection)
// 				if err != nil {
// 					delayBy = delayBy + 1

// 				} else if coin != nil {
// 					update := bson.M{
// 						"$set": coin,
// 					}
// 					res, err := collection.UpdateOne(context, filter, update,
// 						&options.UpdateOptions{Upsert: pointerBool(true)})
// 					if err != nil {
// 						log.Println(err)
// 					}
// 					aggregator(res)
// 				}
// 			}(basicCoin.ID)
// 		}
// 		wg.Wait()
// 		if delayBy > 0 {
// 			fmt.Println(wgIndex, "delayBy", delayBy*2)
// 			time.Sleep(time.Duration(delayBy*2) * time.Second)
// 		}
// 	}

// 	printAggregatedResults(aggregator())
// 	return err
// }

// // // checks if CoinInfo exists in db and calls api if it was last updated more than 12 hours ago.
// // func findCoinInfo(context context.Context, q url.Values, id string, filter interface{}, collection *mongo.Collection) (*coingeko.CoinInfo, error) {
// // 	cgClient := coingeko.New(coingekoBaseUrl)

// // 	var coin *coingeko.CoinInfo
// // 	singleRes := collection.FindOne(context, filter)
// // 	singleRes.Decode(&coin)

// // 	var err error
// // 	if coin == nil {
// // 		coin, err = cgClient.CoinInfo(q, id)
// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 		coin.DocumentLastUpdated = time.Now()
// // 		return coin, nil
// // 	}

// // 	if time.Now().Sub(coin.DocumentLastUpdated) >= 12*time.Hour {
// // 		fmt.Printf("`%s` last updated: %s hours ago \n", id, time.Now().Sub(coin.DocumentLastUpdated)/time.Hour)
// // 		coin, err = cgClient.CoinInfo(q, id)
// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 		coin.DocumentLastUpdated = time.Now()
// // 		return coin, nil
// // 	}
// // 	return nil, nil
// // }

// const cmcConcurrentConnLimit = 10

// func cmcData(context context.Context, db *mongo.Database) error {
// 	err := ProcessMapData(context, db.Collection(cmcMapCollection))
// 	if err != nil {
// 		return err
// 	}
// 	err = ProcessMarketData(context, db.Collection(cmcMarketCollection))
// 	if err != nil {
// 		return err
// 	}
// 	err = ProcessMetadata(context, db.Collection(cmcMetadataCollection))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
