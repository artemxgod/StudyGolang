package main

import (
	"fmt"
	"math"
)

func DigPow(n, p int) int {
	var sum int
	tmp := reverse_num(n)
	for ; tmp > 0; tmp /= 10 {
		sum += int(math.Pow(float64(tmp % 10), float64(p)))
		p++
	}
	if sum % n == 0 {
		return sum / n
	}
	return -1
}

func reverse_num(num int) int {
	var res int
	for ; num > 0; num /= 10 {
		res += num % 10; res *= 10
	}
	return res / 10
}

func Number(busStops [][2]int) int {
	var In, Out int
	for _, elem := range busStops {
		In += elem[0]
		Out += elem[1]
	}
	if In > Out {
		return In - Out
	} else {
		return 0
	}
}

func main() {
	fmt.Println(Number([][2]int{{3,0},{9,1},{4,8},{12,2},{6,1},{7,8}}))
}
