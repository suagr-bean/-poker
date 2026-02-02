package main

import (
	"fmt"
	"poker/controller"
	"poker/Utils"
)

func main() {
	fmt.Println("start")
	//加载全局起始表
	error:=Utils.RangeFile()
	if error !=nil{//处理很重要的错误 说明文件被删或者路径变了 
		fmt.Println("文件没读到")
	}
	controller.Route()
}