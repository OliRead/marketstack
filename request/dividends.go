package request

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Dividends struct {
	symbols  []string
	sort     *SortOrder
	dateFrom *time.Time
	dateTo   *time.Time
	limit    *int
	offset   *int
}

type DividendsOption func(d *Dividends)

func NewDividends(symbols []string, opts ...DividendsOption) (*Dividends, error) {
	if len(symbols) == 0 {
		return nil, fmt.Errorf("%w:no symbols", ErrSymbols)
	}

	msg := Dividends{
		symbols: symbols,
	}

	for _, opt := range opts {
		opt(&msg)
	}

	return &msg, nil
}

func (d *Dividends) Query() url.Values {
	v := url.Values{}
	v.Add("symbols", strings.Join(d.symbols, ","))

	if d.sort != nil {
		v.Add("sort", d.sort.String())
	}

	if d.dateFrom != nil {
		v.Add("date_from", d.dateFrom.Format(dateEncodingFormat))
	}

	if d.dateTo != nil {
		v.Add("date_to", d.dateFrom.Format(dateEncodingFormat))
	}

	if d.limit != nil {
		v.Add("limit", strconv.Itoa(*d.limit))
	}

	if d.offset != nil {
		v.Add("offset", strconv.Itoa(*d.limit))
	}

	return v
}

func (d *Dividends) Action() action {
	return ActionDividends
}

func DividendsWithSortOrder(order SortOrder) DividendsOption {
	return func(d *Dividends) {
		d.sort = &order
	}
}

func DividendsWithDateRange(from, to *time.Time) DividendsOption {
	return func(d *Dividends) {
		if d.dateFrom != nil {
			d.dateFrom = from
		}

		if d.dateTo != nil {
			d.dateTo = to
		}
	}
}

func DividendsWithLimit(limit int) DividendsOption {
	return func(d *Dividends) {
		d.limit = &limit
	}
}

func DividendsWithOffset(offset int) DividendsOption {
	return func(d *Dividends) {
		d.offset = &offset
	}
}
