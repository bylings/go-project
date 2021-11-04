package model

import (
	"fmt"
	"simple_system/util"
)

var (
	filePath   string
	fileSuffix string = ".sql"
	newModels  map[string]interface{}
	config     Config
)

type Config struct {
}

func init() {
	fmt.Println("model --- >>>")
	filePath = util.GetConfig().GetDataPath()
	fmt.Println("path：", path)
	initData()
	initNewModel()
}

// 用于记录需要创建模型的方法¬
func initNewModel() {
	newModels = make(map[string]interface{})
	newModels["user"] = NewUser
	userDatas = make(map[string]Model, 0)
	_ = rfdata("user", "username", userDatas)
}

// 初始化数据
func initData() {
	// 校验文件是否存在
	flag, _ := util.PathExist(path + "user.sql")
	if !flag {
		data := "username,password,age,sex\nroot,123456,18,男\nadminss,123456,18,女\nadminsss,123456,18,男\n"
		b, err := util.WriteFile(path, "user.sql", data)
		fmt.Println("创建结果：", b, "error：", err)
	}

}
