package main

import "fmt"
//formatted output
func main(){
	// var num float64

	// fmt.Scan(&num)
	// if num <= 0 {
	// 	fmt.Printf("число %2.2f не подходит", num)
	// }else if num > 10000{
	// 	fmt.Printf("%e", num)
	// }else { 
	// 	fmt.Printf("%.4f", num*num)
	// }

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


}
