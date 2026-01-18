package service

import (
	"math/bits"
	"poker/model"
)

type Flush struct {
}

func (f *Flush) Processing(d *model.Detail)(bool,int){
	for i:=0;i<4;i++{
	mask:=d.MaskColor[i]
	if bits.OnesCount16(mask)>=5{
		score:=60000
		count:=0
		weight:=[]int{600,45,10,3,1}
      for v:=14;v>=2;v--{
       if mask&(uint16(1) <<uint(v))!=0{
         score+=v*weight[count]
		 count++
		 if count==5{
			break
		 }
	   }
	  }
	  return true,score
	}
	}

	return false,0
}