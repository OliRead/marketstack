package request

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Splits struct {
	symbols  []string
	sort     *SortOrder
	dateFrom *time.Time
	dateTo   *time.Time
	limit    *int
	offset   *int
}

type SplitsOption func(*Splits)

func NewSplits(symbols []string, opts ...SplitsOption) (*Splits, error) {
	if len(symbols) == 0 {
		return nil, fmt.Errorf("%w:no symbols", ErrSymbols)
	}

	msg := Splits{
		symbols: symbols,
	}

	for _, opt := range opts {
		opt(&msg)
	}

	return &msg, nil
}

func (s *Splits) Query() url.Values {
	v := url.Values{}
	v.Add("symbols", strings.Join(s.symbols, ","))

	if s.sort != nil {
		v.Add("sort", s.sort.String())
	}

	if s.dateFrom != nil {
		v.Add("date_from", s.dateFrom.Format(dateEncodingFormat))
	}

	if s.dateTo != nil {
		v.Add("date_to", s.dateFrom.Format(dateEncodingFormat))
	}

	if s.limit != nil {
		v.Add("limit", strconv.Itoa(*s.limit))
	}

	if s.offset != nil {
		v.Add("offset", strconv.Itoa(*s.offset))
	}

	return v
}

func (s *Splits) Action() action {
	return ActionDividends
}

func SplitsWithSortOrder(order SortOrder) SplitsOption {
	return func(s *Splits) {
		s.sort = &order
	}
}

func SplitsWithDateRange(from, to *time.Time) SplitsOption {
	return func(s *Splits) {
		if s.dateFrom != nil {
			s.dateFrom = from
		}

		if s.dateTo != nil {
			s.dateTo = to
		}
	}
}

func SplitsWithLimit(limit int) SplitsOption {
	return func(s *Splits) {
		s.limit = &limit
	}
}

func SplitsWithOffset(offset int) SplitsOption {
	return func(s *Splits) {
		s.offset = &offset
	}
}
