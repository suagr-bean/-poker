package service

func (T*TexasJudge)DealMask(){
  for i:=0;i<len(T.Count.Id);i++{
	number:=T.Count.Id[i].Number
    
	color:=T.Count.Id[i].Color
    
	var mask uint16
	var maskcolor [4]uint16
	for j:=0;j<len(number);j++{
		n:=number[j]
		c:=color[j]
		deal:=uint16(1)<<uint(n)//掩码左移
		mask |=deal
		maskcolor[c]|=deal
		
	  T.Count.Id[i].MaskNumber=mask
	  T.Count.Id[i].MaskColor=maskcolor
	}
  }
}