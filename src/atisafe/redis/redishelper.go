package redis


var (
	connPool RedisPool
)

func AddHash(hash,val string,dbno int){
	conn:=connPool.GetConn(dbno)
	conn.Do("HSET",hash,val)
}