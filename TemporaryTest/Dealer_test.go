package service

import (
	"poker/Service"
	"testing"
)
func TestDeal(t *testing.T){
   game:=Service.GameFactory("texas")
   dealer:=game.MakeDealer()
   
}