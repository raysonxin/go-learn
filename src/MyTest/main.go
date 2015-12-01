// MyTest project main.go
package main

import (
	"fmt"
	"atisafe/redis"
	rg "github.com/garyburd/redigo/redis"
)

func main() {
	
	p:=redis.RedisPool{}
	
	p.BuildPool("localhost:6379","",10)
	
	h:=redis.RedisHelper{
		ConnPool:p,
	}
	
	h.HashSet(0,"rayson","name","abc")
	
	//h.HashDelete(0,"rayson","name")
	
	res,_:=rg.String(h.HashGet(0,"rayson","name"))
	
	
	fmt.Println(res)
}
