package Service

import (
	
	"poker/model"
)
type TexasGame struct{

}
func (T*TexasGame)MakeJudge()model.Judge{
 return NewTexasJudge()
}
func (T*TexasGame)MakeDealer()model.Dealer{
   return NewTexasDealer()
}