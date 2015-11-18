// LearnGo project LearnGo.go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mysql"
)

func main() {
	db,err:=sql.Open("mysql","root:ty123456@tcp(192.168.1.70:3306)/safeguard?charset=utf8")
	if err!=nil{
		panic(err.Error())
	}
	defer db.Close()
	
	rows,err:=db.Query("select * from users")
	if err!=nil{
		panic(err.Error())
	}
	
	columns,err :=rows.Columns()
	if err!=nil{
		panic(err.Error())
	}
	
	values:=make([]sql.RawBytes,len(columns))
	scanArgs:=make([]interface{},len(values))
	for i:=range values{
		scanArgs[i]=&values[i]
	}
	
	for rows.Next(){
		err=rows.Scan(scanArgs...)
		if err!=nil{
			panic(err.Error())
		}
		
		var value string
		for i,col:=range values{
			if col==nil{
				value="NULL"
			}else{
				value=string(col)
			}
			fmt.Println(columns[i],":",value)
		}
	}
}

