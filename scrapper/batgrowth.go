package scrapper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type BatGrowthScrapper struct {
	endpoint string
}

func NewBatGrowthScrapper(ep string) *BatGrowthScrapper {
	return &BatGrowthScrapper{
		endpoint: ep,
	}
}

func (s *BatGrowthScrapper) GetData() []DataPointGrowth {
	var dataRes []DataPointGrowth

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

	strBody := string(respBody)
	//parse data

	spl := strings.FieldsFunc(strBody, SplitFn)
	dt := spl[1]

	// ugly and without time
	//sepA := rune('{')
	//sepB := rune('"')
	//for _, dt := range a {
	//	if rune(a[0]) == sepA {
	//		if rune(a[1]) == sepB {
	//			fmt.Println(dt)
	//		}
	//	}
	//}

	// takes only first occurrence of a data set
	dt = "[" + dt + "]"
	var dat []map[string]interface{}
	if err := json.Unmarshal([]byte(dt), &dat); err != nil {
		log.Fatal(err)
	}

	for _, dp := range dat {
		d := DataPointGrowth{Time: int(dp["Created_at"].(float64)), Results: int(dp["Number_of_results"].(float64))}
		dataRes = append(dataRes, d)
	}

	return dataRes

}

func SplitFn(r rune) bool {
	return r == '[' || r == ']'
}
