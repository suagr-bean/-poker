package service

import (
	"poker/model"
	"poker/utils"
	
)
type TexasGame struct{

}
func (T*TexasGame)MakeJudge()model.Judge{
 return NewJudge()
}
func (T*TexasGame)MakeDealer()model.Dealer{
   return utils.NewTexasDealer()
}