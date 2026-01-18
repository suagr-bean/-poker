package service

import (
	"poker/model"
	
)
 
func GameMain(data*model.Begin,result func(s float32)){
	//1.造对象确认什么游戏
	texas:=GameFactory("texas")
    //添加玩家池和设置自己的牌
	player:=model.NewGameData(data.Person)
	player.Add(data.Id,data.Hand)
	 id:=data.Id//知道自己哪个位置
	 
     
    //添加公牌
	board:=model.NewBoard(data.PublicCard)
     dealer:=texas.MakeDealer()
	 
	 all:=dealer.Init(board,player)//接收一个要处理的二维切片
	judge:=texas.MakeJudge()
	
    judge.CallBack(result)
	judge.InitCard(id,all)
}