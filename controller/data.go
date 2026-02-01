package controller

import (
	"poker/model/httpModel"
	"poker/Service"
	"fmt"
)

func PokerHandler(cal *httpModel.CalData) float32 {
	fmt.Println("进入计算")
	result:= Service.CalRate(cal)
	win:=result.Ev
	fmt.Println("结束计算",win)
	return win
}
