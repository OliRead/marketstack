package marketstack

import "time"

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Count  int `json:"count"`
	Total  int `json:"total"`
}

type EndOfDay struct {
	Pagination Pagination             `json:"pagination"`
	Data       []EndOfDayResponseData `json:"data"`
}

type EndOfDayResponseData struct {
	Open           float32   `json:"open"`
	High           float32   `json:"high"`
	Low            float32   `json:"low"`
	Close          float32   `json:"close"`
	Volume         float32   `json:"volume"`
	AdjustedHigh   float32   `json:"adj_high"`
	AdjustedLow    float32   `json:"adj_low"`
	AdjustedClose  float32   `json:"adj_close"`
	AdjustedOpen   float32   `json:"adj_open"`
	AdjustedVolume float32   `json:"adj_volume"`
	SplitFactor    float32   `json:"split_factor"`
	Symbol         string    `json:"symbol"`
	Exchange       string    `json:"exchange"`
	Date           time.Time `json:"date"`
}
