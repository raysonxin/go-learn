package redis


type RedisHelper struct{
	ConnPool RedisPool
}

func (helper RedisHelper) HashSet(hash,key,val string,dbno int)(error){
	conn:=helper.ConnPool.GetConn(dbno)
	defer conn.Close()
	
	_,err:=conn.Do("HSET",hash,key,val)
	return err
}

func (helper RedisHelper) HashDelete(hash,key string,dbno int) error{
	conn:=helper.ConnPool.GetConn(dbno)
	defer conn.Close()
	
	_,err:=conn.Do("HDEL",hash,key)
	return err
}

