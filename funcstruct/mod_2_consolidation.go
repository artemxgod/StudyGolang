package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func hypotenuse(a,b float64) float64{
	return math.Sqrt(a*a + b*b)
}

func Star(s string) string{
	Rune1 := []rune(s)
	amount := utf8.RuneCountInString(s)
	Rune2 := make([]rune, (amount*2-1))
	j := 0
	for i:= 0; i< amount; i++{
		Rune2[j] = Rune1[i]
		if j < amount*2 - 2{
		Rune2[j+1] = rune('*')
		}
		j += 2
	}
	return string(Rune2)
}

func Max_digit(s string) uint8{
	var max uint8 = 0

	a := []byte(s)
	for _, value := range a{
		if value-48 > max{
			max = (value-48)
		}
	}
	return max
}

func Square(s string) {
	for _, tmp := range s{
		fmt.Print((tmp - '0')*(tmp - '0'))
	}
}

func T(k, p, v float64) float64{
	return 6 / W(k,p,v)
}
func W(k, p, v float64) float64{
	return math.Sqrt(k / M(p, v))
}
func M(p, v float64) float64{
	return p*v
}

func main(){
	//Task 1 - find hypotenuse
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Print(hypotenuse(a,b))

	//Task 2 - put * between
	var s string
	fmt.Scan(&s)
	fmt.Print(Star(s))

	//Task 3 - max digit
	var s string
	fmt.Scan(&s)
	fmt.Print(Max_digit(s))
	
	//Task 3 - ^2 every digit
	var s string 
	fmt.Scan(&s)
	Square(s)

	//Task 4 - pendulum swing
	fmt.Print(T(1296, 6, 6))

}