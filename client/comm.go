package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"SimpleChatServer/proto"
)



func sendReq(conn net.Conn,cmd int,req interface{})(err error){
	var msg proto.Message
	msg.Cmd = cmd

	jsondata,err := json.Marshal(req)
	if err != nil{
		fmt.Println("goCallApi json masrshal error.")
	}
	msg.Data = string(jsondata)

	MsgJson,err := json.Marshal(msg)
	if err != nil{
		fmt.Println("goCallApi json masrshal error.")
	}

	packLen := uint32(len(MsgJson))
	var buf =make([]byte,8)
	binary.BigEndian.PutUint32(buf[0:4], packLen)
	_, err = conn.Write(buf[0:4])
	if err != nil {
		fmt.Println("write data  failed")
		return
	}

	conn.Write(MsgJson)
	if err != nil {
		fmt.Println("goCallApi write error.err:",err.Error())
	}



	return
}