package redis

import (
	"time"
	"github.com/garyburd/redigo/redis"
)

type RedisPool struct{
	Pool		*redis.Pool
}

//build redis connection pool
func (pool *RedisPool) BuildPool(server,password string,connCount int) {
	pool.Pool = &redis.Pool{
		MaxIdle:connCount,
		IdleTimeout:240*time.Second,
		Dial:func()(redis.Conn,error){
			c,err:=redis.Dial("tcp",server)
			if err != nil{
				return nil,err
			}
			
			if 0!=len(password){
				if _,err:=c.Do("AUTH",password);err != nil{
					c.Close()
					return nil,err
				}
			}
			
			return c,err
		},
		TestOnBorrow:func(c redis.Conn,t time.Time) error{
			_,err := c.Do("PING")
			return err
		},
	}
}

//get redis connection by db no
func (pool RedisPool) GetConn(no int) redis.Conn{
	conn:=pool.Pool.Get()
	conn.Do("SELECT",no)
	return conn
}
