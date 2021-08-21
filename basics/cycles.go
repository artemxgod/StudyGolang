//in golang we have only cycle 'for'
package main


import "fmt"

func main(){
	//Task 1 - sum from a to b
	var a, b, sum int

	sum = 0
	fmt.Scan(&a, &b)
	for i := a; i<=b; i++{
		sum += i
	}
	fmt.Print(sum)

	//-----------------------------------\\

	// Task 2 - sum of nums multiples to 8
	var num, cnt, sum1 int

	sum = 0
	cnt = 0
	fmt.Scanln(&cnt)
	for i := 0; i< cnt; i++{
		fmt.Scan(&num)
		if num % 8 == 0 && num < 100 && num > 10{
			sum1 += num
		}
	}
	fmt.Println(sum1)

	//-----------------------------------\\
	//Task 3 - cnt of max

	var cnt1, num1, max int
	
	max = 0
	for fmt.Scan(&num1); num1 != 0; fmt.Scan(&num1){
		if num1 > max{
			max = num1
			cnt1 = 1
		}else if num1 == max{
			cnt1++
		}
	}
	fmt.Print(cnt1)
}