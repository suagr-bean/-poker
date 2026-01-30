package service

import (
	"poker/model"
	"poker/Service"
)
type TexasGame struct{

}
func (T*TexasGame)MakeJudge()model.Judge{
 return NewJudge()
}
func (T*TexasGame)MakeDealer()model.Dealer{
   return Service.NewTexasDealer()
}