package chain

import (
	"math/bits"
	"poker/model"
	"poker/Utils"
)
type StraightFlush struct{
  
}
func (s*StraightFlush)Node(d*model.Count)(bool,int){
	for i:=0;i<4;i++{
    mask:= d.MaskColor[i]
	 if bits.OnesCount16(mask)<5{
		continue
	 }
	result,max:=Utils.DealStright(mask)
	 if result==true{
		score:=90000+max
       return true,score
	}

}
   return false,0
}