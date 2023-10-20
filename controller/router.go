package controller

import (
	"base-go-project/api"
	"base-go-project/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ListApi []api.Api

func SetupRouter(conf util.APPConfig) *echo.Echo {
	//var dbConn *gorm.DB
	//
	//defer func() {
	//	dbConn.Close()
	//}()

	router := setMiddleware()

	// enable these functions below for client connection: db, redis, http

	//restyClient := util.SetupResty(conf.DigitalService)
	//redisClient := util.SetupRedis(conf.Redis)
	//dbConn = util.SetDBConn(conf.Database)

	// registry all your endpoint
	listApi := ListApi{
		api.NewPing(),
	}

	// no need to change this code below
	for _, api := range listApi {
		method, path := api.Endpoint()
		router.Add(method, path, api.Handler)
	}

	return router
}

func setMiddleware() (router *echo.Echo) {
	router = echo.New()
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri}, status=${status}, ${latency_human}\n",
	}))
	router.Use(middleware.BodyLimit("2M"))
	router.Use(middleware.Recover())
	return
}
