package Utils
//掩码方便计算 顺子的
func DealMask(color[]int,number[]int)(uint16,[4]uint16){
    
	var mask uint16
	var maskcolor [4]uint16
	for j:=0;j<len(number);j++{
		n:=number[j]
		c:=color[j]
		deal:=uint16(1)<<uint(n)//掩码左移
		mask |=deal
		maskcolor[c]|=deal
		
	}
	return mask ,maskcolor
  }
