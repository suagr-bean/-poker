package model
type Game interface{
	MakeJudge()Judge
	MakeDealer()Dealer
}
type GameData struct{
   Players []*Player
   Person int 
}
func NewGameData(person int)*GameData{
	G:=&GameData{Players:make([]*Player,0,person),
		Person:person}
    G.SetPlayer()
	return G
}
func (g*GameData)SetPlayer(){
	for i:=0;i<g.Person;i++{
		p:=&Player{Id:i,Hand:make([]int,0)}
       g.Players=append(g.Players,p)
	}
}
func (g*GameData)Add(id int,hand[]int)*GameData{
      g.Players[id].Hand=hand
	return g
}
func (g*GameData)GetAllPlayer()[]*Player{
   return g.Players
}
func (g*GameData)GetIndex(id int)*Player{
   return g.Players[id]
}

