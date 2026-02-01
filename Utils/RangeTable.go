package Utils

import "log"

//读文件的起手表
func RangeTable()AllHand{
	path:="/workspaces/-poker/poker/go/RangeTable/sortvs1.jsonl"
	allhand,error:=ReadFile(path)
	if error!=nil{
	 log.Println("错误",error)
		//这里捕获错误如果没读到文件啥的
	}
	return allhand
}