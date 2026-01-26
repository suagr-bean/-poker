package service

import "poker/model"
type OnePair struct{

}
func(one*OnePair)Processing(d*model.Detail)(bool,int){
	number:=d.CountNumber
	pair:=0
	k1:=0
	k2:=0
	k3:=0
	weight:=[]int{50,10,1}
	for i:=14;i>=2;i--{
	   if number[i]==2{
		pair=i
		break
	   }
	}
	if pair>0{
		for i:=14;i>2;i--{
		   if i!=pair{
			if k1==0{
			  k1=i
			}else if k2==0{
				k2=i
			}else if k3==0{
				k3=i
				break
			}
		   }
		}
	}
	if pair>0&&k1>0{
		score:=20000+pair*500+k1*weight[0]+k2*weight[1]+k3*weight[2]
	 return true,score
	}
	return false,0
}