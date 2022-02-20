package request

import "net/http"

type action byte

const (
	ActionEOD action = iota
	ActionDividends
)

func (a action) endpoint() string {
	switch a {
	case ActionEOD:
		return "eod"
	case ActionDividends:
		return "dividends"
	default:
		return ""
	}
}

func (a action) method() string {
	switch a {
	case ActionEOD:
		return http.MethodGet
	case ActionDividends:
		return http.MethodGet
	default:
		return ""
	}
}
