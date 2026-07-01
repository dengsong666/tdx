package handler

type ResponseDoc struct {
	Code int    `json:"code" example:"0"`
	Msg  string `json:"msg" example:"success"`
	Data any    `json:"data"`
}

type SimplePriceLevelDoc struct {
	Price  float64 `json:"price" example:"10.15"`
	Number int     `json:"number" example:"100"`
}

type SimpleQuoteDoc struct {
	Code       string                `json:"code" example:"000001"`
	Exchange   string                `json:"exchange" example:"sz"`
	Open       float64               `json:"open" example:"10.05"`
	High       float64               `json:"high" example:"10.18"`
	Low        float64               `json:"low" example:"9.99"`
	Close      float64               `json:"close" example:"10.16"`
	Last       float64               `json:"last" example:"10.05"`
	Volume     int64                 `json:"volume" example:"906889"`
	Amount     float64               `json:"amount" example:"915838528"`
	BuyLevels  []SimplePriceLevelDoc `json:"buyLevels"`
	SellLevels []SimplePriceLevelDoc `json:"sellLevels"`
}

type SimpleKlineDoc struct {
	Time   string  `json:"time" example:"2026-07-01T15:00:00+08:00"`
	Last   float64 `json:"last" example:"10.05"`
	Open   float64 `json:"open" example:"10.05"`
	High   float64 `json:"high" example:"10.18"`
	Low    float64 `json:"low" example:"9.99"`
	Close  float64 `json:"close" example:"10.16"`
	Volume int64   `json:"volume" example:"906889"`
	Amount float64 `json:"amount" example:"915838528"`
}

type SimpleMinuteDoc struct {
	Time   string  `json:"time" example:"15:00"`
	Price  float64 `json:"price" example:"10.16"`
	Number int     `json:"number" example:"5609"`
}

type SimpleTradeDoc struct {
	Time   string  `json:"time" example:"2026-07-01T15:00:00+08:00"`
	Price  float64 `json:"price" example:"10.16"`
	Volume int     `json:"volume" example:"5609"`
	Status int     `json:"status" example:"2"`
	Number int     `json:"number" example:"235"`
	Amount float64 `json:"amount" example:"5698744"`
}
