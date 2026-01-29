package ranges

import "sort"



func SortFile(hand AllHand)AllHand{
   sort.Slice(hand.HandIndex,func(i int, j int) bool{
     return hand.HandIndex[i].Win>hand.HandIndex[j].Win
   })
   return hand
}


