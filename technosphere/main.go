package main

import (
	"fmt"
	"strconv"
	"time"
)
func ReturnInt() int{
	return 1
}

func ReturnFloat() float32{
	return float32(1.1)
}

func  ReturnIntArray() [3]int{
	return [3]int{1,3,4}
}

func ReturnIntSlice() []int{
	return []int{1,2,3}
}

func IntSliceToString(arr []int) string{
	var str string
	
	for _, elem := range arr{
		str += strconv.Itoa(elem)
	}
	return str
}

func MergeSlices(farr []float32, arr []int32) []int{
	var newArr []int

	for _, elem := range(farr){
		newArr = append(newArr, int(elem))
	}
	for _, elem := range(arr){
		newArr = append(newArr, int(elem))
	}
	return newArr
}

func GetMapValuesSortedByKey(m map[int]string) []string{
	var s []string

	cnt := 0
	for i:= 0; ;i++{	
		for idx, elem := range(m){
			if cnt == len(m){
				return s
			}
			if i == idx{
				s = append(s, elem)
				cnt++
			}
		}
	}
}

//day 02
func ShowMeTheType(variable interface{}) string{

	switch variable.(type){
	case int8 : return string("int8")
	case int16: return string("int16")
	case int32:return string("int32")
	case int64: return string("int64")
	//...
	default: return string("unknown type")
	}
}

func Sqrt(n float64) float64{
	
	const eps float64 = 0.01
	var x float64 = 1
	var nx float64
	for{
		nx = (x+n/x)/2
		if((x-nx) < eps) && ((x-nx) > 0) || ((x-nx) > -eps) && ((x-nx) < 0) {
			break
		}
		x = nx
	}
	return x	
}


type Calendar struct{
	month string
	quater int
}
func (obj* Calendar) NewCalendar(t time.Time){
	
	obj.month = t.Month().String()
	switch {
	case obj.month == "January" ||
	obj.month == "February" ||
	obj.month == "March": obj.quater = 1
	case obj.month == "April" ||
	obj.month == "May"||
	obj.month == "June": obj.quater = 2
	case obj.month == "July" ||
	obj.month == "August"||
	obj.month == "September": obj.quater = 3
	case obj.month == "October" ||
	obj.month == "Nowember"||
	obj.month == "December": obj.quater = 4
	}
}

func main(){
}