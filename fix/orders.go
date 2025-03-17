package fix

import "time"

type OrderBook struct {
	Symbol         string    `json:"symbol"`
	BidsPrices     []float64 `json:"bidsprices"`
	AsksPrices     []float64 `json:"asksprices"`
	BidsQuantility []int     `json:"bidsquantility"`
	AsksQantility  []int     `json:"asksquantility"`
	TimeStamp      time.Time `json:"timestamp"`
}
