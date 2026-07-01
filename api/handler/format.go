package handler

import (
	"strings"
	"time"

	"github.com/injoyai/tdx/protocol"
)

const (
	formatRaw    = "raw"
	formatSimple = "simple"
)

func responseFormat(s string) (string, bool) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "", formatSimple:
		return formatSimple, true
	case formatRaw:
		return formatRaw, true
	default:
		return "", false
	}
}

func price(p protocol.Price) float64 {
	return p.Float64()
}

type simplePriceLevel struct {
	Price  float64 `json:"price"`
	Number int     `json:"number"`
}

type simpleQuote struct {
	Code       string             `json:"code"`
	Exchange   string             `json:"exchange"`
	Open       float64            `json:"open"`
	High       float64            `json:"high"`
	Low        float64            `json:"low"`
	Close      float64            `json:"close"`
	Last       float64            `json:"last"`
	Volume     int64              `json:"volume"`
	Amount     float64            `json:"amount"`
	BuyLevels  []simplePriceLevel `json:"buyLevels"`
	SellLevels []simplePriceLevel `json:"sellLevels"`
}

func simpleQuotes(quotes protocol.QuotesResp) []simpleQuote {
	result := make([]simpleQuote, 0, len(quotes))
	for _, q := range quotes {
		item := simpleQuote{
			Code:     q.Code,
			Exchange: q.Exchange.String(),
		}
		if q.Kline != nil {
			item.Open = price(q.Kline.Open)
			item.High = price(q.Kline.High)
			item.Low = price(q.Kline.Low)
			item.Close = price(q.Kline.Close)
			item.Last = price(q.Kline.Last)
			item.Volume = q.Kline.Volume
			item.Amount = price(q.Kline.Amount)
		}
		for _, level := range q.BuyLevel {
			item.BuyLevels = append(item.BuyLevels, simplePriceLevel{
				Price:  price(level.Price),
				Number: level.Number,
			})
		}
		for _, level := range q.SellLevel {
			item.SellLevels = append(item.SellLevels, simplePriceLevel{
				Price:  price(level.Price),
				Number: level.Number,
			})
		}
		result = append(result, item)
	}
	return result
}

type simpleKline struct {
	Time   string  `json:"time"`
	Last   float64 `json:"last"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
	Amount float64 `json:"amount"`
}

func simpleKlines(klines protocol.Klines) []simpleKline {
	result := make([]simpleKline, 0, len(klines))
	for _, k := range klines {
		result = append(result, simpleKline{
			Time:   k.Time.Format(time.RFC3339),
			Last:   price(k.Last),
			Open:   price(k.Open),
			High:   price(k.High),
			Low:    price(k.Low),
			Close:  price(k.Close),
			Volume: k.Volume,
			Amount: price(k.Amount),
		})
	}
	return result
}

type simpleMinute struct {
	Time   string  `json:"time"`
	Price  float64 `json:"price"`
	Number int     `json:"number"`
}

func simpleMinutes(items []protocol.PriceNumber) []simpleMinute {
	result := make([]simpleMinute, 0, len(items))
	for _, item := range items {
		result = append(result, simpleMinute{
			Time:   item.Time,
			Price:  price(item.Price),
			Number: item.Number,
		})
	}
	return result
}

type simpleTrade struct {
	Time   string  `json:"time"`
	Price  float64 `json:"price"`
	Volume int     `json:"volume"`
	Status int     `json:"status"`
	Number int     `json:"number"`
	Amount float64 `json:"amount"`
}

func simpleTrades(items protocol.Trades) []simpleTrade {
	result := make([]simpleTrade, 0, len(items))
	for _, item := range items {
		result = append(result, simpleTrade{
			Time:   item.Time.Local().Format(time.RFC3339),
			Price:  price(item.Price),
			Volume: item.Volume,
			Status: item.Status,
			Number: item.Number,
			Amount: price(item.Amount()),
		})
	}
	return result
}
