package ranges

import (
	
	"poker/model"
	"poker/utils"
	
)


func Combination (){
	deck:=utils.NewDeck()
	cards:=deck.AllCards
     model.InitMap()
	 
	 dir:="../../Range"
	 fileName:="handvs1.jsonl"

	 url:="https://fluffy-fortnight-jj64qv579r7g2q4xr-8080.app.github.dev/cal"
	for i:=len(cards)-1;i>=1;i--{
		for j:=i-1;j>=0;j--{
           Suit1:=cards[i]&3
		   Suit2:=cards[j]&3
		   Number1:=cards[i]>>2
		   Number2:=cards[j]>>2 
		 hand1:= model.SuitNumber(Suit1,Number1)
         hand2:=model.SuitNumber(Suit2,Number2)
		 hand:=[]string{}
		hand=append(hand, hand1)
		hand=append(hand,hand2) 
		
         data:=CalHand(hand)
		result,_:= PostStart(url,data)
		
		 
		 SaveFile(dir,fileName,hand,result)
		}
	}  
	  Compt()
	 
}