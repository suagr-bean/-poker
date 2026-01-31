package service

import (
	"poker/model"
 "poker/service/chain"
  
)
   
type TexasJudge struct {
   Id int
  board *model.Board
  Count *model.CountId
  DealScore model.Score
  Win float32
}

func NewJudge()*TexasJudge{
	return &TexasJudge{}
}

func(t*TexasJudge) InitCard(board*model.Board)model.Result{
  
  t.board=board
  
  t.DealCount()//计数桶
  t.DealMask()//预处理 掩码
  result:=t.Start()
  return result
}
 
func (t*TexasJudge)Start()model.Result{
    var scores []int
    
   for i:=0;i<len(t.Count.Id);i++{
     chain:=service.NewChain()
     d:=t.Count.Id[i]
    
     _,score:=chain.Start(d)
    
     scores=append(scores,score)
   }
    win,data:= t.Compare(scores)
   result:= model.Result{
    Win: win,
    Lose: data,
  }
  return result
}
func (t*TexasJudge)Compare(scores[]int)(float32,[]int){ 
   my:=scores[t.Id]
   draw:=0
   maxscore:=scores[0]
   recond:=[]int{}
   for i,v:=range scores{
      if i==t.Id{
        continue
      }
      if v>maxscore{
        maxscore=v
      }
   }
   if my<maxscore{
     recond=append(recond,maxscore)
    return float32(0),recond
   }
   for _,j:=range scores{
     
     if j==maxscore{
      draw++
     }
   }
   if draw!=0{
   return float32(1)/float32(draw),recond
   }
   return float32(1),recond
}
