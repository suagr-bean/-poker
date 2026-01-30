package Service

import"poker/model"
func(t*TexasDealer) Init(b*model.PlayerInfo,data*model.Board){
	deck:=NewDeck()
	t.Board=b
	t.Data=data
	t.Cards=deck
	
	t.Known() //删掉已知牌组
	t.ShuffleCard()
	t.Process()//补玩家牌
	t.Deal()

}