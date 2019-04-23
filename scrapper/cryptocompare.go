package scrapper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"net/url"
)

type CoinMarketScrapper struct {
	endpoint string
	apikey   string
}

type CryptoCompareScrapper struct {
	endpoint string
	apikey   string
}

func NewCryptoCompareScrapper(ep string, k string) *CryptoCompareScrapper {
	return &CryptoCompareScrapper{
		endpoint: ep,
		apikey:   k,
	}
}

func (s *CryptoCompareScrapper) GetData() []DataPoint {
	var dataRes []DataPoint

	client := &http.Client{}
	req, err := http.NewRequest("GET", s.endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)

	var data map[string][]map[string]interface{} //interface{}
	json.Unmarshal(respBody, &data)

	dt := data["Data"]
	for _, d := range dt {
		dp := DataPoint{Price: d["close"].(float64), Time: int(d["time"].(float64))}
		dataRes = append(dataRes, dp)
	}

	return dataRes
}
