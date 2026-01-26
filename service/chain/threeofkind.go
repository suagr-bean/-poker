package service
import "poker/model"
type ThreeOfKind struct{

}
func (t*ThreeOfKind)Processing(d*model.Detail)(bool,int){
    number:=d.CountNumber
	three:=0
	kicker1:=0
	kicker2:=0
	for i:=14;i>=2;i--{
	 if number[i]==3{
       three=i
	   break
	 }
	}
	if three>0{
	  for i:=14;i>=2;i--{
       if i!=three&&number[i]>0{
		  if kicker1==0{
			kicker1=i
		  }else if kicker2==0{
             kicker2=i
			 break
		  }
		  
	   }
	  }
	}
	 if three!=0&&kicker1!=0&&kicker2!=0{
		score:=40000+kicker1*500+kicker2*1
		return true,score
	 }

	return false,0
}