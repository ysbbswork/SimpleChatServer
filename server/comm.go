package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"SimpleChatServer/proto"
)

func (p *Client) readPackage() (msg proto.Message, err error) {
	//这里可以测试一下，两次读IO的速度和，一起读完再检测的速度
	//是先进行网络IO读取包头，判断包头是否正确，再进行包体IO，还是一次性接收完所有，再判断包头包体是否正确。
	n, err := p.conn.Read(p.buf[0:4])
	if err !=nil {
		if err == io.EOF{
			fmt.Println("对方已经断开连接！")
			return
		}
		fmt.Println("readPackage error :",err)
		return
	}
	if n != 4 {
		err = errors.New("read header failed")
		return
	}


	var packLen uint32
	packLen = binary.BigEndian.Uint32(p.buf[0:4])

	fmt.Printf("receive len:%d\n", packLen)
	n, err = p.conn.Read(p.buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}

	fmt.Printf("receive data:%s\n", string(p.buf[0:packLen]))
	err = json.Unmarshal(p.buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
	}
	return
}

func (p *Client) writePackage(data []byte) (err error) {

	packLen := uint32(len(data))

	binary.BigEndian.PutUint32(p.buf[0:4], packLen)
	n, err := p.conn.Write(p.buf[0:4]) //这里需要改，不需要写两次网络IO
	if err != nil {
		fmt.Println("write data  failed")
		return
	}

	n, err = p.conn.Write(data)
	if err != nil {
		fmt.Println("write data  failed")
		return
	}

	if n != int(packLen) {
		fmt.Println("write data  not finished")
		err = errors.New("write data not fninshed")
		return
	}

	return
}
