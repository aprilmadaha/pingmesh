package main


import (
	"os/exec"
	"fmt"
	"time"
	//"regexp"
	"regexp"
	"strings"
)

func main(){
	t_start := time.Now()
	cmd:= exec.Command("fping","-q","-p100","-c5","-f","/Users/aprilmadaha/go/src/pingrpc/fping/ping1.txt")

	vv,_ := cmd.CombinedOutput()
	//cmd = "sshpass -p '{0}' ssh {1}@{2} 'ping -q -c 4 -i 0.2 -w 1 {3}'".format(password, username, s_ip, r_ip)
	//r = os.popen(cmd).read().replace('\n', '')

	//aprilmadaha@liy:~/go/src/golang.org/x$ ping -q -c4 -i 0.2 114.114.114.114
	//PING 114.114.114.114 (114.114.114.114): 56 data bytes
	//
	//--- 114.114.114.114 ping statistics ---
	//	4 packets transmitted, 4 packets received, 0.0% packet loss
	//round-trip min/avg/max/stddev = 23.274/23.686/23.925/0.247 ms

	//res = re.match('.+received, (\d+)% packet loss, time (.*)rtt min/avg/max/mdev = (.*) ms',r).groups()
	//#丢包率, 4次ping总耗时, 平均响应时间
	//packet_loss, total_time, avg_time = res[0], res[1], res[2].split('/')[1]
	//#平均响应时间以0-10ms为基准,换算成百分制(视情况而定，此处延迟都较小，为了页面便于区分故×10)
	//value = int(float(avg_time)*10)
	//title = '丢包率:'+packet_loss+'%，4次ping总耗时:'+total_time+'，平均响应时间:'+avg_time+'ms.'
     //       mesh_data[key] = [value, title]s[0], res[1], res[2].split('/')[1]

	//114.114.114.114 : xmt/rcv/%loss = 5/5/0%, min/avg/max = 12.5/12.8/13.0
	//125.39.52.26    : xmt/rcv/%loss = 5/5/0%, min/avg/max = 8.57/9.04/9.64
	//123.103.122.24  : xmt/rcv/%loss = 5/5/0%, min/avg/max = 6.53/6.97/7.30
	//140.205.220.96  : xmt/rcv/%loss = 5/5/0%, min/avg/max = 28.7/29.4/30.3
	//1.1.1.1         : xmt/rcv/%loss = 5/0/100%

	//liy := regexp.MustCompile(`^http://www.liy.org/([\d]{4})/([\d]{2})/([\d]{2})/([\w-]+).html$`)
	//params := liy.FindStringSubmatch("http://www.liy.org/2018/01/20/golang-goquery-examples-selector.html")
	//
	//for _,param :=range params {
	//	fmt.Println(param)
	//}

//	r:= regexp.MustCompile(`114.114.114.114 : xmt/rcv/%loss = 5/5/0%, min/avg/max = 12.5/12.8/13.0`)
//	r:= regexp.MustCompile(`(\d+).+xmt/rcv/%loss+.+min/avg/max =(.*)`)
//	r:= regexp.MustCompile(`(\d+).+xmt/rcv/%loss+.+min/avg/max =(.*)`)
//	r:= regexp.MustCompile(`(2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})(\\.(2(5[0-5]{1}|[0-4]\\d{1})|[0-1]?\\d{1,2})){3}`)
//	fmt.Println(r.FindAllString(string(vv), -1))
//	fmt.Println(string(vv))
	//fmt.Println(strings.Split(string(vv),"\n"))
	aaa := strings.Split(string(vv),"\n")
	for _,bb := range aaa{
		fmt.Println(bb)
	}

	//re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3} +: xmt/rcv/%loss = (.*), min/avg/max = (.*)`)
	re := regexp.MustCompile(`(.*) +: xmt/rcv/%loss = (.*), min/avg/max = (.*)`)
	//re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
	//	fmt.Println(matched, err)
	submatchall := re.FindAllString(string(vv), -1)
	for _, element := range submatchall {
		//fmt.Println(element)
		submatchall1 := re.FindStringSubmatch(element)
		for _, element1 := range submatchall1 {
			fmt.Println(element1)
		}

	}


	t_end := time.Now()
	fmt.Println(t_end.Sub(t_start))
}
