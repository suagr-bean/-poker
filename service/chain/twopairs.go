package service

import "poker/model"

type TwoPairs struct {
}

func (t *TwoPairs) Processing(d *model.Detail)(bool,int){
	number:=d.CountNumber
	max:=0
	second:=0
	kicker:=0
	for i:=14;i>=2;i--{
	  if number[i]==2{
		if max==0{
			max=i
		}else if second==0{
			second=i
		}
	  }
	}
	if max!=0&&second!=0{
		for i:=14;i>=2;i--{
			if i!=max&& i!=second&&number[i]==1{
				kicker=i
				break
			}
		}
    }
	if kicker!=0{
		score:=30000+max*500+second*20+kicker
	 return true,score
	}
	return false,0
}