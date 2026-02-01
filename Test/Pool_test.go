package service

import (
	
	"poker/model"
	
	"testing"
)
func TestPo(t *testing.T){
	p:=model.NewPlayerInfo(3)//3个玩家
	cards:=[]int{}
	b:=model.NewBoard(cards)
    p.SkipId=0
	p.AddHand(0,[]int{58,59})
	p.AddRange(1,50)
	p.AddRange(2,25)
	
   obj:=model.NewPool(p,b)
   hand1:=[]int{18,19}
   obj.Add(hand1)
   t.Log(obj.GetPool())
}