package utils


func(t*TexasDealer) Process(){
	player:=t.Data.GetAllPlayer()
	for _,v:=range player{
		if len(v.Hand)==0{
         cards:=t.Dealing(2)
		 v.Hand=append(v.Hand,cards...)
		}
	}
	card:= t.Board.GetBoardCards()
	need:=5-len(card)
	if need!=0{
		send:=t.Dealing(need)
		t.Board.Add(send)
	}
	
}