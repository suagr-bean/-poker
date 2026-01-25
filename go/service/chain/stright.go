package service

import (
	"poker/model"
	"poker/utils"
)

type Stright struct {
}

func (s *Stright) Processing(d *model.Detail)(bool,int){
	mask:=d.MaskNumber
	is,max:=utils.DealStright(mask)
	score:=50000
	if is{
     score+=max*10
	 return true,score
	}
	return false,0
}