package request

import "net/http"

type action byte

const (
	ActionEOD action = iota
	ActionDividends
	ActionSplits
)

func (a action) endpoint() string {
	switch a {
	case ActionEOD:
		return "eod"
	case ActionDividends:
		return "dividends"
	case ActionSplits:
		return "splits"
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
	case ActionSplits:
		return http.MethodGet
	default:
		return ""
	}
}
