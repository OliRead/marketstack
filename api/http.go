package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/oliread/marketstack"
)

type HTTP struct {
	baseURL string
	apiKey  string

	httpClient *http.Client
}

func NewClient(baseURL string, apiKey string, httpClient *http.Client) *HTTP {
	return &HTTP{
		baseURL:    baseURL,
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

func (h *HTTP) EndOfDay(ctx context.Context, symbols []string, order marketstack.SortOrder, exchange string, dateFrom, dateTo time.Time, limit, offset int) (*http.Response, error) {
	endpoint := fmt.Sprintf("%s/eod", h.baseURL)

	query := make(map[string][]string)
	query["symbols"] = []string{strings.Join(symbols, ",")}
	query["sort"] = []string{string(order)}
	query["offset"] = []string{strconv.Itoa(offset)}

	if exchange != "" {
		query["exchange"] = []string{exchange}
	}
	if !dateFrom.IsZero() {
		query["date_from"] = []string{dateFrom.Format("2006-01-02T15:04:05-0700")}
	}
	if !dateTo.IsZero() {
		query["date_to"] = []string{dateTo.Format("2006-01-02T15:04:05-0700")}
	}
	if limit != 0 {
		query["limit"] = []string{strconv.Itoa(limit)}
	}

	return h.execute(ctx, endpoint, query)
}

func (h *HTTP) EndOfDayDate(ctx context.Context, date time.Time, symbols []string, order marketstack.SortOrder, exchange string, dateFrom, dateTo time.Time, limit, offset int) (*http.Response, error) {
	endpoint := fmt.Sprintf("%s/eod/%s", h.baseURL, date.Format("2006-01-02"))

	query := make(map[string][]string)
	query["symbols"] = []string{strings.Join(symbols, ",")}
	query["offset"] = []string{strconv.Itoa(offset)}

	if exchange != "" {
		query["exchange"] = []string{exchange}
	}
	if !dateFrom.IsZero() {
		query["date_from"] = []string{dateFrom.Format("2006-01-02T15:04:05-0700")}
	}
	if !dateTo.IsZero() {
		query["date_to"] = []string{dateTo.Format("2006-01-02T15:04:05-0700")}
	}
	if limit != 0 {
		query["limit"] = []string{strconv.Itoa(limit)}
	}

	return h.execute(ctx, endpoint, query)
}

func (h *HTTP) EndOfDayLatest(ctx context.Context, symbols []string, order marketstack.SortOrder, exchange string, dateFrom, dateTo time.Time, limit, offset int) (*http.Response, error) {
	endpoint := fmt.Sprintf("%s/eod/latest", h.baseURL)

	query := make(map[string][]string)
	query["symbols"] = []string{strings.Join(symbols, ",")}
	query["offset"] = []string{strconv.Itoa(offset)}

	if exchange != "" {
		query["exchange"] = []string{exchange}
	}
	if !dateFrom.IsZero() {
		query["date_from"] = []string{dateFrom.Format("2006-01-02T15:04:05-0700")}
	}
	if !dateTo.IsZero() {
		query["date_to"] = []string{dateTo.Format("2006-01-02T15:04:05-0700")}
	}
	if limit != 0 {
		query["limit"] = []string{strconv.Itoa(limit)}
	}

	return h.execute(ctx, endpoint, query)
}

func (h *HTTP) execute(ctx context.Context, endpoint string, query url.Values) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	query["access_key"] = []string{h.apiKey}
	req.URL.RawQuery = query.Encode()

	return h.httpClient.Do(req)
}
