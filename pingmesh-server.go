package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"strings"
	"database/sql"
	//"encoding/json"
)

var Upiparrayrequet UpIpArrayRequet

type Ip struct{					//IP结构体
}


type GetIpRequest struct {		//得到IP请求结构体

}

type GetIpRespone struct{		//得到IP返回结构体
	Hostip []string				//返回所有IP地址
}


type UpIpRequest struct{	//返回ping结果结构体
	Dst,Ploss,Pavg string		//结构体变量 Dst:目的地址,Ploss:丢包率,Pavg:平均丢包率
}

type UpIpArrayRequet struct{	//返回ping结果组结构体
	Src string			`json:"Src"`		//解析Src json
	Tss string			`json:"Tss"`		//解析Tss json
	Upiprequest	[]UpIpRequest	`json:"UpIpArrayRequet"`			//解析PING结果 json
}

type UpIpRespone struct{		//得到IP返回结构体

}


func (this *Ip)GetIp(getiprequest GetIpRequest,getiprespone *GetIpRespone) error{		//上传IP方法UpIp(客户端角度)
    getiprespone.Hostip = pingList()		//返回值从数据库函数getHostip获取
	 return nil
}

func (this *Ip)UpIp(upiparrayrequet UpIpArrayRequet,upiprespone *UpIpRespone) error{	//得到IP方法GetIp
	Upiparrayrequet = upiparrayrequet
	return nil
}

func insertHostip(ipadd string) {			//将连接服务器的ip放入数据库
	dba,err :=sql.Open("mysql","root:123456@tcp(localhost:3306)/ping")
	defer dba.Close()
	row,err := dba.Prepare(`INSERT host (host) value (?)`)

	if err != nil{
		log.Println(err)
	}
	row.Exec(ipadd)
}

func pingList()([]string){			//返回数据库ip地址组
	dba,err :=sql.Open("mysql","root:123456@tcp(localhost:3306)/ping")
	defer dba.Close()
	rows,err:=dba.Query("SELECT * FROM host" )
	log.Println(err)

	var hostIp string
	var ipArray []string		//保存ip数组
	for rows.Next() {			//数据库查询
		err := rows.Scan(&hostIp)
		log.Println(err)
		ipArray = append(ipArray,hostIp)
	}
	return ipArray				//返回地址组
}


func listenIp(){			//多线程监听
	rpc.Register(new(Ip))	//rpc注册IP方法

	lis,err := net.Listen("tcp","127.0.0.1:8098")		//监听端口
	if err != nil {
		log.Fatalln("fatal error: ", err)

	}
	defer lis.Close()
	//getHostip()


	fmt.Fprint(os.Stdout, "%s", "start connection BBBBBBBBAAAA")
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}

		go func(conn net.Conn) {
			fmt.Fprintf(os.Stdout, "%s", "new client in comming\n")
			//fmt.Println(conn.RemoteAddr())
			jsonrpc.ServeConn(conn)
			ipaddress := conn.RemoteAddr().String()
			ipadd := strings.Split(ipaddress, ":")[0]
			fmt.Printf("%s\n", ipadd)

			insertHostip(ipadd)		//先调用插入ip函数
			//getHostip()
		}(conn)
	}
}

func main(){

	listenIp()
}