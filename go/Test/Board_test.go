package service

import (
	"poker/model"
	"testing"
)
//测试公牌对象
func TestBoard(t *testing.T){
    board:=[]int{30,50,40,45,30}
	boards:=model.NewBoard(board)
	cards:=boards.GetCards()
	t.Log("cards",cards)
}