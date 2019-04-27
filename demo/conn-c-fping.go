package main

import (

	"log"
	"net/rpc/jsonrpc"
	"fmt"
	"net"
	"os"
	//"os/exec"
	//"strings"
	"time"
	//"os/exec"

	//"os/exec"
	"os/exec"
	//"strings"
	"regexp"
	"strings"
	//"encoding/json"
)

type GetIpRequest struct {		//得到IP请求结构体

}

type GetIpRespone struct{		//得到IP返回结构体
	Hostip []string				//返回所有IP地址
}

type Pingstruct struct {
	Tss int64
	Src,Dst,Loss,Avg string
}

type UpIpArrayRequet struct{	//返回ping结果组结构体
	UpIparrayrequet []Pingstruct
}

type UpIpRespone struct{		//得到IP返回结构体

}
//
//type PingStructGroup struct {
//	Src,Tss string
//	PingStructArray	[]PingStruct
//}
//
//var PingStructArray []PingStruct
////var ccc chan int
////ccc := make(chan int,10)

var pingStructArray []Pingstruct

func GetLocalIp()(string){ 		//获取本地ip
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

func runCommand(name string,arg ...string){		//运行fping命令
	var pingstruct Pingstruct
	var Pingstructlist []Pingstruct
	arg = append(arg,"-q")
	arg = append(arg,"-p")
	arg = append(arg,"120")
	arg = append(arg,"-c")
	arg = append(arg,"5")
	cmd := exec.Command(name,arg...)
	vv,_ := cmd.CombinedOutput()
	re := regexp.MustCompile(`(.*) +: xmt/rcv/%loss = (.*), min/avg/max = (.*)`) //正则过滤下丢包率为100%的
	submatchall := re.FindAllStringSubmatch(string(vv),-1)
	for _, element := range submatchall {
		//fmt.Println(element[0])
		//fmt.Println(element[1])
		//fmt.Println(strings.Split(element[2],"/")[2])
		//fmt.Println(strings.Split(element[3],"/")[2])
		pingstruct.Src = GetLocalIp()
		pingstruct.Tss = time.Now().Unix()
		pingstruct.Dst = element[1]
		pingstruct.Loss = strings.Split(element[2],"/")[2]
		pingstruct.Avg = strings.Split(element[3],"/")[1]
		Pingstructlist = append(Pingstructlist,pingstruct)
	}
	//res,_ := json.Marshal(Pingstructlist)
	pingStructArray = Pingstructlist
//	return res
//	fmt.Println(res)
	fmt.Println(pingStructArray)

}


func fPing(ipadd []string){	//获取目标ip,丢包率，ping平均延迟

	runCommand("fping",ipadd...)
//	fmt.Println(ipadd)
	////fmt.Println(b)
	//aa := []string{ipadd}
	////args := strings.Fields(ipadd)
	//result  := exec.Command("fping","-A" ,"-c 4" ,"-u", ipadd)	//执行ping命令
	//out,_ := result.Output()
	//fmt.Println(out)
}

func pingHost()([]string){
	//wg := sync.WaitGroup{}

	conn,err := jsonrpc.Dial("tcp","127.0.0.1:8098")
	if err != nil{
		log.Fatalln("dailing error:",err)
	}

	getiprequest :=GetIpRequest{}
	var getiprespone GetIpRespone
	err = conn.Call("Ip.GetIp",getiprequest,&getiprespone)
	if err != nil{
		log.Fatalln("getip error: ",err)
	}


	return getiprespone.Hostip

}

func UpIp(){
	conn,err := jsonrpc.Dial("tcp","127.0.0.1:8098")
	if err != nil{
		log.Fatalln("dailing error:",err)
	}

	upip := UpIpArrayRequet{pingStructArray}
	var rippr UpIpRespone
	err  = conn.Call("Ip.UpIp",upip,&rippr)
	if err != nil{
		log.Fatalln("return error: ",err)
	}

}


func main(){
	t_start := time.Now()

	//for t := range time.Tick(time.Second * 60) {
	//	//fmt.Println(t, "hello world")
	//	fPing(pingHost())
	//	fmt.Println(t)
	//}
	fPing(pingHost())
	UpIp()
	//fmt.Println(strings.Split(pingHost(),"\n"))
	t_end := time.Now()
	fmt.Println(t_end.Sub(t_start))
	}

