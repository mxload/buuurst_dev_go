package main

import (
	"net/http"

	"github.com/labstack/echo"
	buuurst_dev "github.com/mxload/buuurst_dev_go"
)

func main() {
	e := echo.New()

	e.Use(echo.WrapMiddleware(buuurst_dev.MiddlewareFunc(
		&buuurst_dev.BuuurstDevConfig{
			Enabled:      true,
			CollectorURL: "https://lambda-public.buuurst.dev/put-request-log",
			ProjectID:    YOUR_PROJECT_ID,
			ServiceKey:   "YOUR_SERVICE_KEY",
			CustomHeaders: []string{
				"Authorization",
			},
			IgnorePaths: []string{
				"/ignored",
			},
		},
	)))

	e.GET("/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "get")
	})
	e.POST("/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "post")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
