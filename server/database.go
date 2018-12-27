package main

import (
	"fmt"
	"sync"
)

var DATABASE DataBase

type DataBase struct {
	data map[string]string
	mu sync.Mutex
}

func DataBaseInit(){
	DATABASE.data = make(map[string]string)

}
func (p *DataBase) addUser(name, password string) {
	p.mu.Lock()
	p.data[name] = password
	p.mu.Unlock()
}

func (p *DataBase) delUser(name string) {
	p.mu.Lock()
	delete(p.data,name)
	p.mu.Unlock()
}

func (p *DataBase) checkUserandPassword(name, password string) bool{
	res,ok :=DATABASE.data[name]
	if ok&& res == password {
		return true
	}
	return false
}

func (p *DataBase) checkUser(name string) bool{
	_,ok :=DATABASE.data[name]
	fmt.Println(ok)
	return ok
}