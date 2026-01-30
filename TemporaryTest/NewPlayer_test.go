package service

import (
	"poker/model"
	
	"testing"
)

//测试添加玩家手牌和range是否正常
func TestNewPlayer(t *testing.T){
	
	 player:=model.NewPlayerInfo(6)
	 hand:=[]int{58,59}
	 
	 player.AddHand(0,hand)
     t.Log("\n","data",player.Players[0])
}