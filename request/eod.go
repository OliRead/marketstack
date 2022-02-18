package request

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type EOD struct {
	symbols  []string
	exchange *string
	sort     *SortOrder
	dateFrom *time.Time
	dateTo   *time.Time
	limit    *int
	offset   *int
}

func NewEOD(symbols []string, opts ...EODOption) (*EOD, error) {
	if len(symbols) == 0 {
		return nil, fmt.Errorf("%w:no symbols", ErrSymbols)
	}

	msg := EOD{
		symbols: symbols,
	}

	for _, opt := range opts {
		opt(&msg)
	}

	return &msg, nil
}

func (e *EOD) Query() url.Values {
	v := url.Values{}
	v.Add("symbols", strings.Join(e.symbols, ","))

	if e.exchange != nil {
		v.Add("exchange", *e.exchange)
	}

	if e.sort != nil {
		v.Add("sort", e.sort.String())
	}

	if e.dateFrom != nil {
		v.Add("date_from", e.dateFrom.Format(dateEncodingFormat))
	}

	if e.dateTo != nil {
		v.Add("date_to", e.dateFrom.Format(dateEncodingFormat))
	}

	if e.limit != nil {
		v.Add("limit", strconv.Itoa(*e.limit))
	}

	if e.offset != nil {
		v.Add("offset", strconv.Itoa(*e.limit))
	}

	return v
}

func (e *EOD) Action() action {
	return ActionEOD
}

type EODOption func(e *EOD)

func EODWithExchange(exchange string) EODOption {
	return func(e *EOD) {
		if exchange == "" {
			return
		}

		e.exchange = &exchange
	}
}

func EODWithSortOrder(order SortOrder) EODOption {
	return func(e *EOD) {
		e.sort = &order
	}
}

func EODWithDateRange(from, to *time.Time) EODOption {
	return func(e *EOD) {
		if e.dateFrom != nil {
			e.dateFrom = from
		}

		if e.dateTo != nil {
			e.dateTo = to
		}
	}
}

func EODWithLimit(limit int) EODOption {
	return func(e *EOD) {
		e.limit = &limit
	}
}

func EODWithOffset(offset int) EODOption {
	return func(e *EOD) {
		e.offset = &offset
	}
}
