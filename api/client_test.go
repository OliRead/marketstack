package api_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/oliread/marketstack/api"
	"github.com/oliread/marketstack/request"
)

func NewMockServer(responseCode int, responseData []byte) *httptest.Server {
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(responseCode)
			w.Write(responseData)
		}),
	)

	return s
}

func TestEOD(t *testing.T) {
	tcs := []struct {
		name       string
		server     *httptest.Server
		eodOptions request.EODOption
		err        error
		eod        api.EOD
	}{
		{
			name: "Success",
			server: NewMockServer(200, []byte(`
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
						"dividend": 0.0,
						"symbol": "AAPL",
						"exchange": "XNAS",
						"date": "2021-04-09T00:00:00+0000"
					}
				]
			}
			`)),
			eod: api.EOD{
				Pagination: api.Pagination{
					Limit:  100,
					Offset: 0,
					Count:  100,
					Total:  9944,
				},
				Data: []api.EODValue{
					{
						Open:           129.8,
						High:           133.04,
						Low:            129.47,
						Close:          132.995,
						Volume:         106686703,
						AdjustedOpen:   129.8,
						AdjustedHigh:   133.04,
						AdjustedLow:    129.47,
						AdjustedClose:  132.995,
						AdjustedVolume: 106686703,
						SplitFactor:    1,
						Dividend:       0,
						Symbol:         "AAPL",
						Exchange:       "XNAS",
						Date: api.DateTime(
							time.Date(2021, time.April, 9, 0, 0, 0, 0, time.UTC),
						),
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Parallel()

		builder, err := request.NewBuilder(
			request.BuilderWithBaseURL(tc.server.URL + "/"),
		)
		if err != nil {
			t.Fatal(err)
		}

		client := api.NewClient(builder)

		eod, err := client.EOD(context.Background(), []string{"AAPL"})
		if err != nil {
			if !errors.Is(err, tc.err) {
				t.Fatalf("%s does not unwrap to %s", err, tc.err)
			}
		}

		if !reflect.DeepEqual(tc.eod, eod) {
			t.Fatalf("EXPECTED: %+v\nACTUAL: %+v", tc.eod, eod)
		}
	}
}

func TestDividends(t *testing.T) {
	tcs := []struct {
		name      string
		server    *httptest.Server
		options   request.DividendsOption
		err       error
		dividends api.Dividends
	}{
		{
			name: "Success",
			server: NewMockServer(200, []byte(`
			{
				"pagination": {
					"limit": 100,
					"offset": 0,
					"count": 100,
					"total": 50765
				},
				"data": [
					{
						"date": "2021-05-24",
						"dividend": 0.5,
						"symbol": "IAU"
					}
				]
			}
			`)),
			dividends: api.Dividends{
				Pagination: api.Pagination{
					Limit:  100,
					Offset: 0,
					Count:  100,
					Total:  50765,
				},
				Data: []api.DividendsData{
					{
						Date: api.Date(
							time.Date(2021, time.May, 24, 0, 0, 0, 0, time.UTC),
						),
						Dividend: 0.5,
						Symbol:   "IAU",
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Parallel()

		builder, err := request.NewBuilder(
			request.BuilderWithBaseURL(tc.server.URL + "/"),
		)
		if err != nil {
			t.Fatal(err)
		}

		client := api.NewClient(builder)

		dividends, err := client.Dividends(context.Background(), []string{"AAPL"})
		if err != nil {
			if !errors.Is(err, tc.err) {
				t.Fatalf("%s does not unwrap to %s", err, tc.err)
			}
		}

		if !reflect.DeepEqual(tc.dividends, dividends) {
			t.Fatalf("EXPECTED: %+v\nACTUAL: %+v", tc.dividends, dividends)
		}
	}
}

func TestSplits(t *testing.T) {
	tcs := []struct {
		name    string
		server  *httptest.Server
		options request.SplitsOption
		err     error
		splits  api.Splits
	}{
		{
			name: "Success",
			server: NewMockServer(200, []byte(`
			{
				"pagination": {
					"limit": 100,
					"offset": 0,
					"count": 100,
					"total": 50765
				},
				"data": [
					{
						"date": "2021-05-24",
						"split_factor": 0.5,
						"symbol": "IAU"
					}
				]
			}
			`)),
			splits: api.Splits{
				Pagination: api.Pagination{
					Limit:  100,
					Offset: 0,
					Count:  100,
					Total:  50765,
				},
				Data: []api.SplitsData{
					{
						Date: api.Date(
							time.Date(2021, time.May, 24, 0, 0, 0, 0, time.UTC),
						),
						SplitFactor: 0.5,
						Symbol:      "IAU",
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Parallel()

		builder, err := request.NewBuilder(
			request.BuilderWithBaseURL(tc.server.URL + "/"),
		)
		if err != nil {
			t.Fatal(err)
		}

		client := api.NewClient(builder)

		splits, err := client.Splits(context.Background(), []string{"AAPL"})
		if err != nil {
			if !errors.Is(err, tc.err) {
				t.Fatalf("%s does not unwrap to %s", err, tc.err)
			}
		}

		if !reflect.DeepEqual(tc.splits, splits) {
			t.Fatalf("EXPECTED: %+v\nACTUAL: %+v", tc.splits, splits)
		}
	}
}
