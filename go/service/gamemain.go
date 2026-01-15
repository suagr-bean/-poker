package service

import (
	"poker/model"
)
func GameMain(){
	//1.造对象确认什么游戏
	texas:=GameFactory("texas")
	//造发牌员 新牌组
   
    
	//前面deck ，后面发牌员
    //造玩家池 加玩家手牌
	player:=model.NewGameData(6)
	player.Add(0,[]int{15,30})
	
    //公牌
	board:=model.NewBoard([]int{8,9,10,25,36})
     dealer:=texas.MakeDealer()
	 dealer.Init(board,player)
	
}