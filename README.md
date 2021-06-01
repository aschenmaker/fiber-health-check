# Fiber-health-check 🩺 Middleware
🩺Fiber-health-check middleware support health-check for [Fiber](https://github.com/gofiber/fiber)⚡️ framework.

## Signatures

```go
func New(config ...Config) fiber.Handler
```

## How to use
Frist import the midllware from Github,
```shell
go get -u github.com/gofiber/v2
go get -u github.com/aschenmaker/fiber-health-check
```
Then create a Fiber app with `app := fiber.New()`.
## Default Config
```go
app.Use(healthcheck.New())
```
## Custom Config
```go
app.Use(healthcheck.New(
	HeaderName:   "X-Custom-Header",
	HeaderValue:  "customValue",
	ResponseCode: http.StatusTeapot,
	ResponseText: "teapot",
))
```
## Config
```go
// Config defines the config for middleware
type Config struct {
    // HeaderName defines the health check header key
    //
    // Optional. Default: "X-Health-Check"
    HeaderName   string
    // HeaderValue defines the health check header val
    //
    // Optional. Default: "1"
    HeaderValue  string
    // ResponseCode defines the health check response code
    //
    // Optional. Default: http.StatusOK
    ResponseCode int
    // ResponseText defines the health check response description
    //
    // Optional. Default: "ok"
    ResponseText string
}
```
