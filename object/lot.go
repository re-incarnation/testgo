package stalcraftgo

type Lot struct {
	ItemId       string      `json:"item_id"`
	Amount       int32       `json:"amount"`
	StartPrice   int64       `json:"start_price"`
	CurrentPrice int64       `json:"current_price"`
	BuyOutPrice  int64       `json:"buy_out_price"`
	StartTime    string      `json:"start_time"`
	EndTime      string      `json:"end_time"`
	Additional   interface{} `json:"additional"`
}
