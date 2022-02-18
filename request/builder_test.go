package request_test

import (
	"context"
	"errors"
	"testing"

	"github.com/oliread/marketstack/request"
)

func TestNewBuilder(t *testing.T) {
	tcs := []struct {
		name    string
		options []request.BuilderOption
		err     error
	}{
		{
			name:    "NoOptions",
			options: []request.BuilderOption{},
			err:     nil,
		},
		{
			name:    "NilOptions",
			options: nil,
			err:     nil,
		},
		{
			name: "WithBaseURL",
			options: []request.BuilderOption{
				request.BuilderWithBaseURL("valid"),
			},
			err: nil,
		},
		{
			name: "InvalidBaseURL",
			options: []request.BuilderOption{
				request.BuilderWithBaseURL("\x7f"),
			},
			err: request.ErrBuilderOption,
		},
		{
			name: "WithAPIKey",
			options: []request.BuilderOption{
				request.BuilderWithAPIKey("valid"),
			},
			err: nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := request.NewBuilder(tc.options...); tc.err != nil {
				if !errors.Is(err, tc.err) {
					t.Fatalf("EXPECTED: %s\nACTUAL: %s", tc.err, err)
				}
			}
		})
	}
}

func TestBuilderEOD(t *testing.T) {
	tcs := []struct {
		name    string
		options []request.EODOption
		err     error
	}{
		{
			name:    "NoOptions",
			options: []request.EODOption{},
		},
	}

	builder, _ := request.NewBuilder()

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_, err := builder.EOD(context.TODO(), []string{"symbols"})
			if tc.err != nil {
				if !errors.Is(err, tc.err) {
					t.Fatalf("EXPECTED: %s\nACTUAL: %s", tc.err, err)
				}
			}
		})
	}
}
