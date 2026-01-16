package service

import (
	"poker/model"
   "fmt"
)

func(t*TexasJudge) DealCount(){
   c:=model.CountId{}
   for i:=0;i<len(t.Deal);i++{
      cards:=t.Deal[i]//deal是一个二维切片
       card:=[]int{}
       color:=make([]int,5)
       number:=make([]int,15)
   for _,v:=range cards{
        card=append(card,v)
       c:=v&3
       color[c]++
       n:=v>>2
       number[n]++
   }
   c.SetAll(color,number,card)
}
  all:=c.GetAll()
  fmt.Println(all)
  for i:=0;i<len(all);i++{
    fmt.Println(all[i].Number)
  }
}