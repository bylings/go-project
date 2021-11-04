package simple_system_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type myFun func()

func (f myFun) call() {
	fmt.Println("================ >>> start ================ >>>")
	f() //
	fmt.Println("================ >>> end   ================ >>>")
}

/*
route :={
	"index":{
		"oper":[],
		"index":[],
	}
}
*/

func TestFun(t *testing.T) {
	newSumcoo := myFun(dispatch)
	newSumcoo.call()
}

func dispatch() {
	fmt.Println("dddd")
}

type ConfigData struct {
	DataPath string `json:"data_path"`
	BasePath string `json:"base_path"`
}

func TestJson(t *testing.T) {
	// type => json	结构体转json
	c := ConfigData{"data_path", "base_path"}
	fmt.Println(c)
	cj, _ := json.Marshal(c)

	cj1, _ := json.MarshalIndent(c, "", " ")
	fmt.Println(string(cj))
	fmt.Println(string(cj1))

	js := `{
		"data_path":"这是数据路径",
		"base_path":"这是项目路径"
	}`

	cb := []byte(js)
	var c1 ConfigData
	json.Unmarshal(cb, &c1)
	fmt.Println(c1)

	cj2, _ := json.Marshal(c1)
	cj3, _ :=  json.MarshalIndent(c1, "", " ")
	fmt.Println(string(cj2))
	fmt.Println(string(cj3))
	//fmt.Printf("%T \n", c1)

}
