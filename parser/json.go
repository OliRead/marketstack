package parser

import (
	"encoding/json"
	"io"
	"time"

	"github.com/oliread/marketstack"
)

type JSON struct {
}

func (j *JSON) ParseEndOfDay(r io.ReadCloser) (marketstack.EndOfDay, error) {
	msg := endOfDay{}

	if err := json.NewDecoder(r).Decode(&msg); err != nil {
		return marketstack.EndOfDay{}, err
	}

	data := make([]marketstack.EndOfDayResponseData, len(msg.Data))
	for i, res := range msg.Data {
		data[i] = marketstack.EndOfDayResponseData{
			Open:           res.Open,
			High:           res.High,
			Low:            res.Low,
			Close:          res.Close,
			Volume:         res.Volume,
			AdjustedOpen:   res.AdjustedOpen,
			AdjustedHigh:   res.AdjustedHigh,
			AdjustedLow:    res.AdjustedLow,
			AdjustedClose:  res.AdjustedClose,
			AdjustedVolume: res.AdjustedVolume,
			SplitFactor:    res.SplitFactor,
			Symbol:         res.Symbol,
			Exchange:       res.Exchange,
			Date:           time.Time(res.Date),
		}
	}

	return marketstack.EndOfDay{
		Pagination: msg.Pagination,
		Data:       data,
	}, nil
}
