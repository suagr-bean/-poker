package service

import (
	
	"poker/Service"
	"poker/model"
	"testing"
)
func TestDealerPlayer(t*testing.T){
   texas:=Service.GameFactory("texas")
    dealer:=texas.MakeDealer()
	p:=model.NewPlayerInfo(4)
    p.SkipId=2
	p.AddHand(2,[]int{10,15})
    p.AddRange(3,25)
	b:=model.NewBoard(nil)
	pool:=dealer.DealHand(p,b)
	t.Log(pool.GetPool())
	
}