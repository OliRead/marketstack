package request

import (
	"context"
	"fmt"
	"net/http"
)

type Builder struct {
	apiKey  string
	baseURL string
}

func NewBuilder(opts ...BuilderOption) (*Builder, error) {
	b := Builder{}

	for _, opt := range opts {
		if err := opt(&b); err != nil {
			return nil, fmt.Errorf("%w:%s", ErrBuilderOption, err)
		}
	}

	return &b, nil
}

func (b *Builder) EOD(ctx context.Context, symbols []string, opts ...EODOption) (*http.Request, error) {
	msg := EOD{
		symbols: symbols,
	}

	for _, opt := range opts {
		opt(&msg)
	}

	return b.Build(ctx, &msg)
}

func (b *Builder) Build(ctx context.Context, msg Message) (*http.Request, error) {
	query := msg.Query()
	query.Add("access_key", b.apiKey)

	return http.NewRequestWithContext(ctx, msg.Action().method(), b.baseURL+msg.Action().endpoint(), nil)
}
