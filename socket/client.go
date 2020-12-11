package main

import (
	"fmt"
	"net"
)

func main(){
	conn, err := net.Dial("tcp","127.0.0.1:10000")
	if err != nil {
		fmt.Println(err)
	}
	temp := make([]byte,5)
	temp[0] = byte(1)
	conn.Write(temp)
}
