package orders

import "time"

type OrderBook struct {
	Symbol       string    `json:"symbol"`
	BidsPrices   []float64 `json:"bidsprices"`
	AsksPrices   []float64 `json:"asksprices"`
	BidsQuantity []int     `json:"bidsquantility"`
	AsksQantity  []int     `json:"asksquantility"`
	TimeStamp    time.Time `json:"timestamp"`
}
