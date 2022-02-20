package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oliread/marketstack/request"
)

type Client struct {
	httpClient *http.Client
	builder    *request.Builder
}

func NewClient(b *request.Builder) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		builder:    b,
	}
}

func (c *Client) EOD(ctx context.Context, symbols []string, opts ...request.EODOption) (EOD, error) {
	req, err := c.builder.EOD(ctx, symbols, opts...)
	if err != nil {
		return EOD{}, err
	}

	res, err := c.execute(req)
	if err != nil {
		return EOD{}, fmt.Errorf("%w:%s", ErrTransport, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return EOD{}, fmt.Errorf("%w:%s", ErrInternal, err)
	}

	msg := EOD{}
	if err := json.Unmarshal(body, &msg); err != nil {
		return EOD{}, fmt.Errorf("%w:%s", ErrDecode, err)
	}

	return msg, nil
}

func (c *Client) Dividends(ctx context.Context, symbols []string, opts ...request.DividendsOption) (Dividends, error) {
	req, err := c.builder.Dividends(ctx, symbols, opts...)
	if err != nil {
		return Dividends{}, err
	}

	res, err := c.execute(req)
	if err != nil {
		return Dividends{}, fmt.Errorf("%w:%s", ErrTransport, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Dividends{}, fmt.Errorf("%w:%s", ErrInternal, err)
	}

	msg := Dividends{}
	if err := json.Unmarshal(body, &msg); err != nil {
		return Dividends{}, fmt.Errorf("%w:%s", ErrDecode, err)
	}

	return msg, nil
}

func (c *Client) Splits(ctx context.Context, symbols []string, opts ...request.SplitsOption) (Splits, error) {
	req, err := c.builder.Splits(ctx, symbols, opts...)
	if err != nil {
		return Splits{}, err
	}

	res, err := c.execute(req)
	if err != nil {
		return Splits{}, fmt.Errorf("%w:%s", ErrTransport, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Splits{}, fmt.Errorf("%w:%s", ErrInternal, err)
	}

	msg := Splits{}
	if err := json.Unmarshal(body, &msg); err != nil {
		return Splits{}, fmt.Errorf("%w:%s", ErrDecode, err)
	}

	return msg, nil
}

func (c *Client) execute(req *http.Request) (*http.Response, error) {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w:%s", ErrTransport, err)
	}

	return res, nil
}
