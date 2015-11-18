package main

import (
	"time"
	"fmt"

	"github.com/shirou/gopsutil/mem"
"github.com/shirou/gopsutil/cpu"
"github.com/shirou/gopsutil/process"
)

func main() {
	v, _ := mem.VirtualMemory()
	infos := MonitorInfo{}
	infos.MemTotal = v.Total
	infos.MemUsedPercent = v.UsedPercent
	fmt.Println(infos)

	a,_:=cpu.CPUTimes(false)
    fmt.Println(a)
pro,_:=process.NewProcess(2716)
//pro.Kill()
//for true{


b,_:=pro.CPUPercent()
    fmt.Println(b)
	time.Sleep(10000)
}


}
