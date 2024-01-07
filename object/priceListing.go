package stalcraftgo

type PriceListing struct {
	Total  int64    `json:"total"`
	Prices []Prices `json:"prices"` //?{}/[]
}

type Prices struct {
	Amount     int32       `json:"amount"`
	Price      int64       `json:"price"`
	Time       string      `json:"time"`
	Additional interface{} `json:"additional"`
}

type ItemPriceHistoryResponse struct {
	Total  int64    `json:"total"`
	Prices []Prices `json:"prices"`
}

type ActiveItemLotsResponse struct {
	Total  int64    `json:"total"`
	Prices []Prices `json:"prices"`
}
