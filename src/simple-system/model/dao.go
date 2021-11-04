package model

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// 统一的模型接口
type Model interface {
	ToString() string // 格式化输出数据信息
}

var (
	path   string                 = "" // 数据路径
	suffix string                 = ".sql"
	models map[string]interface{} // 记录标识，user=>结构体
)

// 初始化
func init() {
	sysType := runtime.GOOS
	//fmt.Println("当前系统是：", sysType)
	switch sysType {
	case "windows":
		path = "E:\\GoWorks\\src\\simple-system\\data\\"
	case "darwin":
		path = "/Users/mzj/Documents/code/server/go/go-project/src/simple-system/data/"
	}

	models = make(map[string]interface{})
	//models["user"] = &User{}
	models["user"] = NewUser

	userDatas = make(map[string]Model, 0)
	rfdata("user", "username", userDatas)
}

// 读文件 ->通过配置设置
// name 数据库表名称，user,admin
// primary 查询主键
// models 存放数据
func rfdata(name, primary string, datas map[string]Model) error {
	// 1、读取数据库文件=》读取哪个文件
	f, ferr := os.Open(path + name + suffix)
	if ferr != nil {
		fmt.Println("文件读取异常", ferr)
		return errors.New("文件查询失败 error")
	}

	// 延迟执行释放文件句柄
	defer f.Close()

	// 创建文件缓冲区，重复读取不消耗io
	buf := bufio.NewReader(f)

	// 判断是否为字符串
	field := make([]string, 0)
	// 2、遍历每一行数据  字段根据,分割；数据通过\n分割
	for {
		row, rerr := buf.ReadBytes('\n')

		//fmt.Println("值：---- ", string(row), rerr) // 先打印，再判断err2 (如果文件末尾没有'\n'，那么需要先打印再判断err2)
		if rerr != nil {
			// 判断文件是否读取结束
			if rerr == io.EOF {
				//fmt.Println("--------- 结束", rerr)
				break
			}
			fmt.Println("抛出缓存读取文件异常", rerr)
		}
		// 读取到的数据换行，分割数据
		data := strings.Split(strings.TrimSuffix(string(row), "\n"), ",")
		//fmt.Println("读取到的文件信息\n", data)

		// 	2.1 是否为字段
		// 根据数据判断操作：是记录字符串还是设置数据
		if len(field) == 0 {
			field = data
			for k, v := range data {
				field[k] = strings.TrimSpace(strings.TrimSuffix(v, "\n"))
			}
		} else {
			//  2.2 存储数据到models
			// 		2.2.1 根据name得到模型
			// 		2.2.2 利用反射-》对模型赋值
			// 		2.2.3 再把模型存储在datas
			/**
			datas := {
				"primary":model(data),
				"primary1":model(data),
				"primary2":model(data),
			}
			*/
			toModel(name, primary, datas, data, field)
		}
		// 读取到存储的字段
		//fmt.Println("读取到字段\n", field)
	}
	return nil

}

//  2.2 存储数据到models
// 		2.2.1 根据name得到模型
// 		2.2.2 利用反射-》对模型赋值
// 		2.2.3 再把模型存储在datas
func toModel(name, primary string, datas map[string]Model, data, field []string) error {
	// 2.2.1 根据name得到模型
	if models[name] == nil {
		return errors.New("不存在的模型：" + name)
	}
	// 2.2.2 利用反射-》对模型赋值   -> 如果是采用构造函数的方式则需要利用反射获取
	modelV := reflect.ValueOf(models[name]).Call([]reflect.Value{})[0]
	//fmt.Printf("modelV type：%T \n", modelV)

	//fmt.Println("data 数据：\n", data) // admins,123456,18,男

	var primaryValue string // 记录当前主键的数值
	for k, v := range data {
		// 判断是否为主键字段的值
		if field[k] == primary {
			primaryValue = v
			//fmt.Println("查询到的主键值：", primaryValue)
		}
		// 得到model中对应字段的set方法，使用反射 MethodByName,例如  SetUsername
		fset := modelV.MethodByName("Set" + strings.Title(field[k]))
		//fmt.Printf("fset type %T \n", fset)
		//fmt.Println("fset：", fset)

		// 调用方法,并且传参
		// 参数值 可能是int（意思是要求传参值的类型，传入的参数要保持一致）  ----> 从模型中获取
		// 1、根据模型获取属性字段
		// 2、
		fset.Call([]reflect.Value{
			reflect.ValueOf(ToTypeValue(modelV, field[k], v)),
		})
		//fmt.Println("field[k] :", field[k])

		//mtype := modelV.Elem().FieldByName(field[k]).Type().Name()
		// reflect.Value.Elem() 得到类型
		// reflect.Zero()   &User{}  =>  User{}  根据指针得到原有对象

		//modelZ := reflect.Zero(modelV.Type().Elem()).FieldByName(field[k]).Type().Name()
		//
		//fmt.Println("mtype", mtype)

	}

	//fmt.Println("model ", modelV)  ,断言
	datas[primaryValue] = modelV.Interface().(Model)
	return nil
}

// 字符类型
func ToTypeValue(modelV reflect.Value, field, value string) interface{} {

	mtype := modelV.Elem().FieldByName(field).Type().Name()
	switch mtype {
	case "int":
		b, _ := strconv.Atoi(value)
		return b
	}
	return string(value)
}

// 写文件
func fwrite(name string, models map[string]Model) bool {
	// 获取到转化的数据内容
	content := getModelsToString(models)
	outfile, outErr := os.OpenFile(path+name+suffix, os.O_WRONLY|os.O_CREATE, 0666)
	if outErr != nil {
		fmt.Println("文件找不到")
		return false
	}
	defer outfile.Close()

	outbufwrite := bufio.NewWriter(outfile)
	_, _ = outbufwrite.WriteString(content + "\n")
	_ = outbufwrite.Flush()
	return true
}

// 把模型数据源转化为字符串
func getModelsToString(models map[string]Model) string {
	// 记录字段内容
	var fields string
	// 循环处理数据
	var content string
	for _, model := range models {
		if fields == "" {
			// 利用反射获取字段内容
			// rmodel := reflect.TypeOf(model)
			rmodel := reflect.ValueOf(model)
			modelZ := reflect.Zero(rmodel.Type().Elem())
			for i := 0; i < modelZ.NumField(); i++ {
				fields = fields + modelZ.Type().Field(i).Name + ","
			}
			fields = strings.TrimSuffix(fields, ",")
		}
		// 记录数据内容
		content = content + model.ToString() + "\n"
	}
	// 最终内容
	return fields + "\n" + strings.TrimSuffix(content, "\n")
}
