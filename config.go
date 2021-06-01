package healthcheck

import "net/http"

type Config struct {
	HeaderName   string
	HeaderValue  string
	ResponseCode int
	ResponseText string
}

var (
	// DefaultHeaderName default header name
	DefaultHeaderName = "X-Health-Check"

	// DefaultHeaderValue default header value
	DefaultHeaderValue = "1"

	// DefaultResponseCode default response code
	DefaultResponseCode = http.StatusOK

	// DefaultResponseText default response text
	DefaultResponseText = "ok"

	// DefaultConfig default config
	DefaultConfig = Config{
		HeaderName:   DefaultHeaderName,
		HeaderValue:  DefaultHeaderValue,
		ResponseCode: DefaultResponseCode,
		ResponseText: DefaultResponseText,
	}
)

func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return DefaultConfig
	}
	cfg := config[0]
	if cfg.HeaderName == "" {
		cfg.HeaderName = DefaultHeaderName
	}
	if cfg.HeaderValue == "" {
		cfg.HeaderValue = DefaultHeaderValue
	}
	if cfg.ResponseCode == 0 {
		cfg.ResponseCode = DefaultResponseCode
	}
	if cfg.ResponseText == "" {
		cfg.ResponseText = DefaultResponseText
	}
	return cfg
}
