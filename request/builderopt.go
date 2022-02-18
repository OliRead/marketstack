package request

import "net/url"

type BuilderOption func(*Builder) error

func BuilderWithAPIKey(key string) BuilderOption {
	return func(b *Builder) error {
		b.apiKey = key

		return nil
	}
}

func BuilderWithBaseURL(u string) BuilderOption {
	return func(b *Builder) error {
		if _, err := url.Parse(u); err != nil {
			return err
		}

		b.baseURL = u

		return nil
	}
}
