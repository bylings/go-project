package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 缓冲区
var inputReader *bufio.Reader

type Cfunc func() (bool, error)

func init() {
	inputReader = bufio.NewReader(os.Stdin)
}

func CReturn(cf Cfunc) bool {
	fmt.Println("============= >>> start ============= >>> ")
	flag, err := cf()
	if err != nil {
		fmt.Println("系统异常", err)
	}
	fmt.Println("============= >>> end   ============= >>> ")
	return flag
}

func CRead() string {
	// 获取控制台输入的信息
	input, _ := inputReader.ReadString('\n')
	input = strings.TrimSpace(strings.TrimSuffix(input, "\n"))
	return input
}

// 输出指令
func Coper(operate []string) {
	for k, v := range operate {
		fmt.Printf("(%d)：%s \n", k, v)
	}
	fmt.Println("退出请输 x ")
}
