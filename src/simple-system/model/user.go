package model

import (
	"strconv"
)

type User struct {
	username string
	password string
	age      int
	sex      string
}

var userDatas map[string]Model

func NewUser() *User {
	return &User{}
}

func (u *User) SetUsername(username string) {
	u.username = username
}
func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) SetAge(age int) {
	u.age = age
}

func (u *User) SetSex(sex string) {
	u.sex = sex
}

func (u *User) GetUsername() string {
	return u.username
}
func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetAge() int {
	return u.age
}

func (u *User) GetSex() string {
	return u.sex
}

func GetUser(username string) *User {
	if userDatas == nil {
		return nil
	}
	//fmt.Println(userDatas)  使用断言
	flag := false
	for userKey, _ := range userDatas {
		//fmt.Println("k：", userKey, "v：", v)
		if username == userKey {
			flag = true
			break
		}
	}
	if flag == false {
		return nil
	}

	return userDatas[username].(*User)
}

// 格式化输出数据信息
func (u *User) ToString() string {
	return u.username + "," + u.password + "," + strconv.Itoa(u.age) + "," + u.sex
}

func (u *User) save() bool {
	userDatas[u.username] = u
	return fwrite("user", userDatas)
}

func (u *User) All() []*User {
	var users []*User = make([]*User, 0)
	for _, user := range userDatas {
		users = append(users, user.(*User))
	}
	return users
}
