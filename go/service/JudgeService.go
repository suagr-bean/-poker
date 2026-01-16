package service

type TexasJudge struct {
  Deal [][]int
   
}

func NewJudge()*TexasJudge{
	return &TexasJudge{}
}
func(t*TexasJudge) InitCard(card[][]int){
  t.Deal=card
  t.DealCount()//计数桶
  
}
