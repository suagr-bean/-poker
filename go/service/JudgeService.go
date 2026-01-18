package service

import (
	"poker/model"
	service "poker/service/process"
)
   
type TexasJudge struct {
   Id int
  Deal [][]int
  Count *model.CountId
  DealScore model.Score
  Win float32
  callback(func(score float32))
}

func NewJudge()*TexasJudge{
	return &TexasJudge{}
}
func (t*TexasJudge)CallBack(f func(score float32)){
    t.callback=f //回调函数
}
func(t*TexasJudge) InitCard(id int,card[][]int){
  t.Id=id
  t.Deal=card
  t.DealCount()//计数桶
  t.DealMask()//预处理 掩码
  t.Start()
  
  t.callback(t.Win)
}
 
func (t*TexasJudge)Start(){
    var scores []int
   for i:=0;i<len(t.Count.Id);i++{
     chain:=service.NewChain()
     d:=t.Count.Id[i]
     _,score:=chain.Process(d)
     scores=append(scores,score)
   }
  t.Win+= t.Compare(scores)
  
}
func (t*TexasJudge)Compare(scores[]int)float32{
    
   my:=scores[t.Id]
   
   draw:=0
   maxscore:=scores[0]
   for _,v:=range scores{
      if v>maxscore{
        maxscore=v
      }
   }
   if my<maxscore{
    return float32(0)
   }
   for _,j:=range scores{
     if j==maxscore{
      draw++
     }
   }
   if draw!=0{
   return float32(1)/float32(draw)
   }
   return float32(1)
}
