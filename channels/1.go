package main

// code with ryan
import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"sync"
)

func unbufferedBlock() {
	dataChan := make(chan int)

	dataChan <- 20
	n := <-dataChan

	log.Println("unbufferedBlocking: ", n)
}

func bufferedBlock() {
	// more input than channel size: deadlocks
	// equal or more size than input: no deadlock
	dataChan := make(chan int, 1)

	dataChan <- 20
	dataChan <- 30

	n := <-dataChan
	log.Println("bufferedBlocking: ", n)

	// dataChan <- 30 // put this here does not deadlock when size of chan is 1
	n = <-dataChan
	log.Println("bufferedBlocking: ", n)
}

func unbufferedBlockSol1() {
	// use threads
	dataChan := make(chan int)

	go func() {
		dataChan <- 20

	}()
	n := <-dataChan

	log.Println("unbufferedBlockSol1: ", n)

}

func blockChanRange() {
	dataChan := make(chan int)

	// dataChan <- 20 // cause deadlock

	go func() {
		for i := 0; i < 1000; i++ {
			dataChan <- i
		}
	}()

	// causes deadlock on last input as the dataChan as its still trying to receive 
	// when go routine is done
	for n := range dataChan {
		log.Println("blockChanRange: ", n)

	}
	log.Println("DONE")

}


func doWork(i int) string {
	t := rand.Intn(2)
	time.Sleep(time.Duration(t) * time.Second)
	return fmt.Sprintf("%d: %d", i, t)
}

func unblockChanRangeSequential() {
	dataChan := make(chan string)

	// dataChan <- 20 // cause deadlock

	go func() {
		for i := 0; i < 1000; i++ {
			dataChan <- doWork(i)
		}
		close(dataChan)
	}()

	for n := range dataChan {
		log.Println("unblockChanRange: ", n)

	}
	log.Println("DONE")

}


func unblockChanRangeParallel() {
	// not ordered
	dataChan := make(chan string)

	// dataChan <- 20 // cause deadlock
	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done() //without this deadloock occurs
				dataChan <- doWork(i)
			}(i)
		}
		wg.Wait()
		close(dataChan)
	}()

	for n := range dataChan {
		log.Println("unblockChanRange: ", n)

	}
	log.Println("DONE")

}

func main() {
	// unbufferedBlock()
	// bufferedBlock()
	// unbufferedBlockSol1()
	// blockChanRange()
	// unblockChanRangeSequential()
	unblockChanRangeParallel()
}
