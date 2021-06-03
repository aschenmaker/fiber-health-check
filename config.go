package healthcheck

import "net/http"

// Config defines the config for middleware
type Config struct {
	// HeaderName defines the health check header key
	//
	// Optional. Default: "X-Health-Check"
	HeaderName string
	// HeaderValue defines the health check header val
	//
	// Optional. Default: "1"
	HeaderValue string
	// ResponseCode defines the health check response code
	//
	// Optional. Default: http.StatusOK
	ResponseCode int
	// ResponseText defines the health check response description
	//
	// Optional. Default: "ok"
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

// helper defines the default config for middleware.
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
