package tcp

import(
	"strconv"
	"regexp"
)

var pattern string=`(?:\[)l=(\d+)&s=(\d+)&c=(\d+)&f=(\w+)&t=(\w+)(?:\])`

func CheckPacket(content string)(remain string,pks []MyPacket){
	regex,err:=regexp.Compile(pattern)
	
	if err!=nil{
		return content,nil
	}
	
	matchStrs:=regex.FindAllStringSubmatch(content,-1)
	matchIdxs:=regex.FindAllStringSubmatchIndex(content,-1)
		
	matches:=len(matchStrs)
	if matches<1{
		return content,nil
	}
	
	total:=len(content)
	var last string
	packets:=make([]MyPacket,matches)
	count:=0
	index:=0
	
	for k,v:=range matchStrs{
		p:=MyPacket{
			From:v[4],
			To:v[5],
		}
		p.Length,_=strconv.Atoi(v[1])
		p.Sequence,_=strconv.Atoi(v[2])
		p.Command,_=strconv.Atoi(v[3])
		
		startIndex:=matchIdxs[k][1]
		data:=SubString(content,startIndex,p.Length)
		
		if p.Length==len(data){
			p.Data=data
			packets[k]=p
			count++
			index=startIndex+p.Length
		}else{
			break
		}
	}
	
	if count!=matches{
		packets=packets[:count]
	}
	
	if index<total{	
		last=SubString(content,index,total-index)
	}
	
	return last,packets
}

func SubString(str string,begin int,length int) (substr string){
	rs:=[]rune(str)
	lth:=len(rs)
	
	if begin<0{
		begin=0
	}
	
	if begin>=lth{
		begin=lth
	}
	
	end:=begin+length
	if(end>lth){
		end=lth
	}
	return string(rs[begin:end])
}