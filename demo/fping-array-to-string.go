package main

import (
	"fmt"
	"os/exec"
)

func runCommand(name string,arg ...string){
	cmd := exec.Command(name,arg...)
	vv,_ := cmd.CombinedOutput()
	fmt.Println(string(vv))
}

func main(){
	aa := []string{"114.114.114.114","1.1.1.1","8.8.8.8"}
	//fmt.Println(aa)
	//cmd:= exec.Command("fping","-q","-p100","-c5",aa)
	//cmd:= exec.Command("fping",aa...)
	//vv,_ := cmd.CombinedOutput()
	//fmt.Println(string(vv))
//	bbb := append(bbb,"-f")

	runCommand("fping",aa...)
}
