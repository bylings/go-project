package controller

import (
	"errors"
	"fmt"
	"reflect"
	"simple_system/util"
	"strconv"
	"strings"
)

var (
	autoController *AutoController
	next           string
	view           string
)

func init() {
	autoController = &AutoController{}
}

func Run() {
	next = "index::Welcome"

	for {
		flag := util.CReturn(util.Cfunc(dispatch))
		if flag {
			break
		}
	}
	fmt.Println("结束")
}

func dispatch() (bool, error) {
	// 1、根据指定的控制器和方法执行
	args := strings.Split(next, "::")
	controller, ok := controllers[args[0]]
	if ok != true {
		return false, errors.New("获取不到控制器" + args[0])
	}

	// 反射执行方法
	// 	1、传递执行的控制器对象	controller
	// 	2、根据方法名，执行方法
	cr := reflect.ValueOf(controller) // 返回结构体 reflect.value

	// 得到控制器中的方法MethodByName()，再调用Call执行方法
	cr.MethodByName(args[1]).Call([]reflect.Value{})

	// 2、获取下一步执行的操作,结构如下
	//opers = [][3]string{
	//	0: {0: "auto", 1: "login", 2: "登陆系统"},
	//	1: {0: "auto", 1: "register", 2: "注册用户"},
	//}

	opers, ok := views[view]
	if ok != true {
		return false, errors.New("获取不到视图" + view)
	}

	// 3、数据处理
	//methods{
	//	0: "auto::login",
	//	1: "auto::register",
	//}
	//desc{
	//	0: "登陆系统",
	//	1: "注册用户",
	//}
	methods, desc := toModelFormate(opers)

	util.Coper(desc)
	// 4、用户的界面展示
	for {
		input := util.CRead()
		if input == "x" {
			return true, nil
		}
		flag, err := strconv.Atoi(input)

		//fmt.Println("获取到的命令：", flag)

		if err == nil && flag < len(methods) {
			next = methods[flag]
			break
		}

		fmt.Println("信息输入有误，请重新输入")
	}
	return false, nil
}
