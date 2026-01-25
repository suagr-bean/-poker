package service

import "poker/model"
type FourOfKind struct{
   
}
func (f*FourOfKind)Processing(d*model.Detail)(bool,int){
	 Max:=0
	 Kicker:=0
	for i:=14;i>=2;i--{
     if d.CountNumber[i]==4{
		Max=i
		break
	 }
	}
	if Max>0{
	for i:=14;i>=2;i--{
      if Max!=i&&d.CountNumber[i]>0{
		Kicker=i
		break
	  }
	}
	if Max!=0{
	score:=80000+Max*4+Kicker
	return true,score
    }
	
}
   return false,0
}