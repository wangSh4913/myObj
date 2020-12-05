package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup
var resMap = make(map[string]map[string]int, 1024)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage:jiankong.exe filename")
		os.Exit(0)
	}
	//打开文件
	fileR, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileR.Close()
	//打开写文件
	fileW, err1 := os.OpenFile("result.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 644)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	writer := bufio.NewWriter(fileW)
	defer fileW.Close()
	//读取文件
	reader := bufio.NewReader(fileR)
	runtime.GOMAXPROCS(2)
	i := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		//去除每一行后面的\r\n
		url := strings.TrimRight(line, "\r\n")
		if len(url) == 0 {
			continue
		}
		//fmt.Printf("url is:%#v\n", url)
		res := strings.Split(url, " ")
		name := res[0]
		url = res[1]
		//链接输出，写到文件里
		wg.Add(1)
		go link(name, url)
		i += 1
		if i > 500 {
			fmt.Println(len(resMap))
			for name, m := range resMap {
				for url, retCode := range m{
					res := name + "  " + url + "  " + strconv.Itoa(retCode) + "\r\n"
					writer.WriteString(res)
					delete(resMap, name)
				}
			}
			writer.Flush()
			time.Sleep(10 * time.Second)
			i = 0
		}
	}
	wg.Wait()
}

//尝试链接
func link(name string,url string) {
	defer wg.Done()

	tmp := make(map[string]int,1)
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	//设置超时为5s
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		tmp[url] = -1
		resMap[name] = tmp
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		//fmt.Printf("链接异常, url is:%s, statusCode is : %d, errInfo is :%s", url, resp.StatusCode, err)
		tmp[url] = resp.StatusCode
		resMap[name] = tmp
	}

	return
}
