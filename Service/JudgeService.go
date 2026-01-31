package Service

import (
	"poker/model"
 
  
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

