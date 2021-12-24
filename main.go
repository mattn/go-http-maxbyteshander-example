package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func middleware(next http.Handler) http.Handler {
	return http.MaxBytesHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}), 4096)
}

func main() {
	e := echo.New()

	e.Use(echo.WrapMiddleware(middleware))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})

	e.POST("/", func(c echo.Context) error {
		var v interface{}
		err := json.NewDecoder(c.Request().Body).Decode(&v)
		if err != nil {
			return err
		}
		fmt.Println(v)
		return c.String(http.StatusOK, "")
	})
	e.Logger.Fatal(e.Start(":8989"))
}
