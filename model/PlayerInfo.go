package model
type Game interface{
	MakeJudge()Judge
	MakeDealer()Dealer
}
//玩家信息
type PlayerInfo struct{
   Players []*Player
   Person int 
}
type Player struct{
   Hand [] int
   PlayerRange int
}
func NewPlayerInfo(person int)*PlayerInfo{
	G:=&PlayerInfo{Players:make([]*Player,0,person),
		Person:person}
    G.Init()
	return G
}
func (g*PlayerInfo)Init(){
	for i:=0;i<g.Person;i++{
		p:=&Player{Hand:make([]int,0),
		PlayerRange:100}
       g.Players=append(g.Players,p)
	}
}
func (g*PlayerInfo)AddHand(index int,hand[]int){
	g.Players[index].Hand=hand
	
}
func (g*PlayerInfo)AddRange(index int,ranges int){
	g.Players[index].PlayerRange=ranges
}