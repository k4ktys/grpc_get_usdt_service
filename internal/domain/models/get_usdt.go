package models

type UsdtTrade struct {
	Timestamp int64 `json:"timestamp"`

	Asks []UsdtCandle `json:"asks"`
	Bids []UsdtCandle `json:"bids"`
}

type UsdtCandle struct {
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Amount string `json:"amount"`
	Factor string `json:"factor"`
	Type   string `json:"type"`
}
