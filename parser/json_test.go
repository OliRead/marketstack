package parser_test

import (
	"io"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/oliread/marketstack"
	"github.com/oliread/marketstack/parser"
)

func TestJsonEndOfDaySuccess(t *testing.T) {
	payload := `
	{
		"pagination": {
			"limit": 100,
			"offset": 0,
			"count": 100,
			"total": 9944
		},
		"data": [
			{
				"open": 129.8,
				"high": 133.04,
				"low": 129.47,
				"close": 132.995,
				"volume": 106686703.0,
				"adj_high": 133.04,
				"adj_low": 129.47,
				"adj_close": 132.995,
				"adj_open": 129.8,
				"adj_volume": 106686703.0,
				"split_factor": 1.0,
				"symbol": "AAPL",
				"exchange": "XNAS",
				"date": "2021-04-09T00:00:00+0000"
			}
		]
	}
	`

	expect := marketstack.EndOfDay{
		Pagination: marketstack.Pagination{
			Limit:  100,
			Offset: 0,
			Count:  100,
			Total:  9944,
		},
		Data: []marketstack.EndOfDayResponseData{
			{
				Open:           129.8,
				High:           133.04,
				Low:            129.47,
				Close:          132.995,
				Volume:         106686703.0,
				AdjustedOpen:   129.8,
				AdjustedHigh:   133.04,
				AdjustedLow:    129.47,
				AdjustedClose:  132.995,
				AdjustedVolume: 106686703.0,
				SplitFactor:    1.0,
				Symbol:         "AAPL",
				Exchange:       "XNAS",
				Date:           time.Date(2021, time.April, 9, 0, 0, 0, 0, time.FixedZone("", 0)),
			},
		},
	}

	parser := parser.JSON{}
	msg, err := parser.ParseEndOfDay(io.NopCloser(strings.NewReader(payload)))
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expect, msg) {
		t.Fatalf("EXPECT: %+v\nACTUAL: %+v", expect, msg)
	}
}
