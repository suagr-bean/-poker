package service

import (
	"sync"
	"poker/model"
)
var mu sync.Mutex
func Note(stats*model.Stats,score int){ 
	mu.Lock()
	defer mu.Unlock()
	switch {
	case score==100000:
		stats.RoyalFlush++
    case score>=90000:
       stats.StraightFlush++
	case score>=80000:
	   stats.FourOfKind++
	case score>=70000:
	    stats.FullHouse++
    case score>=60000:
		stats.Flush++
	case score>=50000:
		stats.Straight++
	case score>=40000:
		stats.ThreeOfKind++
	case score>=30000:
		stats.TwoPairs++
	case score>=20000:
		stats.OnePair++
	case score>=10000:
		stats.HighCard++
	}
	
}
