package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup
var resMap = make(map[string]int, 1024)

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
	fileW, err1 := os.OpenFile("result.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	writer := bufio.NewWriter(fileW)
	defer fileW.Close()
	//读取文件
	reader := bufio.NewReader(fileR)
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
		fmt.Printf("url is:%#v\n", url)
		if len(url) == 0 {
			continue
		}
		//链接输出，写到文件里
		wg.Add(1)
		go link(url)
	}
	wg.Wait()
	fmt.Println(len(resMap))
	for k, v := range resMap {
		res := k + "  " + strconv.Itoa(v) + "\r\n"
		writer.WriteString(res)
	}
	writer.Flush()
}

//尝试链接
func link(url string) {
	defer wg.Done()

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	//设置超时为5s
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		resMap[url] = -1
		return
	}
	if resp.StatusCode != 200 {
		fmt.Printf("链接异常, url is:%s, statusCode is : %d, errInfo is :%s", url, resp.StatusCode, err)
		resMap[url] = resp.StatusCode
	}
	return
}
