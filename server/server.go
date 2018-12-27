package main

import (
	"fmt"
	"io"
	"net"
)


type UsersList struct{
	users []Client
}


func runServer(addr string) (err error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("runServer listen failed, ", err)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("runServer accept failed, ", err)
			continue
		}

		go proccess(conn)
	}
}

func proccess(conn net.Conn) {

	defer conn.Close()
	client := &Client{
		conn: conn,
	}

	err := client.Process()
	if err != nil && err != io.EOF{
		fmt.Println("proccess client failed, ", err)
		return
	}
}
