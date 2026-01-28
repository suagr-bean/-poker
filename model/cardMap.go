package model
var SuitMap=map[string]int{
  "C":0,
   "S":1,
   "D":2,
   "H":3,
}
var NumberMap=map[string]int{
	"2":2,
	"3":3,
	"4":4,
	"5":5,
	"6":6,
	"7":7,
	"8":8,
	"9":9,
	"10":10,
	"J":11,
	"Q":12,
	"K":13,
	"A":14,
}
var Suit=make(map[int]string)
var Number=make(map[int]string)
func InitMap(){
	for i,v:=range SuitMap{
		Suit[v]=i
	}
	for j,k:=range NumberMap{
		Number[k]=j
	}
}
func SuitNumber(color int,number int)string{
	 stringSuit:=Suit[color]
	 stringNumber:=Number[number]
	 return stringSuit+stringNumber
}
type HandInfo struct {
	Hand []string
	Win string
	
}
var HandMap=make(map[int]HandInfo)
func SetHandMap(id int, hand[]string,win string){
	HandMap[id]=HandInfo{
		Hand:hand,
		Win :win,
	}
}
func GetMap()map[int]HandInfo{
	return HandMap
}