package Service

import (
	"poker/model"
)
type TexasGame struct{

}
func (T*TexasGame)MakeJudge()model.Judge{
 return NewJudge()
}
func (T*TexasGame)MakeDealer()model.Dealer{
   return NewTexasDealer()
}