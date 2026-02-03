package model
type Game interface{
	MakeJudge()Judge
	MakeDealer()Dealer
}
//玩家信息
type PlayerInfo struct{
   Players []*Player
   Person int 
   SkipId int
}
type Player struct{
   Hand [] int
   PlayerRange int
   Score int
}
func NewPlayerInfo(person int)*PlayerInfo{
	G:=&PlayerInfo{Players:make([]*Player,0,person),
		Person:person}
    G.init()//内部初始化
	return G
}

func (g*PlayerInfo)init(){
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
