package dao

import (
	rdx "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	redis *rdx.Client
)
