package main

import (
	"fmt"
	"stopwatch/swer"
	"time"
)

func main() {
	var w swer.Sw = swer.NewPpStopWatch()
	w.Start("step1")
	time.Sleep(time.Second)
	w.Stop()
	time.Sleep(time.Second * 2)
	w.Start("step1")
	time.Sleep(time.Second * 2)
	w.Stop()
	fmt.Println(w.PrettyPrint())
}
