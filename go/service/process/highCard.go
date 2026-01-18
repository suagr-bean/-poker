package service

import "poker/model"
type HighCard struct{

}
func(h*HighCard)Processing(d*model.Detail)(bool,int){
	number:=d.CountNumber
	h1:=0 
    h2:=0
	h3:=0
	h4:=0
	h5:=0
    weight:=[]int{600,150,30,5,1}
	for i:=14;i>=2;i--{
	  if number[i]==1{
	  if h1==0{
		h1=i
	  }else if h2==0{
		h2=i
	  }else if h3==0{
		h3=i
	  }else if h4==0{
		h4=i
	  }else if h5==0{
		h5=i
		break
	  }
	}
}
	score:=h1*weight [0]+h2*weight [1]+h3*weight[2]+h4*weight[3]+h5*weight[4]
	return true,score
}