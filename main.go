package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/gpestana/bat-pp/predictor"
	sc "github.com/gpestana/bat-pp/scrapper"
	"io/ioutil"
)

const (
	apikey                 = "d19289bf83ad76ecd1848c2576381182cd2c3365e3d00040a049b9200d622bf6"
	endpoint               = "https://min-api.cryptocompare.com/data/histoday?fsym=BAT&tsym=USD&limit=10"
	historicalDataPriceDb  = "./histDataPriceDb.json"
	historicalDataGrowthDb = "./histDataGrowthDb.json"
)

func main() {

	// Gets historical price data from CrytpoCompare
	cm := sc.NewCryptoCompareScrapper(endpoint, apikey)
	data := cm.GetData()
	// re-write historical data into local DB
	dataPrice, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(historicalDataPriceDb, dataPrice, 0644)

	//Gets historical growth data from batgrowth.com
}
