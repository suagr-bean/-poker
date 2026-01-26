package service

import (
	"math/bits"
	"poker/model"
	"poker/utils"
)
type StraightFlush struct{
  
}
func (s*StraightFlush) Processing(d*model.Detail)(bool,int){
	for i:=0;i<4;i++{
    mask:= d.MaskColor[i]
	 if bits.OnesCount16(mask)<5{
		continue
	 }
	result,max:=utils.DealStright(mask)
	 if result==true{
		score:=90000+max
       return true,score
	}

}
   return false,0
}