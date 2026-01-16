package service

import (
	"poker/model"
)
func GameMain(){
	//1.造对象确认什么游戏
	texas:=GameFactory("texas")
    //添加玩家池和设置自己的牌
	player:=model.NewGameData(6)
	player.Add(0,[]int{15,30})
	
    //添加公牌
	board:=model.NewBoard([]int{8,9,10})
     dealer:=texas.MakeDealer()
	 all:=dealer.Init(board,player)//接收一个要处理的二维切片
	judge:=texas.MakeJudge()
	judge.InitCard(all)
}