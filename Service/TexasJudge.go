package Service

import (
	"poker/Service/chain"
	"poker/model"
)


type TexasJudge struct{
   
}
func NewTexasJudge()*TexasJudge{
	return&TexasJudge{}
}
func (T*TexasJudge)Init(p*model.PlayerInfo,b*model.Board)model.Result{
	T.DealSingleScore(p,b)
   result:= T.Compare(p)
   return result

}
//给每个玩家打分
func (T*TexasJudge)DealSingleScore(p*model.PlayerInfo,b*model.Board){
      publicCards:=b.GetCards()
	  cards:=make([]int,7)
	  copy(cards[2:7],publicCards)
	  for _,player:=range p.Players{
		copy(cards[0:2],player.Hand)
     count:=DealCount(cards) 
	 chains:=chain.NewChain()//造责任链打分	
	err,score:=chains.Start(count)//打分
    if err==false{//打分错误
		return 
	}
    player.Score=score
}  
}
//计数
func DealCount(cards[]int)*model.Count{
	count:=model.NewCount(cards)
	countcolor:=make([]int,4)
	countNumber:=make([]int,15)
	for _,v:=range cards{
		color:=v&3
		count.Color=append(count.Color, color)
		number:=v>>2
		count.Number=append(count.Number, number)
		countcolor[color]++
		countNumber[number]++
	}
	count.SetCount(countcolor,countNumber)
	 mask,maskcolor:=DealMask(count.Color,count.Number)

	 count.MaskNumber=mask
	 count.MaskColor=maskcolor
	
	return count
}
func (T*TexasJudge)Compare(p*model.PlayerInfo)model.Result{
	 draw:=0
	 win:=0
	myid:=p.SkipId
	result:=model.Result{}
	maxscore:=p.Players[myid].Score
	for i,player:=range p.Players{
	  if player.Score>maxscore{
		result.LoseIndex=append(result.LoseIndex,i)
		result.LoseScore=append(result.LoseScore,player.Score)
		result.Ev=0.0
		 return result
	  } else if player.Score==maxscore{
        draw++
	  }else {
		continue
	  }
	}
	if draw>0&&win==0{
	  result.WinScore=0
	  result.Ev=float32(1)/float32(draw)
	}
     result.WinScore=float32(maxscore)
	 result.Ev=1.0
	return result
}