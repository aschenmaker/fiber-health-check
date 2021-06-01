package healthcheck

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// Test_NoConfig go test -run Test_NoConfig
func Test_NoConfig(t *testing.T) {
	app := fiber.New()
	app.Use(New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusNoContent).SendString("test default")
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/", nil))
	utils.AssertEqual(t, nil, err, "app.Test(req)")
	var content []byte
	content, err = ioutil.ReadAll(resp.Body)
	utils.AssertEqual(t, nil, err, "resp.body(req)")
	utils.AssertEqual(t, fiber.StatusNoContent, resp.StatusCode, "Status Code")
	utils.AssertEqual(t, "", utils.GetString(content), "Textcontent")
}

func Test_NoHeader(t *testing.T) {
	app := fiber.New()
	app.Use(New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusNoContent).SendString("test default")
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "app.Test(req)")
	var content []byte
	content, err = ioutil.ReadAll(resp.Body)
	utils.AssertEqual(t, nil, err, "resp.body(req)")
	utils.AssertEqual(t, fiber.StatusNoContent, resp.StatusCode, "Status Code")
	utils.AssertEqual(t, "", utils.GetString(content), "Textcontent")
}

func Test_Default_Config(t *testing.T) {
	app := fiber.New()
	app.Use(New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusNoContent).SendString("test default")
	})

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set(DefaultHeaderName, DefaultHeaderValue)
	resp, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "app.Test(req)")
	var content []byte
	content, err = ioutil.ReadAll(resp.Body)
	utils.AssertEqual(t, nil, err, "resp.body(req)")
	utils.AssertEqual(t, DefaultResponseCode, resp.StatusCode, "Status Code")
	utils.AssertEqual(t, DefaultResponseText, utils.GetString(content), "Textcontent")
}

func Test_Custom_Config(t *testing.T) {
	app := fiber.New()
	app.Use(New(Config{
		HeaderName:   "X-Custom-Header",
		HeaderValue:  "customValue",
		ResponseCode: http.StatusTeapot,
		ResponseText: "teapot",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusNoContent).SendString("test default")
	})

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Custom-Header", "customValue")
	resp, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "app.Test(req)")
	var content []byte
	content, err = ioutil.ReadAll(resp.Body)
	utils.AssertEqual(t, nil, err, "resp.body(req)")
	utils.AssertEqual(t, http.StatusTeapot, resp.StatusCode, "Status Code")
	utils.AssertEqual(t, "teapot", utils.GetString(content), "Textcontent")
}

func Test_Custom_Config_empty(t *testing.T) {
	app := fiber.New()
	app.Use(New(Config{}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusNoContent).SendString("test default")
	})

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set(DefaultHeaderName, DefaultHeaderValue)
	resp, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "app.Test(req)")
	var content []byte
	content, err = ioutil.ReadAll(resp.Body)
	utils.AssertEqual(t, nil, err, "resp.body(req)")
	utils.AssertEqual(t, DefaultResponseCode, resp.StatusCode, "Status Code")
	utils.AssertEqual(t, DefaultResponseText, utils.GetString(content), "Textcontent")

}
