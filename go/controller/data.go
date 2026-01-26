package controller

import (
	"fmt"
	"poker/model"
	"poker/model/httpModel"
	"poker/service"
	"poker/utils"
	"sync"
	"time"
)
func Begining(cal*httpModel.CalData)float32{
	start:=time.Now()
	 stats:=&model.Stats{}
	 record:=func(score int){
		service.Note(stats,score)
	 }
	var req float32
	 Hand:=cal.Cards.Hand
	 dealHand:=utils.DealMap(Hand)
     dealPublic:=utils.DealMap(cal.Cards.Hand)
	 fmt.Println(dealHand)
	win:=float32(0)
	beginner:=&model.Begin{}
	beginner.Hand=dealHand
	beginner.Id=cal.Cards.Position
	beginner.PublicCard=dealPublic
	beginner.Person=cal.Table.Person
    beginner.Frequency=10000
	
	job:=make(chan struct{},200)
	chanResult:=make(chan float32,200)
	worker:=2
     var wait sync.WaitGroup
	wait.Add(worker)
	for w:=0;w<worker;w++{
		
		go func (workId int){
		 defer wait.Done()
          for range job{
			var single float32
			service.GameMain(beginner,func(s float32){
				single+=s
			},record)
			chanResult <- single
		  }
		}(w)
	}
	go func(){
		for i:=0;i<beginner.Frequency;i++{
			job <- struct{}{}
		}
		close(job)
	}()
    go func(){
      wait.Wait()
	  close(chanResult)
	}()
	 for r:=range chanResult{
		win+=r
	 }
	req=(float32(win)/float32(beginner.Frequency))*100
	fmt.Println(req)
	fmt.Println(stats)
	end:=time.Since(start)
    fmt.Println(end)

    return req
}