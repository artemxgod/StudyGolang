package main

import "fmt"
//formatted output
func main(){
	var num float64

	fmt.Scan(&num)
	if num <= 0 {
		fmt.Printf("число %2.2f не подходит", num)
	}else if num > 10000{
		fmt.Printf("%e", num)
	}else { 
		fmt.Printf("%.4f", num*num)
	}

	//Array and slice
	//	Task 1
	var(workArray [10]uint8 
		idx1, idx2 int
	)
	
	for i := 0; i < len(workArray); i++{ // input array
		fmt.Scan(&workArray[i])
	}
	
	for i := 0; i < 3; i++{ //	switching elements
		fmt.Scan(&idx1, &idx2)
		workArray[idx1],workArray[idx2] = workArray[idx2],workArray[idx1]
	}
	
	for i := 0; i < len(workArray); i++{ //	output array
		fmt.Print(workArray[i], " ")
	}

	//Task 2 - slices
	var	N,value int
	
	baseslice := make([]int, 4, 4)
	fmt.Scan(&N)
	for i := 0; i<N;i++{
		if i < 4{
			fmt.Scan(&baseslice[i])
		}else{
			fmt.Scan(&value)
			baseslice = append(baseslice, value)
		}
	}
	fmt.Print(baseslice[3])
	
	//Task 3 - find and output max elem
	array := [5]int{}
	var a int
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
	}
	var max int

	max = -1000

	for i:=0;i < 5; i++{
		if max < array[i]{
			max = array[i]
		}
	}
	fmt.Print(max)
	
	//Task 4 - output even elems
	var( N int
		baseArray [100]int)
	
	fmt.Scan(&N)
	for i:=0; i < N; i++{
		fmt.Scan(&baseArray[i])
	}
	for i:=0; i < N; i++{
		if(i % 2 == 0){
			fmt.Print(baseArray[i], " ")
		}
	}

	//Task 5 - amount of positive numbers
	var (N,cnt int
		array [100]int)

		cnt = 0
		fmt.Scan(&N)
		for i:=0; i < N; i++{
			fmt.Scan(&array[i])
			if(array[i] >= 0){
				cnt++
			}
		}
		fmt.Print(cnt)
}
