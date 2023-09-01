package core

import (
	"github.com/zurvan-lab/TimeTraceDB/utils"
)

type User struct {
	Name, Token string
	Cmds        []string
}

type Users map[int]User

func CreateUsers() *Users {
	return &Users{}
}

func (u *Users) NewUser(name, Token string, cmds []string) {
	for {
		id, _ := utils.GenerateRandomNumber(1, 99999999)
		if _, ok := (*u)[id]; !ok {
			(*u)[id] = User{Name: name, Cmds: cmds, Token: Token}
			break
		}
	}
}
