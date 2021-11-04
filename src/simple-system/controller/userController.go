package controller

import "fmt"

type UserController struct {
}

func (c *UserController) List() {
	fmt.Println("展示用户信息")
}
