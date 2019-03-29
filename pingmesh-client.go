package main

import (
	"os/exec"
	"strings"
	"fmt"
	"net"
	"os"
	"time"
)



type Pings struct{			//申明结构体
	Dip,Packetloss,Pavg string
	//Pavg string
}


func GetIp()(string){ 		//获取本地ip
	var IP string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//fmt.Println(ipnet.IP.String())
				IP = ipnet.IP.String()
			}
		}
	}
	return IP
}

func Ping(ipadd string,PingChan chan Pings){	//获取目标ip,丢包率，ping平均延迟
	result  := exec.Command("ping","-q" ,"-c 4" ,"-i 0.2", ipadd)	//执行ping命令
	out,_ := result.Output()

	r := strings.Split(string(out),"---")				//根据结果进行分割


	//fmt.Println(r[2])
	dipString := strings.Split(r[1]," ")				//获取目的ip地址
	dip := dipString[1]
	//fmt.Println(dip)

	plString := strings.Split(r[2],",")				//获取丢包率
	pkloss :=strings.Split(plString[2]," ")[1]
	packetloss := strings.Trim(pkloss,"%")
	//fmt.Println(packetloss)

	pavgString := strings.Split(plString[2]," ")[6]	//获取ping的平均丢包率
	pavg := strings.Split(pavgString,"/")[1]
	//fmt.Println(pavg)


	PingR := Pings{Dip:dip,Packetloss:packetloss,Pavg:pavg}	//将数据保存到结构体上
	PingChan <- PingR										//将数据保存到通道上
	//return dip,packetloss,pavg


}

func main(){
	//abc := "114.114.114.114"
	t_start := time.Now()  	// 运行程序前先记录时间
	PingChan := make(chan Pings,10)		//通道初始化
	for i:=0;i<20;i++{
		go Ping("172.31.160.114",PingChan)
	}

	//fmt.Println(GetIp())
	for i:=0;i<20;i++{
		fmt.Println(<-PingChan)			//从通道/信道里面读取数据
	}
	t_end := time.Now()

	//fmt.Println(sum)
	fmt.Println(t_end.Sub(t_start))
	fmt.Println(GetIp())
	//fmt.Println(net.Dial("icmp","114.114.114.114"))
}