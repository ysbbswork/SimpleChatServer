package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"SimpleChatServer/proto"
)

type Client struct {
	conn   net.Conn
	name string
	buf    [8192]byte
}


func (p *Client) Process() (err error) {

	for {
		var msg proto.Message
		//接收消息
		msg, err = p.readPackage()
		if err != nil {
			userManger.DelClient(p.name)

			message := fmt.Sprintf("user:%s 下线了", p.name)
			fmt.Println(message)
			_ = p.SendContext2allUser(message)
			return err
		}
		//根据具体命令字具体处理
		switch msg.Cmd {
		case proto.Register:
			rsp,rerr := p.register(msg)
			if rerr != nil {
				err =rerr
				fmt.Println("login error", err.Error())

				return err
			}
			err = p.sendRsp(proto.Register,rsp)

		case proto.Login:
			rsp,err := p.login(msg)
			if err != nil {
				fmt.Println("login error", err.Error())
				return err
			}
			err = p.sendRsp(proto.Login,rsp)

		case proto.Msg:
			err = p.SendMessage(msg)
			fmt.Println("SendMessage")
		default:
			err = errors.New("unsupport message")
			return
		}

		if err != nil {
			fmt.Println("process msg failed, err:", err)
			continue
			//return
		}
	}
}


func (p *Client)sendRsp(cmd int,rsp interface{})(err error){
	JsonRsp, jsonErr := json.Marshal(rsp)
	if jsonErr != nil {
		fmt.Println("sendRsp json marshal faild.")
		err = jsonErr
		return
	}
	var rspMsg proto.Message
	rspMsg.Cmd = cmd
	rspMsg.Data = string(JsonRsp)
	JsonRspMsg,errj := json.Marshal(rspMsg)
	if errj != nil {
		fmt.Println("sendRsp json marshal faild.")
		err =errj
		return
	}
	err = p.writePackage(JsonRspMsg)
	if err != nil {
		fmt.Println("sendRsp writePackage error.")
		return
	}
	return
}


func (p *Client)login(msg proto.Message) (rsp proto.LoginRsp,err error){
	var loginData proto.LoginReq
	err = json.Unmarshal([]byte(msg.Data), &loginData)
	if err != nil {
		fmt.Println("login unmarshal failed, err:", err)
		return
	}
	realPassword,ok :=DATABASE.data[loginData.Name]
	if ok {
		if realPassword == loginData.Password{
			rsp.Flag = proto.Success
			userManger.AddClient(loginData.Name ,p)
			p.name = loginData.Name
			message := fmt.Sprintf("user:%s 上线了", p.name)
			fmt.Println(message)
			_ = p.SendContext2allUser(message)
		}else{
			rsp.Flag = proto.Faild
		}
	}else{
		//没有这个用户
		rsp.Flag = proto.UserNotExist
	}

	return
}

func (p *Client)register(msg proto.Message) (rsp proto.RegisterRsp,err error){
	var data proto.RegisterReq
	err = json.Unmarshal([]byte(msg.Data), &data)
	if err != nil {
		fmt.Println("login unmarshal failed, err:", err)
		return
	}
	if DATABASE.checkUser(data.Name) {
		rsp.Flag = proto.UserAlreadyExist
		fmt.Println("User name all already exists.")
		return
	}
	DATABASE.addUser(data.Name, data.Password)
	fmt.Println("DATABASE now:",DATABASE)
	if !(DATABASE.checkUserandPassword(data.Name, data.Password)) {
		err = errors.New("register error,add user in database faild")
	}
	return
}


func (p *Client)SendMessage(msg proto.Message) (err error){
	var data proto.MessageReq
	err = json.Unmarshal([]byte(msg.Data), &data)
	if err != nil {
		fmt.Println("login unmarshal failed, err:", err)
		return
	}
	for _,client := range userManger.GetAllUsers(){

		var msg proto.MessageRsp
		msg.Msg=data.Msg
		err = client.sendRsp(2,&msg)
		if err != nil {
			fmt.Println("sendRsp error.",err.Error())
		}

	}
	return
}

func (p *Client)SendContext2allUser(data string) (err error){

	for _,client := range userManger.GetAllUsers(){
		var msg proto.MessageRsp
		msg.Msg=data
		err = client.sendRsp(2,&msg)
		if err != nil {
			fmt.Println("SendContext2allUser sendRsp error.",err.Error())
		}
	}
	return
}