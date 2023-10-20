package auth

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type auth struct {
	redis *redis.Client
}

func New(redis *redis.Client) auth {
	return auth{redis: redis}
}

var key = "login_auth_token"

func (a auth) StoreToken(token string, expires time.Duration) {
	a.redis.Set(context.Background(), key, token, expires)
}

func (a auth) GetToken() (token string, err error) {
	token, err = a.redis.Get(context.Background(), key).Result()
	return
}
