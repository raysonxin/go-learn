package redis


type RedisHelper struct{
	ConnPool RedisPool
}

//设置hash的某个值，hash:key:value
func (helper RedisHelper) HashSet(dbno int,hash string,key string,val interface{})(error){
	
	conn:=helper.ConnPool.GetConn(dbno)
	defer conn.Close()
	
	_,err:=conn.Do("HSET",hash,key,val)
	return err
}

//删除hash的某个值，hash:key
func (helper RedisHelper) HashDelete(dbno int,hash string,keys ...interface{}) error{
	conn:=helper.ConnPool.GetConn(dbno)
	defer conn.Close()
	
	_,err:=conn.Do("HDEL",hash,keys)
	return err
}

//移除整个hash内存储的内容，hash
func (helper RedisHelper) KeyDelete(dbno int,hash string) error{
	conn:=helper.ConnPool.GetConn(dbno)
	defer conn.Close()
	
	_,err:=conn.Do("DEL",hash)
	return err
}

//获取hash中key存储的值
func (helper RedisHelper) HashGet(dbno int,hash string,key string) (interface{},error){
	conn:=helper.ConnPool.GetConn(dbno)
	defer conn.Close()
	
	res,err:=conn.Do("HGET",hash,key)
	return res,err
}
 
