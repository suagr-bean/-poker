package controller

import (
	"fmt"
	"poker/Service"
	"poker/model"
	"poker/model/httpModel"
)
type DataResult struct{
	Win float32 
	Ev  float32 

}
//控制数据流向
func PokerHandler(cal *httpModel.CalData) model.CalResult {
	
	result:= Service.CalRate(cal)
	win:=result.Win
	fmt.Println("结束计算",win)
    
	return result
}
