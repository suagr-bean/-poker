package service

import (
	"poker/model"
   
)

func(t*TexasJudge) DealCount(){
   c:=model.CountId{}
   t.Count=&c
   for i:=0;i<len(t.Deal);i++{
      cards:=t.Deal[i]//deal是一个二维切片
       card:=[]int{}
       countcolor:=make([]int,4)
       countnumber:=make([]int,15)
       color:=[]int{}
       number:=[]int{}
   for _,v:=range cards{
        card=append(card,v)
       c:=v&3
       color=append(color,c)
       countcolor[c]++
       n:=v>>2
       number=append(number,n)
       countnumber[n]++
   }
   c.SetAll(color,number,card,countcolor,countnumber)
}
  
  

}