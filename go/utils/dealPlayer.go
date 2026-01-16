package utils


func(T*TexasDealer) Deal ()[][]int{
	all:=[][]int{}
	play:=T.Data.GetAllPlayer()
    cards:=T.Board.GetBoardCards()
  for _,v:=range play{
	
	process:=append([]int(nil),cards...)
    process=append(process,v.Hand...)
	all=append(all,process)
  }
  return all
}