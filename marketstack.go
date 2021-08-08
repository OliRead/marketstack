package marketstack

import (
	"context"
	"io"
	"net/http"
	"time"
)

type SortOrder string

const (
	SortOrderAscending  = SortOrder("ASC")
	SortOrderDescending = SortOrder("DESC")
)

type API interface {
	EndOfDay(ctx context.Context, symbols []string, order SortOrder, exchange string, dateFrom, dateTo time.Time, limit, offset int) (*http.Response, error)
	EndOfDayDate(ctx context.Context, date time.Time, symbols []string, order SortOrder, exchange string, dateFrom, dateTo time.Time, limit, offset int) (*http.Response, error)
	EndOfDayLatest(ctx context.Context, symbols []string, order SortOrder, exchange string, dateFrom, dateTo time.Time, limit, offset int) (*http.Response, error)
}

type Parser interface {
	ParseEndOfDay(r io.ReadCloser) (EndOfDay, error)
}
