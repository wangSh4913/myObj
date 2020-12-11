package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

/*
 flag包的用法：
	flag.BoolVar(&onCpu,    "cpu",    false,    "turn on cpu monitor")
                 结果存放  命令行参数     默认值       提示信息
    e.g   1)go run main.go  --help
         -cpu
              turn on cpu?
          2)go run main.go -cpu=true
 pprof包用法：
     pprof.StartCPUProfile(file)
     写到文件后，使用
     go tool pprof filename
*/

func test(){
	var ch chan bool
	for{
		v := <-ch
		switch {
		case v:
			fmt.Printf("111")
		default:
		}
	}
}

func main() {
	//两种写法
	var onCpu bool
	var onMem *bool
	flag.BoolVar(&onCpu, "cpu", false, "turn on cpu monitor")
	onMem = flag.Bool("mem", false, "turn on memory monitor")
	//必须写，否则无法解析到
	flag.Parse()

	for i:=0;i<6;i++{
		go test()
	}


	if onCpu {
		//启动cpu监控，输出到文件
		file, _ := os.Create("cpu.pprof")
		pprof.StartCPUProfile(file)
		defer file.Close()
		defer pprof.StopCPUProfile()
	}
	if *onMem {
		//启动cpu监控，输出到文件
		file1, _ := os.Create("mem.pprof")
		pprof.WriteHeapProfile(file1)
		defer file1.Close()
	}
	time.Sleep(20*time.Second)
}
