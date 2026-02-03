package Service

import "poker/model"
//设置自己的手牌加上 设置对手的range范围
func PlayerManage(data*model.InitData)*model.PlayerInfo{
  
  player:=model.NewPlayerInfo(data.Person)
  player.AddHand(data.Id,data.Hand)
  for i,v:=range data.Action{//把玩家范围设置好
	 player.AddRange(i+1,v)
  }
  return player
}