package service

import ("poker/ranges"
  "testing"
)

func TestSortMap(t *testing.T){

	result:=ranges.SortMap()
	deal:=result[1]
    t.Log("test",deal)
}