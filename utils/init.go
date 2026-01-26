package utils

import"poker/model"
func(t*TexasDealer) Init(b*model.Board,data*model.GameData){
	deck:=NewDeck()
	t.Board=b
	t.Data=data
	t.Cards=deck
	
	t.Known() //删掉已知牌组
	t.ShuffleCard()
	t.Process()//补玩家牌
	t.Deal()

}