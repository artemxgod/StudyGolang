package main

import (
	"fmt"
	"strings"
)
func work(num int) int{
	if num < 4{
		return num - 1
	}else{
		return num + 1
	}
}
func main(){

	//Task 1 - time limit
	var workArray [10]int
	s:= map[int]int{
		2 : 1, 
		3 : 2,
		4 : 5,
		5 : 6,
		6 : 7,
		7 : 8,
	}
	for i := 0; i < 10;i++{
		fmt.Scan(&workArray[i])
		if s[workArray[i]] != 0{
		fmt.Print(s[workArray[i]]," ")
		}else{
			 fmt.Print(work(workArray[i])," ")
			 }
		}

	//Task 2 - wrong database
	groupCity := map[int][]string{
		10:   []string{"Деревня", "Село"}, // города с населением 10-99 тыс. человек
		100:  []string{"Город", "Большой город"}, // города с населением 100-999 тыс. человек
		1000: []string{"Миллионик"}, // города с населением 1000 тыс. человек и более
	}
	 cityPopulation := map[string]int{
		"Село" : 100,
		"Миллионик" : 500,
		"Город" : 600,
		//город: население города в тысячах человек,
	}
	for idx, _ := range cityPopulation{
		for _, elem := range groupCity[10]{
			if idx == elem {
				delete(cityPopulation, idx)
			}
		}
	}
	for idx, _ := range cityPopulation{
		for _, elem := range groupCity[1000]{
			if idx == elem {
				delete(cityPopulation, idx)
			}
		}
	}	
	fmt.Print(cityPopulation)
}