package controller

import "fmt"

type IndexController struct {
}

func (c *IndexController) Welcome() {
	fmt.Println("欢迎来到XXX系统")
	fmt.Println("你要执行的操作")
	// 设置下一个页面的流程
	view = "login_view"
}

func (c *IndexController) Index() {
	fmt.Println("进入首页")
	view = "login_view"
}

func (c *IndexController) List() {
	fmt.Println("信息展示页")
	view = "login_view"
}
