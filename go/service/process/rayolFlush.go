package service

import "poker/model"
type RoyalFlush struct{

}

func (r*RoyalFlush) Processing(d*model.Detail)(bool,int){
	const mask uint16=0x7c00
    for i:=0;i<4;i++{
		if (d.MaskColor[i]&mask)==mask{
          return true,100000
		}
	}
	return false,0
}