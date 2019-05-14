package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	//"strings"
	"database/sql"
)



type Ip struct{					//IP结构体
}



type UpIpRequest struct{	//返回ping结果结构体
	Tss  int64
	Src string
	Dst string
	Loss string
	Avg string
	Min string
	Max string
}

type UpIpArrayRequet struct{	//返回ping结果组结构体
	UpIparrayrequet []UpIpRequest
}

type UpIpRespone struct{		//得到IP返回结构体

}
var Upiparrayrequet []UpIpRequest



func (this *Ip)UpIp(upiparrayrequet UpIpArrayRequet,upiprespone *UpIpRespone) error{	//得到IP方法GetIp
	Upiparrayrequet = upiparrayrequet.UpIparrayrequet

	return nil

}


func insertIP(ip []UpIpRequest){		//把客户端的结果存入数据库
	db1,err := sql.Open("mysql","root:123456@tcp(localhost:3306)/ping")
	defer db1.Close()
	for _,singlePing := range ip{
		if err != nil{
			log.Println(err)
		}
		row,err := db1.Prepare(`INSERT valu (tss,src,dst,loss,rttmin,rttavg,rttmax) value (?,?,?,?,?,?,?)`)
		if err != nil{
			log.Println(err)
		}
		row.Exec(singlePing.Tss,singlePing.Src,singlePing.Dst,singlePing.Loss,singlePing.Min,singlePing.Avg,singlePing.Max)

	}
}

func pingList()([]string){			//返回数据库ip地址组
	dba,err :=sql.Open("mysql","root:123456@tcp(localhost:3306)/ping")
	checkError(err)

	defer dba.Close()

	rows,err:=dba.Query("SELECT * FROM host" )
	checkError(err)

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
	lis,err := net.Listen("tcp","172.19.129.11:58099")		//监听端口
	checkError(err)
	defer lis.Close()
	fmt.Fprint(os.Stdout, "%s", "start connection aprilmadaha 58099")
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}

		go func(conn net.Conn) {
			jsonrpc.ServeConn(conn)
			insertIP(Upiparrayrequet)
		}(conn)
	}
}


func checkError(err error) {		//错误函数
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func main(){

	listenIp()

}