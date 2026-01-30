package Service

import (
	"poker/model"
	
)
 //反回judge计算的结果集
func Start(data*model.Begin)model.Result{
	//1.造对象确认什么游戏
	texas:=GameFactory("texas")
    //添加玩家池
	player:=PlayerManage(data)
	
    //添加公牌
	board:=model.NewBoard(data.PublicCard)
	//造发牌员
     dealer:=texas.MakeDealer()
	 
	 dealer.Init(player,board)

	judge:=texas.MakeJudge()
	
    
     result:=judge.InitCard(id,board)
	 return result
}