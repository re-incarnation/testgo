package stalcraftgo

type PriceEntry struct {
	Amount     int32       `json:"amount"`
	Price      int64       `json:"price"`
	Time       string      `json:"time"`
	Additional interface{} `json:"additional"`
}
