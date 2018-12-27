package main

import (
	"fmt"
	"net"
	"SimpleChatServer/proto"
)

func registerProcess(conn net.Conn,name,password string) (err error){

	var loginData proto.RegisterReq
	loginData.Name = name
	loginData.Password = password

	err=sendReq(conn,proto.Register,&loginData)

	return
}

//发消息
func process(conn net.Conn) (err error){
	var input string
	fmt.Scanf("%s\n",&input)

	err = msgProcess(conn,input)
	if err != nil {
		fmt.Println("msgProcess error.")
	}


	return
}

func msgProcess(conn net.Conn,msg string) (err error){

	var req proto.MessageReq
	req.Msg =msg
	err=sendReq(conn,proto.Msg,&req)
	return
}


func loginProcess(conn net.Conn,name,password string) (err error){

	var req proto.LoginReq
	req.Name = name
	req.Password = password
	err=sendReq(conn,proto.Login,&req)

	return
}
