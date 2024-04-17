package main

import (
	"log"
	"os/exec"
	"strings"
	"time"
)

var ifnotresponding bool = false
var pid string

func check_unresponding(processname string) {
	println("start check")
	command := []string{
		"/V",
		"/FI",
		"imagename eq " + processname,
	}
	cmd := exec.Command("tasklist", command...)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("result is %s\n", &out)

	parseoutput(string(out))
}

func parseoutput(output string) {
	// println(output)
	parse1 := strings.Fields(output)
	if len(parse1) == 2 {
		println("未找到相应进程，请启动")
		// os.Exit(0)
	} else {
		pid = parse1[20]
		// println(pid)
		ifnotresponding = (parse1[25] != "Running")
	}
	// fmt.Printf("%q\n", parse1)

	// println(ifnotresponding)
}

func killprocess(processname string) {
	command := []string{
		"/F",
		"/FI",
		"imagename eq " + processname,
	}
	cmd := exec.Command("taskkill", command...)
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	// println(cmd.String())
}

// func restart(processname string, filelocation string) {
// 	cmd := exec.Command(filelocation)
// 	println(cmd.String())
// 	if err := cmd.Run(); err != nil {
// 		log.Fatalln(err)
// 	}
// }

func main() {
	var processname = "IdleDragons.exe"
	// var filelocation = "D:\\Epic games\\IdleChampions\\IdleDragons.exe"
	// killprocess(processname)
	//restart(processname, filelocation)
	println("每10s检查IdleDragons.exe进程是否未响应，如未响应则强制杀死进程")
	for true {
		check_unresponding(processname)
		if ifnotresponding {
			println("检测到程序未响应，将强制关闭")
			killprocess(processname)
		}
		time.Sleep(10 * time.Second)
	}
}
