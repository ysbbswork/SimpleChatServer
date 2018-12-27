package main

import "fmt"


type UserMgr struct {
	onlineUsers map[string]*Client
}

var (
	userManger *UserMgr
)

func init() {
	userManger = &UserMgr{
		onlineUsers: make(map[string]*Client, 1024),
	}
}

func (p *UserMgr) AddClient(userId string, client *Client) {
	p.onlineUsers[userId] = client
}

func (p *UserMgr) GetClient(userId string) (client *Client, err error) {
	client, ok := p.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("user %d not exists", userId)
		return
	}

	return
}

func (p *UserMgr) GetAllUsers() map[string]*Client {
	return p.onlineUsers
}

func (p *UserMgr) DelClient(userId string) {
	delete(p.onlineUsers, userId)
}

