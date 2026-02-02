package main

import (
	"fmt"
	"poker/controller"
	"poker/Utils"
)

func main() {
	fmt.Println("start")
	//加载全局起始表
	Utils.RangeFile()
	controller.Route()
}