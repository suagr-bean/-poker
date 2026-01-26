package service

import "poker/model"
func NewChain()*model.Chain{
	chain:=&model.Chain{}
	chain.Add(&RoyalFlush{})
	chain.Add(&StraightFlush{})
	chain.Add(&FourOfKind{})
	chain.Add(&FullHouse{})
	chain.Add(&Flush{})
	chain.Add(&Stright{})
	chain.Add(&ThreeOfKind{})
	chain.Add(&TwoPairs{})
	chain.Add(&OnePair{})
	chain.Add(&HighCard{})
	return chain
}