package main


import (
	"os/exec"
	"fmt"
	"time"
	"regexp"
)
type Pingstruct struct{
	Dst string
	Avg float64
}

func main(){
	t_start := time.Now()
	cmd:= exec.Command("fping","-q","-p100","-c5","-f","/home/aprilmadaha/go/src/pingmesh/ip.txt")

	vv,_ := cmd.CombinedOutput()
	//abc :=strings.Split(string(vv),"\n")
	//r, _ := regexp.Compile("p([a-z]+)ch")
	//res = re.match('.+received, (\d+)% packet loss, time (.*)rtt min/avg/max/mdev = (.*) ms',r).groups()
	//#丢包率, 4次ping总耗时, 平均响应时间
	//packet_loss, total_time, avg_time = res[0], res[1], res[2].split('/')[1]

	//114.114.114.114 : xmt/rcv/%loss = 5/5/0%, min/avg/max = 12.5/12.8/13.0
	//125.39.52.26    : xmt/rcv/%loss = 5/5/0%, min/avg/max = 8.57/9.04/9.64
	//123.103.122.24  : xmt/rcv/%loss = 5/5/0%, min/avg/max = 6.53/6.97/7.30
	//140.205.220.96  : xmt/rcv/%loss = 5/5/0%, min/avg/max = 28.7/29.4/30.3
	//1.1.1.1         : xmt/rcv/%loss = 5/0/100%


		r, _ := regexp.Compile(".+xmt/rcv/%loss.+.min/avg/max = (.*)")
	fmt.Println(r.FindAllString(string(vv), -1)[1])

	//fmt.Println(abc)
	//fmt.Println(len(abc))
	//for i,cba := range abc{
	//	fmt.Println(i,cba)
	//}


	t_end := time.Now()
	fmt.Println(t_end.Sub(t_start))
}
