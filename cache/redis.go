package cache

import (
	"github.com/gomodule/redigo/redis"
	_ "github.com/jinzhu/gorm/dialects/postgres" // For configuration
)

// GetRedisConn returns Redis connection
func GetRedisConn() redis.Conn {
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	return conn
}
