package Service

import (

	"math/rand/v2"
      "poker/Utils"
	"poker/model"
	
)

	//发牌员内部补玩家牌 发牌 
func (t*TexasDealer)DealHand(p*model.PlayerInfo,b*model.Board)*model.Pool{
	 poolobj:=model.NewPool(p,b)
    t.PlayerCards(p,poolobj)
     pool:=poolobj.GetPool()
	 t.DeleteKnowDeck(pool)
	 t.Shuffle()
     t.ProcessBoard(b)
   	return poolobj
}
//删掉已知牌
func (t*TexasDealer)DeleteKnowDeck(pool []int){
	
   for _,j:=range pool{
	for i,x:=range t.Deck.Cards{
		if j==x{
		  length:=len(t.Deck.Cards)-1
          t.Deck.Cards[i]=t.Deck.Cards[length]
		  t.Deck.Cards=t.Deck.Cards[:length]
		  break
		}
	}
   } 
}
//洗牌
func (t*TexasDealer)Shuffle(){
	rand.Shuffle(len(t.Deck.Cards),func(i,j int){
	   t.Deck.Cards[i],t.Deck.Cards[j] =t.Deck.Cards[j],t.Deck.Cards[i]
	})
}
//给玩家发手牌 根据玩家范围
func (t*TexasDealer)PlayerCards(p*model.PlayerInfo ,pool*model.Pool){
	
   for i:=0;i<p.Person;i++{
     if i==p.SkipId{
		continue
	 }
	 cards:=pool.GetPool()//要删除的牌
	 if len(p.Players[i].Hand)==0{
	   hand,error:= Utils.Rand(p.Players[i].PlayerRange,cards)//根据玩家range发牌
	   if error!=nil{
		//处理错误
	   }
	   pool.Add(hand)
	   p.Players[i].Hand=append(p.Players[i].Hand,hand...)
	 }
   }
}
func (t*TexasDealer)ProcessBoard(b*model.Board){
	if len(b.GetCards())!=5{
		need:=5-len(b.GetCards())
		cards:=t.DealingCards(need)
		b.AddCards(cards)
	}

}
func (t*TexasDealer)DealingCards(need int)[]int{
	cards:=t.Deck.Cards[:need]
	t.Deck.Cards=t.Deck.Cards[need:]
	return cards
}