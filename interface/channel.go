package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
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

//техносфера

//цикл горутин
func CicleGo(){
	for i:= 0; i< 10; i++{
		//функция перехватывает переменную из области видимости
		//чтобы перехватить значения требуется явно передавать его в функцию
		go func(i int){
			fmt.Println("Got - ", i)
		}(i)
	}
	fmt.Scanln()//нужно чтобы функция не закончилась раньше горутины
}

//блокировка каналов
func greet(input chan<- string){ //write only(chan <- string) 
	//бесконечный цикл
	for {
		input <- fmt.Sprintf("Я")
		input <- fmt.Sprintf("нормальный!")
	}
}
func process(input <- chan int) (res int){ //read only(<- chan int)
	//будет ждать пока поступит новое значение
	for r:= range input{
		res += r
	}
	return 
}
func BlockChan(){
	c := make(chan string)
	go greet(c)
	//пока не будет прочитано канал заблокирован
	for idx := 0; idx < 2 ;idx++{
		fmt.Println(<-c, ",", <-c)
	}
	stuff := make(chan int,7) //buffered (7 раз пишется до блокировкки)
	for i:= 0; i< 19; i+=3{
		stuff <- i
	}
	close(stuff)//не даем больше писать в канал
	fmt.Print("Res = ", process(stuff))
}

//читаем несколько каналов
func boring(msg string) <-chan string{
	c := make(chan string)
	go func(){
		for i := 0 ; ; i++{
			c <- fmt.Sprintf("%s %d",msg, i)
			time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		}
	}()
	return c
}
func fanIn(input1, input2 <-chan string) <-chan string{
	c := make(chan string)
	go func(){
		for{
		c <- <-input1
		}
	}()
	go func(){
		for{
			c <- <-input2
		}
	}()
	return c
}
func few(){
	c := fanSelect(boring("Hi!"), boring("Bye :<"))
	for i:= 0; i<10; i++{
		fmt.Println(<-c)
	}
	fmt.Println("Ok bye!")
}

//select
func fanSelect(input1, input2 <-chan string) <-chan string{
	c:= make(chan string)

	go func(){ //уменьшаем вдвое количество горутин
		for{
			select{
			case s := <-input1:
				c <- s
			case s := <-input2:
			c <- s
			}
		}
	}()
	return c
}
func Select(){ //селект ожидает данные хотя бы из одного case, проходя все кейсы
	//runtime.GOMAXPROCS(runtime.NumCPU(20)) - ограничиваем количество ядер, которое будет использоваться
	//runtime.Gosched() - говорим рантайму что выполнение можно прервать и перключиться на что-то другое
}

//deadlock
type Ball struct {hits int}

func player(name string, table chan *Ball){
	for{
	//block
	ball := <-table
	//incriment
	ball.hits++
	fmt.Println(name, ball.hits)
	//wait
	time.Sleep(100* time.Millisecond)
	//send the ball back in channel
	//program will be blocked until second player read from table
	table <- ball
	}
}
func DL(){
	table := make(chan *Ball) //тк не buffered будет по очереди
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball) //запускаем мяч, если убрать то будет deadlock
	time.Sleep(1* time.Second)
	<- table // end of the game
}

//явно завершить горутины
func Bboring(die chan bool) <-chan string { // Возвращаем канал строк только для чтения.
	c := make(chan string)
	go func() {
		for {
			select {
			case c <- fmt.Sprintf("boring %d", rand.Intn(100)):
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			case <-die:
				fmt.Println("Jobs done!")
				die <- true
				return
			}
		}
	}()
	return c
}

func Cancel() {
	die := make(chan bool)
	res := Bboring(die)

	for i := 0; i < 5; i++ {
		// Читаем из канала
		fmt.Printf("You say: %q\n", <-res)
	}
	die <- true
	// Ждем, пока все горутины закончат выполняться
	<-die
}


func main(){
}

