package main

import (
	"fmt"
	"net"
)

func main(){
	listen, err := net.Listen("tcp","127.0.0.1:10000")
	if err != nil {
		fmt.Println(err)
	}
	conn,err := listen.Accept()
	if err != nil {
		fmt.Println(err)
	}
	var temp [128]byte
	n, err := conn.Read(temp[:])
	if n > 0 {
		fmt.Println(temp[:n])
	}
}
