package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	type location struct {
		// 字段名大写是公共变量，小写是私有变量
		Lat  float64 `json:"lat"` // json标签 ``
		Long float64 `json:"long"`
	}
	// 类型可以被复用
	var spirit location
	spirit.Lat = -14.243
	spirit.Long = 175.5234

	fmt.Println(spirit)

	opportunity := location{Lat: -1.342, Long: 2.5646}
	fmt.Println(opportunity)

	// 初始化
	insight := location{Lat: -1.342}
	fmt.Println(insight)

	spirit = location{-1.345, 3.34}

	// 打印结构体
	fmt.Printf("%v\n", spirit)
	fmt.Printf("%+v\n", insight)
	fmt.Println("---------------------------")

	bytes, err := json.Marshal(insight)
	exitOnError(err)
	fmt.Println(string(bytes))

}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
