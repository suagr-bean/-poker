package Service

import (
	"poker/model"
	"poker/model/httpModel"
	
	"poker/Utils"
)

//业务入口 蒙特卡洛计算德州
func CalRate(cal*httpModel.CalData)model.CalResult{
    
	for k:=range model.ScoreMap{
		model.ScoreMap[k]=0
	}
   data:= ToDealData(cal)//处理好的对象
   win:=float32(0.0)
   for i:=0;i<data.Frequency;i++{
    result:=SingleGame(data)//每一次的循环
     if result.Lose==false{//赢了
		win+=result.Ev
	 }
	 if result.Lose==true{//输了 统计
       for _,v:=range result.LoseScore{
		CountScore(v)
	   }
	 }
   }
   //计算输哪些牌
     cardtype:=LoseScores(data.Frequency)
	 winresult:=win/float32(data.Frequency)*100
	 ev:=float32(0.0)
     if cal.Table.DeadMoney>=0&&cal.Table.CallMoney>=0{ 
       ev=ActionEv(cal.Table.DeadMoney,cal.Table.CallMoney,winresult)
	 }
	final:=model.CalResult{
		Win: winresult,
        LoseInfo: cardtype,
		Ev :ev,
	}
   return final
}
//把数据简单处理了 该map的map
func ToDealData(cal*httpModel.CalData)*model.InitData{
	hand:=cal.Cards.Hand
	public:=cal.Cards.PublicCards
	Handmap:=Utils.DealMap(hand)
	PublicMap:=Utils.DealMap(public)
      data:=&model.InitData{
	  Person: cal.Table.Person,
	  Id: cal.Cards.Position,
      Hand: Handmap,
      PublicCard: PublicMap,
	  Frequency: 10000,
	  Action:cal.Table.PlayerAction,
	}
	return data
}

func LoseScores(frequency int)model.CardType{
    
	cardtype:=model.CardType{
			RoyalFlush: float32(model.ScoreMap[10])/float32(frequency)*100,
			StraightFlush: float32(model.ScoreMap[9])/float32(frequency)*100,
			FourOfKind: float32(model.ScoreMap[8])/float32(frequency)*100,
			FullHouse: float32(model.ScoreMap[7])/float32(frequency)*100,
			Flush: float32(model.ScoreMap[6])/float32(frequency)*100,
			Stright: float32(model.ScoreMap[5])/float32(frequency)*100,
			ThreeOfKind: float32(model.ScoreMap[4])/float32(frequency)*100,
			TwoPairs: float32(model.ScoreMap[3])/float32(frequency)*100,
			OnePair: float32(model.ScoreMap[2])/float32(frequency)*100,
			HighCard: float32(model.ScoreMap[1])/float32(frequency)*100,
	}
      return cardtype
}


func CountScore(score int){
    s:=GetScore(score) 
   model.ScoreMap[s]++
   
}
func GetScore(score int)int{
	   n:=score/10000 //最小的分也是 1万
	   return n
}
func ActionEv(deadMoney float32,callMoney float32, win float32)float32{
	lose:=(float32(100.0)-win)/100
	
	ev:=(win/100)*(deadMoney)-lose*callMoney
	return ev
}