package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"strings"
	"database/sql"
)



type Ip struct{					//IP结构体
}


type GetIpRequest struct {		//得到IP请求结构体

}

type GetIpRespone struct{		//得到IP返回结构体
	Hostip []string				//返回所有IP地址
}

func (this *Ip)GetIp(getiprequest GetIpRequest,getiprespone *GetIpRespone) error{		//上传IP方法UpIp(客户端角度)
    getiprespone.Hostip = pingList()		//返回值从数据库函数getHostip获取
	 return nil
}

func pingList()([]string){			//返回数据库ip地址组
	dba,err :=sql.Open("mysql","root:123456@tcp(localhost:3306)/ping")//填写mysql的账户名密码
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

func insertHostip(ipadd string) {			//将连接服务器的ip放入数据库
	dba,err :=sql.Open("mysql","root:123456@tcp(localhost:3306)/ping")
	checkError(err)

	defer dba.Close()

	row,err := dba.Prepare(`INSERT host (host) value (?)`)
	checkError(err)

	row.Exec(ipadd)
}

func listenIp(){			//多线程监听
	rpc.Register(new(Ip))	//rpc注册IP方法

	lis,err := net.Listen("tcp","172.19.129.11:58098")		//监听端口
	checkError(err)

	defer lis.Close()

	//fmt.Fprint(os.Stdout, "%s", "start connection aprilmadaha 58098")
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
			//fmt.Printf("%s\n", ipadd)

			insertHostip(ipadd)		//先调用插入ip函数

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