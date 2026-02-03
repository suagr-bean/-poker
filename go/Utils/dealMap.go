package Utils

import "poker/model"

//把数据过map
func DealMap(pool[]string)[]int{
	list:=make([]int,0)
	
	for i:=0;i<len(pool);i++{
	    deal:=pool[i]
		dataColor:=deal[:1]
		dataNumber:=deal[1:]
        color:=model.SuitMap[dataColor]
		number:=model.NumberMap[dataNumber]
		wise:=number<<2|color
		
		list=append(list,wise)
	
	}
	return list
}