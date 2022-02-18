package request

import (
	"errors"
	"net/url"
)

const (
	dateEncodingFormat = "2006-01-02T15:04:05-0700"
)

var (
	ErrBuilderOption = errors.New("builder option")
	ErrSymbols       = errors.New("invalid symbols")
)

type Message interface {
	Query() url.Values
	Action() action
}
