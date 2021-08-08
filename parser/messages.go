package parser

import (
	"strings"
	"time"

	"github.com/oliread/marketstack"
)

type date time.Time

func (d *date) UnmarshalJSON(data []byte) error {
	t, err := time.Parse("2006-01-02T15:04:05-0700", strings.Trim(string(data), `"`))
	if err != nil {
		return err
	}

	*d = date(t)
	return nil
}

type endOfDay struct {
	Pagination marketstack.Pagination `json:"pagination"`
	Data       []endOfDayData         `json:"data"`
}

type endOfDayData struct {
	Open           float32 `json:"open"`
	High           float32 `json:"high"`
	Low            float32 `json:"low"`
	Close          float32 `json:"close"`
	Volume         float32 `json:"volume"`
	AdjustedOpen   float32 `json:"adj_open"`
	AdjustedHigh   float32 `json:"adj_high"`
	AdjustedLow    float32 `json:"adj_low"`
	AdjustedClose  float32 `json:"adj_close"`
	AdjustedVolume float32 `json:"adj_volume"`
	SplitFactor    float32 `json:"split_factor"`
	Symbol         string  `json:"symbol"`
	Exchange       string  `json:"exchange"`
	Date           date    `json:"date"`
}
