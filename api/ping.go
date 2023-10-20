package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type PingImpl struct {
}

func NewPing() PingImpl {
	return PingImpl{}
}

func (m PingImpl) Endpoint() (method, path string) {
	method = echo.GET
	path = "/ping"
	return
}

func (m PingImpl) Handler(ctx echo.Context) error {
	apiResponse := "PONG"
	fmt.Println("start:", time.Now())
	return ctx.JSON(http.StatusOK, apiResponse)
}
