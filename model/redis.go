package model

import (
	"ginTest/config"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

var Redis Rdb

type rdber struct {
	rdbClient *redis.Client
}

type Rdb interface {
	GetRdb() *redis.Client
	Get(key string) interface{}
	Set(key string, value string, expiration time.Duration) bool
	Del(key string) bool
}

func init() {
	Redis = &rdber{}
}

func (r rdber) GetRdb() (goClient *redis.Client) {

	r.rdbClient = redis.NewClient(&redis.Options{
		Addr:     config.Dft.Get().Redis.Host + ":" + strconv.Itoa(config.Dft.Get().Redis.Port),
		Password: config.Dft.Get().Redis.Pass,
		DB:       0, // use default DB
	})

	goClient = r.rdbClient

	return

}

func (r *rdber) Get(key string) (val interface{}) {
	val, err := r.GetRdb().Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			log := "redis key (" + key + ") is not exist"
			println(log)
			return false
		}
		panic(err)
	}
	return
}

func (r *rdber) Set(key string, value string, expiration time.Duration) bool {
	expiration *= 1000000
	Rdb := r.GetRdb()
	setRe, err := Rdb.Set(key, value, expiration).Result()
	if err != nil {
		panic(err)
	}
	if setRe == "OK" {
		println("redis set succeeded")
		return true
	} else {
		println("redis set failed")
		return false
	}
}

func (r *rdber) Del(key string) bool {
	Rdb := r.GetRdb()
	re, err := Rdb.Del(key).Result()
	if err != nil {
		panic(err)
	}
	if re > 0 {
		println("redis del succeeded")
		return true
	} else {
		println("redis del failed or key (" + key + ") not exist")
		return false
	}
}
