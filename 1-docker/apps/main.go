package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
    kvHostname := os.Getenv("KV_HOSTNAME")
    kv, err := NewKV(kvHostname)

    if err != nil {
        panic(err)
    }

    e := echo.New()

    e.POST("/:key", func(c echo.Context) error {
        key := c.Param("key")
        value := c.FormValue("value")

        err := kv.Set(key, value)
        if err != nil {
            return c.String(http.StatusInternalServerError, err.Error())
        }

        return c.String(http.StatusOK, "OK")
    })

    e.GET("/:key", func(c echo.Context) error {
        key := c.Param("key")

        val, err := kv.Get(key)
        if err != nil {
            return c.String(http.StatusInternalServerError, err.Error())
        }

        return c.String(http.StatusOK, val)
    })

    e.Logger.Fatal(e.Start(":1323"))
}