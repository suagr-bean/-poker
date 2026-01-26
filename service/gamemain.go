package service

import (
	"poker/model"
	
)
 
func GameMain(action[]int,data*model.Begin,result func(s float32),record func(int)){
	//1.造对象确认什么游戏
	texas:=GameFactory("texas")
    //添加玩家池和设置自己的牌
	player:=model.NewGameData(data.Person)
	player.Add(data.Id,data.Hand,1)//1先代表全下
	 id:=data.Id//知道自己哪个位置
	SetPlayer(player,action)
	
    //添加公牌
	board:=model.NewBoard(data.PublicCard,data.Person)
     dealer:=texas.MakeDealer()
	
	 dealer.Init(board,player)
	judge:=texas.MakeJudge()
	
    judge.CallBack(result)
	judge.Record(record)
	judge.InitCard(id,board)
}