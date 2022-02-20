package api

import (
	"strings"
	"time"
)

type Date time.Time
type DateTime time.Time

func (d *Date) UnmarshalJSON(data []byte) error {
	t, err := time.Parse("2006-01-02", strings.Trim(string(data), `"`))
	if err != nil {
		return err
	}

	*d = Date(t.UTC())
	return nil
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	t, err := time.Parse("2006-01-02T15:04:05-0700", strings.Trim(string(data), `"`))
	if err != nil {
		return err
	}

	*d = DateTime(t.UTC())
	return nil
}

type Error struct {
	Code    string                    `json:"code"`
	Message string                    `json:"message"`
	Context map[string][]ErrorContext `json:"context"`
}

type ErrorContext struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

type Pagination struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Count  int    `json:"count"`
	Total  uint64 `json:"total"`
}

type EOD struct {
	Pagination Pagination `json:"pagination"`
	Data       []EODValue `json:"data"`
	Error      Error      `json:"error"`
}

type EODValue struct {
	Open           float32  `json:"open"`
	High           float32  `json:"high"`
	Low            float32  `json:"low"`
	Close          float32  `json:"close"`
	Volume         float32  `json:"volume"`
	AdjustedOpen   float32  `json:"adj_open"`
	AdjustedHigh   float32  `json:"adj_high"`
	AdjustedLow    float32  `json:"adj_low"`
	AdjustedClose  float32  `json:"adj_close"`
	AdjustedVolume float32  `json:"adj_volume"`
	SplitFactor    float32  `json:"split_factor"`
	Dividend       float32  `json:"dividend"`
	Symbol         string   `json:"symbol"`
	Exchange       string   `json:"exchange"`
	Date           DateTime `json:"date"`
}

type Dividends struct {
	Pagination Pagination      `json:"pagination"`
	Data       []DividendsData `json:"data"`
	Error      Error           `json:"error"`
}

type DividendsData struct {
	Date     Date    `json:"date"`
	Dividend float32 `json:"dividend"`
	Symbol   string  `json:"symbol"`
}

type Splits struct {
	Pagination Pagination   `json:"pagination"`
	Data       []SplitsData `json:"data"`
	Error      Error        `json:"error"`
}

type SplitsData struct {
	Date        Date    `json:"date"`
	SplitFactor float32 `json:"split_factor"`
	Symbol      string  `json:"symbol"`
}
