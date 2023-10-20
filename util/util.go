package util

import (
	"crypto/tls"
	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/gorm"
	"time"
)

func SetupResty(conf ThirdPartyServiceConfig) (restyClient *resty.Client) {
	restyClient = resty.New()
	restyClient.SetTimeout(time.Duration(conf.Timeout) * time.Second)
	restyClient.SetDebug(conf.Debug)
	//restyClient.SetRetryCount(3)
	restyClient.SetRetryWaitTime(time.Duration(conf.Timeout) * time.Second)
	if conf.UseProxy {
		restyClient.SetProxy(conf.HttpProxy)
	}
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	restyClient.SetHeader("Content-Type", "application/json")

	return
}

func SetupRedis(conf Redis) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		Username: conf.Username,
		Password: conf.Password,
	})

	return redisClient
}

func SetDBConn(db DB) *gorm.DB {
	dbConn := NewDBConnection(db)
	dbConn.Begin()
	return dbConn
}
