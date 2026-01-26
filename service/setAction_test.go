package service

import (
	"poker/model"
	"testing"
	"fmt"
)

func TestAction(t *testing.T){
	
	 player:=model.NewGameData(6)
    action:=[]int{1,3,5}
    SetPlayer(player,action)
	fmt.Println(player.Players[1].Action)
}