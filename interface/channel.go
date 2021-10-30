package main

import (
	"fmt"
	"time"
	"sync"
)
func work() {
	fmt.Println("Goroutine start")
	time.Sleep(time.Second * 1)
	fmt.Println("Goroutine stop")
}

func task(c chan int, N int){
	c <- N+1
}

func task2(c chan string, s string){
	s += " "
	for i:= 0; i<5; i++{
		c <- s
	}
}

func removeDuplicates(inputStream, outputStream chan string){
	prev := ""
	for v := range inputStream{
		if v != prev{
			outputStream <- v
			prev = v
		}
	}
	close(outputStream)
}

func waitResult(){
	done := make(chan struct{})
	go func(){
		//work()
	close(done)
	}()
	<-done
}

func WaitRes(){
	wg := new(sync.WaitGroup)
	for i:= 0; i<10; i++{
		wg.Add(1)
		go func(){
			work()
			defer wg.Done()
		}()
	}
	wg.Wait()
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		select{
		case num := <-firstChan:
			c <- num*num
		case num := <-secondChan:
			c <-num*3
		case <-stopChan:
			return
		default: 
			return
		}
	}()
	return c
}

func calculator2(arguments <-chan int, done <-chan struct{}) <-chan int {
	c := make(chan int)
	go func(){
		var sum int = 0
		defer close(c)
		for {
			select{
			case num := <- arguments:
				sum += num
			case <-done:
				c <- sum
				return
			}
		}
	}()
	return c
}

func merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int) {
		done := make(chan bool, n*2)
		results := make([]*int, n*2)

		go func() {
			nDone := 0
			for <-done {
				for i := nDone; i < n; i++ {
					if results[i] != nil && results[i+n] != nil {
						out <- (*results[i] + *results[i+n])
						if nDone++; nDone == n {
							return
						}
					} else {
						break
					}
				}
			}
		}()
		input := func(ch <-chan int, results []*int) {
			for i := 0; i < n; i++ {
				x := <-ch
				go func(i int, x int) {
					result := f(x)
					results[i] = &result
					done <- true
				}(i, x)
			}
		}
		go input(in1, results[:n])
		go input(in2, results[n:])

	}(f, in1, in2, out)
}
func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	// since there could be a long stream of numbers coming from the channels.
	results := make([]int, n)
	go func() {
		// creating a wg with n elements for each calculation.
		wg := new(sync.WaitGroup)
		wg.Add(n)
		for i := 0; i < n; i++ {
			// reading from the channels.
			x1 := <-in1
			x2 := <-in2
			index := i
			go func() {
				// creating channels for concurrent calculations.
				x1res := make(chan int)
				x2res := make(chan int)
				// concurrent calculations.
				go func() {
					x1res <- f(x1)
				}()
				go func() {
					x2res <- f(x2)
				}()
				// writing the result.
				results[index] = (<-x1res) + (<-x2res)
				wg.Done()
			}()

		}
		wg.Wait()
		// sending results to the output channel.
		for i := 0; i < n; i++ {
			out <- results[i]
		}
	}()
}
func main(){

}