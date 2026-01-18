package service

import "poker/model"
type FullHouse struct{
	Max int
	TwoPair int
}
func(f*FullHouse)Processing(d*model.Detail)(bool,int){
	three:=0
	two:=0
	for i:=14;i>=2;i--{
	 if d.CountNumber[i]==3{
       three=i
	   break
	 }
	 if three>0{
     if  i!=three&&d.CountNumber[i]>=2{
		two=i
		break
	 }
	 }    
}
    if three!=0&&two!=0{
	  score:=70000+three*3+two*2
	  return true,score
	}
   return false,0
}