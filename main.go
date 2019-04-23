package main

import (
	"encoding/json"
	"fmt"
	prd "github.com/gpestana/bat-pp/predictor"
	sc "github.com/gpestana/bat-pp/scrapper"
	"io/ioutil"
)

const (
	// coincompare confs
	// #TODO: attributes outside
	ccApikey   = "d19289bf83ad76ecd1848c2576381182cd2c3365e3d00040a049b9200d622bf6"
	ccEndpoint = "https://min-api.cryptocompare.com/data/histoday?fsym=BAT&tsym=USD&limit=200"

	//batgrowth confs
	bgEndpoint = "https://batgrowth.com/"

	historicalDataPriceDb  = "./histDataPriceDb.json"
	historicalDataGrowthDb = "./histDataGrowthDb.json"
)

func main() {
	// Gets historical price data from CrytpoCompare
	dp := getAndSaveDataPrice()

	// Gets historical growth data from batgrowth.com
	dg := getAndSaveDataGrowth()

	fmt.Println(dg, dp)

	// Runs predictors (ideally should return value)
	prd.LinearRegression(dg, dp)
}

func getAndSaveDataPrice() []sc.DataPoint {
	cm := sc.NewCryptoCompareScrapper(ccEndpoint, ccApikey)
	data := cm.GetData()
	// re-write historical data into local DB
	dataPrice, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(historicalDataPriceDb, dataPrice, 0644)
	return data
}

func getAndSaveDataGrowth() []sc.DataPointGrowth {
	bgs := sc.NewBatGrowthScrapper(bgEndpoint)
	dataGrowth := bgs.GetData()
	// re-write historical data into local DB
	dataGrowthJson, _ := json.MarshalIndent(dataGrowth, "", " ")
	_ = ioutil.WriteFile(historicalDataGrowthDb, dataGrowthJson, 0644)
	return dataGrowth
}
