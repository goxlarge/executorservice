package main

import (
	"fmt"
	"time"

	"github.com/goxlarge/executorservice/scheduledservice"
)

type R struct {
	S string
}
func (r *R) Run(){
	fmt.Println(r.S)
}

func main(){
	a := &R{"hello"}
	// b := &R{"world"}
	// ch := scheduledservice.Schedulejob(5, a)
	//  scheduledservice.Schedulejob(1, b)
	// <- ch
	//<- ch2

	 scheduledservice.SchedulejobWith(1, a, 10)
	

	time.Sleep(20* time.Second)
}
