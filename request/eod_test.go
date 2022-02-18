package request_test

import (
	"errors"
	"testing"

	"github.com/oliread/marketstack/request"
)

func TestNewEOD(t *testing.T) {
	tcs := []struct {
		name    string
		symbols []string
		err     error
	}{
		{
			name:    "Success",
			symbols: []string{"SOME, SYMBOLS"},
			err:     nil,
		},
		{
			name:    "NoSymbols",
			symbols: []string{},
			err:     request.ErrSymbols,
		},
		{
			name:    "NilSymbols",
			symbols: nil,
			err:     request.ErrSymbols,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_, err := request.NewEOD(tc.symbols)
			if err != nil {
				if !errors.Is(err, tc.err) {
					t.Fatalf("EXPECTED: %s\nACTUAL: %s", tc.err, err)
				}
			}
		})
	}
}
