package api

import (
	"github.com/labstack/echo/v4"
)

type Api interface {
	Endpoint() (method, path string)
	Handler(ctx echo.Context) error
}
