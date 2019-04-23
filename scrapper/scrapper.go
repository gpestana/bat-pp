package scrapper

type Scrapper interface {
	GetData() []DataPoint
}

type DataPoint struct {
	Time  int
	Price float64
}

type DataPointGrowth struct {
	Time    int
	Results int
}
