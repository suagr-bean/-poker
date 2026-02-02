package model
type Judge interface{
	Init(p*PlayerInfo,b*Board)Result
}