package service

import (
	"poker/model"
   
)

func(t*TexasJudge) DealCount(){
   count:=model.CountId{}
   t.Count=&count
       
       
   for i:=0;i<len(t.board.Entry);i++{
      cards:=t.board.Entry[i].GetEntryCards()
       card:=[]int{}
       countcolor:=make([]int,4)
       countnumber:=make([]int,15)
       color:=[]int{}
       number:=[]int{}
      for i:=0;i<len(cards);i++{
         card=append(card,cards[i])
         c:=cards[i]&3
         color=append(color,c)
         countcolor[c]++
        n:=cards[i]>>2
        number=append(number,n)
        countnumber[n]++
        
      }
       
       /* card=append(card,cards...)
       c:=v&3
       color=append(color,c)
       countcolor[c]++
       n:=v>>2
       number=append(number,n)
       countnumber[n]++
       fmt.Println(card)*/
     count.SetAll(color,number,card,countcolor,countnumber)
     
   }
   
}
   


  
