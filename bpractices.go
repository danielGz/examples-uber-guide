package main

import (
	"armory.io/bpractices/copy_slices"
	"armory.io/bpractices/icompliance"
	ptrs "armory.io/bpractices/ptrs_interfaces"
	"armory.io/bpractices/semaphore"
	"fmt"
	"time"
)

func main() {
    ptrs.Demo()
	icompliance.ScareTrespassers(&icompliance.GermanShepperd{})

	lottery := copy_slices.Lottery{}
	lottery.SetLuckyNumbers(copy_slices.LuckyNumbers)
	//lottery.SetLuckyNumbersProperly(copy_slices.LuckyNumbers)
	lottery.Shuffle()
	lottery.PrintLuckyNumbers()
	fmt.Println(copy_slices.LuckyNumbers)
	TestSemaphore()
}

func TestSemaphore(){
	sem := semaphore.New(3)
	doneC := make(chan bool, 1)
	totProcess := 10
	for i := 1; i <= totProcess; i++ {
		sem.Acquire()
		go func(v int) {
			defer sem.Release()
			longRunningTask(v)
			if v == totProcess {
				doneC <- true
			}
		}(i)
	}
	<-doneC
}

func longRunningTask(taskID int) {
	fmt.Println(time.Now().Format("15:04:05"),
		"Running task with ID",taskID)
	time.Sleep(2 * time.Second)
}
