package redis

import (
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

var pool *redisPool

type redisPool struct {
	*redis.Pool
	sync.Mutex
}

func init() {
	config, err := loadConfig()
	if err == nil {
		pool = newPool(&config)
	} else {
		pool = &redisPool{}
	}
}

func newPool(conf *Config) *redisPool {
	pool := &redis.Pool{
		MaxIdle:     conf.MaxIdle,
		MaxActive:   conf.MaxActive,
		IdleTimeout: time.Duration(conf.IdleTimeout) * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp",
				conf.Address,
				redis.DialPassword(conf.Password),
				redis.DialConnectTimeout(time.Duration(conf.ConnTimeout) * time.Second),
				redis.DialReadTimeout(time.Duration(conf.ReadTimeout) * time.Second),
				redis.DialWriteTimeout(time.Duration(conf.WriteTimeout) * time.Second))

			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}

	return &redisPool{Pool: pool}
}

func Get() (redis.Conn, error) {
	if pool != nil {
		return pool.Get(), nil
	}

	pool.Lock()
	defer pool.Unlock()
	if pool != nil {
		return pool.Get(), nil
	}

	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	pool = newPool(&config)

	return pool.Get(), nil
}

func Do(cmd string, args ...interface{}) (interface{}, error) {
	conn, err := Get()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	for i := 0; i < 3; i++ {
		r, e := conn.Do(cmd, args...)
		if e == nil {
			return r, nil
		} else {
			if e == redis.ErrNil {
				return nil, e
			} else {
				err = e
			}
		}
	}

	return  nil, err
}