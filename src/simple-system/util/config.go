package util

import (
	"flag"
	"runtime"
)

var (
	instance   *uconfig
	configFile string
	conf       = flag.String("conf", "../etc/config.json", "这是命令描述")
)

type uconfig struct {
	DataPath string `json:"data_path"`
	BasePath string `json:"base_path"`
}

func init() {
	sysType := runtime.GOOS
	//fmt.Println("当前系统是：", sysType)
	switch sysType {
	case "windows":
		configFile = "E:\\GoWorks\\src\\simple-system\\etc\\config.json"
	case "darwin":
		configFile = "/Users/mzj/Documents/code/server/go/src/simple-system/etc/config.json"
	}

	flag.Parse()
	//fmt.Println(*conf)
	// 读取命令行-conf参数
	//newUConfigWithFile(configFile)
	newUConfigWithFile(*conf)
}

func newUConfigWithFile(configFileDir string) {
	if instance == nil {
		c := &uconfig{}
		// 读取文件
		_ = ReadJson(configFile, c)
		instance = c
	}
}

func GetConfig() *uconfig {
	//fmt.Println(*conf)
	return instance
}

func (c *uconfig) GetDataPath() string {
	return c.DataPath
}
