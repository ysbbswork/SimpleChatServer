package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"SimpleChatServer/proto"
)

var(
	alreadyLogin bool

)

func main() {
	clientProcess("localhost:10000")
}

func clientProcess(address string) {


	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	fmt.Println(address)
	go handleMsg(conn)
	alreadyLogin = false

	for {
		if !alreadyLogin {
			userInit(conn)
		}
		res :=process(conn)
		if res != nil{
			return
		}
	}


}

func loginInit(conn net.Conn) (err error){
	var name,password string
	fmt.Printf("请输入注册用户名和密码\n")
	_, err = fmt.Scanf("%s\n%s\n", &name, &password)
	if err != nil{
		fmt.Println("scanf error.",err.Error())
	}
	err = loginProcess(conn,name,password)
	if err != nil {
		fmt.Println("login error.")
	}
	return
}
func registerInit(conn net.Conn)(err error){
	var name,password string
	fmt.Printf("请输入注册用户名和密码\n")
	_, err = fmt.Scanf("%s\n%s\n", &name, &password)
	if err != nil{
		fmt.Println("scanf error.",err.Error())
	}

	err = registerProcess(conn,name,password)
	if err != nil {
		fmt.Println("registerProcess error.")
	}
	return
}
func userInit(conn net.Conn){
	fmt.Printf("登录：1\t注册用户：2\n")
	var number int
	_, err:= fmt.Scanf("%d\n", &number)
	if err != nil {
		fmt.Println("scanf error",err.Error())
	}
	switch number{
	case 1:
		loginInit(conn)
	case 2:
		registerInit(conn)
	default:
		fmt.Println("选项非法！")
	}
	return
}


func handleMsg(conn net.Conn) {

	for {
		err := listenAndHandle(conn)
		if err != nil{
			fmt.Println(err.Error())
		}
	}
	return
}


//监听消息，并处理
func listenAndHandle(conn net.Conn)(err error) {
	msg, err := readPack(conn)
	fmt.Println("cmd", msg.Cmd)
	switch msg.Cmd {
		case proto.Register:
			err := registerHandle(msg.Data)
			if err != nil {
				fmt.Println("registerHandle err")
			}

		case proto.Login:
			err := loginHandle(msg.Data)
			if err != nil {
				fmt.Println("loginHandle error.")
			}
		case proto.Msg:
			err := msgHandle(msg.Data)
			if err != nil {
				fmt.Println("msgHandle error.")
			}
		default:
			fmt.Println("handle default")
		}

	return
}

func readPack(conn net.Conn)(msgPack proto.Message,err error){
	var buf =make([]byte,8)
	n, err := conn.Read(buf[0:4])
	if err !=nil {
		if err == io.EOF{
			fmt.Println("readPackData:对方已经断开连接！")
			return
		}
		fmt.Println("readPackage error :",err)
		return
	}
	if n != 4 {
		fmt.Println("read header failed")
		return
	}
	packLen := binary.BigEndian.Uint32(buf[0:4])
	var buff    [8192]byte
	n, err = conn.Read(buff[0:packLen])
	if n != int(packLen) {
		fmt.Println("read body failed")
		return
	}

	err = json.Unmarshal(buff[:packLen],&msgPack)
	if err != nil{
		fmt.Println("readPackData json Unmasrshal1 error.")
	}
	return
}



func msgHandle(data string)(err error){

	var msg proto.MessageRsp
	err = json.Unmarshal([]byte(data),&msg)
	if err != nil{
		fmt.Println("json Unmasrshal error.")
	}

	fmt.Println("新消息：",msg.Msg)

	return
}


func loginHandle(data string)(err error){

	var msg proto.LoginRsp
	err = json.Unmarshal([]byte(data),&msg)
	if err != nil{
		fmt.Println("json Unmasrshal2 error.")
	}

	switch msg.Flag {
	case 0:
		fmt.Println("登录成功")
		alreadyLogin = true
	case 1:
		fmt.Println("密码错误")
		alreadyLogin = false
	case 2:
		fmt.Println("没有此用户")
		alreadyLogin = false
	}

	return
}


func registerHandle(data string)(err error){

	var msg proto.RegisterRsp
	err = json.Unmarshal([]byte(data),&msg)
	if err != nil{
		fmt.Println("json Unmasrshal error.")
	}

	switch msg.Flag {
	case 0:
		fmt.Println("注册成功")

	case 1:
		fmt.Println("注册失败")
	case 2:
		fmt.Println("用户名已经被使用了")
	}

	return
}