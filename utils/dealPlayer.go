package utils

import "poker/model"




func(T*TexasDealer) Deal (){
     
  for i:=0;i<T.Data.Person;i++{
	 action:=T.Data.Players[i].Action
	 hand:=T.Data.Players[i].Hand
	 boardCard:=T.Board.GetBoardCards()
	 if action==1{
	  process:=append([]int(nil),boardCard...)
	  process=append(process,hand...)
	  e:=model.NewEntryInfor(process)
	  
	  T.Board.Entry=append(T.Board.Entry,e)
	 } 
  }


}