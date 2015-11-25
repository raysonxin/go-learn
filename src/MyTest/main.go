// MyTest project main.go
package main

import (
	"fmt"
	"atisafe/redis"
)

func main() {
	
	p:=redis.RedisPool{}
	
	p.BuildPool("localhost:6379","",10)
	
	h:=redis.RedisHelper{
		ConnPool:p,
	}
	
	h.HashSet("rayson","age","27",0)
	
	h.HashDelete("rayson","name",0)
	
	fmt.Println("Hello World!")
}
