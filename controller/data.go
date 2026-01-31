package controller

import (
	"poker/model/httpModel"
	"poker/service"
)

func PokerHandler(cal *httpModel.CalData) float32 {
	reuslt := service.CalRate(cal)
	win := reuslt.Win
	return win
}
