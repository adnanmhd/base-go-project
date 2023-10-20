package model

import (
	"time"
)

type IAuth interface {
	StoreToken(token string, exp time.Duration)
	GetToken() (string, error)
}
