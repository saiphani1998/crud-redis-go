package redis

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	_ "sync"
	"time"
	"fmt"
)

//types
type (
	Course struct {
		Department string
		Code       string
		Section    string
	}

	Student struct {
		Id string `json:"id"`
	}

	PoolConn struct {
		pool *redis.Pool
	}
)

func New() *PoolConn {
	return &PoolConn{
		pool: &redis.Pool{
			MaxIdle:     80,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", "localhost:6379")
				if err != nil {
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		},
	}
}

//keep hmap with course -> code ->receivers map([int]Receivers)
func (c *PoolConn) Add(course Course, receiver string) error {
	key, err := json.Marshal(course)
	if err != nil {
		return err
	}
	conn := c.pool.Get()
	defer conn.Close()
	if _, err := conn.Do("LPUSH", string(key), receiver); err != nil {
		return err
	}
	return nil
}

func (c *PoolConn) Remove(course *Course, receiver string) error {
	key, err := json.Marshal(course)
	if err != nil {
		return err
	}
	conn := c.pool.Get()
	defer conn.Close()
	if _, err := conn.Do("LREM", string(key), -1, receiver); err != nil {
		return err
	}
	return nil
}

func (c *PoolConn) Get(course Course) ([]string, error) {
	out, err := json.Marshal(course)
	if err != nil {
		return nil, err
	}
	conn := c.pool.Get()
	defer conn.Close()
	values, err := redis.Strings(conn.Do("LRANGE", string(out), 0, -1))
	fmt.Println("The ids of the registered students are: ",values)	
	if err != nil {
		return nil, err
	}
	return values, nil
}

