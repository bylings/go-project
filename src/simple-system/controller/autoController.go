package controller

import (
	"fmt"
	"simple_system/model"
	"simple_system/util"
)

type AutoController struct {
}

func (c *AutoController) Login() {
	view = "login_view"
	fmt.Print("输入你的用户名 : ")
	username := util.CRead()
	fmt.Print("输入你的密码 : ")
	password := util.CRead()
	
	user := model.GetUser(username)
	if user == nil {
		fmt.Println("查询不到用户", username)
		return
	}
	if user.GetPassword() == password {
		fmt.Println("登入成功")
		view = "index_view"
		return
	} else {
		fmt.Println("密码错误")
		return
	}
}

func (c *AutoController) Register() {
	fmt.Println("注册用户")
}
