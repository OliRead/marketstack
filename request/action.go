package request

import "net/http"

type action byte

const (
	ActionEOD action = iota
)

func (a action) endpoint() string {
	switch a {
	case ActionEOD:
		return "eod"
	default:
		return ""
	}
}

func (a action) method() string {
	switch a {
	case ActionEOD:
		return http.MethodGet
	default:
		return ""
	}
}
