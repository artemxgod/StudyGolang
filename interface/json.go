package main

import (
	"os"
	"encoding/json"
	"fmt"
	"io/ioutil"
)


type myStruct struct {
	// при кодировании / декодировании будет использовано имя name, а не Name
	Name string `json:"name"`
	
	// при кодировании / декодировании будет использовано то же имя (Age),
	// но если значение поля равно 0 (пустое значение: false, nil, пустой слайс и пр.),
	// то при кодировании оно будет опущено
	Age int `json:",omitempty"`
	
	// при кодировании / декодировании поле всегда игнорируется
	Status bool `json:"-"`
}

type (
	Group struct{
		Students [] struct{
			Rate []int `json:"Rating"`
		}
	}
)
type FindId struct{
	ID int `json:"global_id"`
}
//task 1 - find avg rating, 2 solutions
func Avg_rating(){
	var group Group
	file, err := os.Open("interface/stud.json")
	if err != nil{
		panic("Unable to open file")
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&group); err != nil { //.json to struct
		panic(err)
	}
	
	var avg float64
	for _, st := range group.Students {
		avg += float64(len(st.Rate))
	}
	avg /= float64(len(group.Students))
	
	e := json.NewEncoder(os.Stdout) // 
	e.SetIndent("", "    ")
	e.Encode(struct {Average float64}{avg})
}
func New_Avg_rating(){
	var avg float64 = 0
	var obj Group

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil{
		panic("trouble with reading .json file")
	}

	if err := json.Unmarshal(data, &obj); err != nil{
		panic("error writing data in object")
	}
	for _, studs := range obj.Students{
		
		avg += float64(len(studs.Rate))
	}
	avg /= float64(len(obj.Students))
	
	data, err = json.MarshalIndent(struct {Average float64}{avg}, "", "    ")
	_, err = os.Stdout.Write(data)
}

//task 2 - find sum of id's in .json file
func ID_sum(){
	
	file, err := os.Open("interface/decode.json")
	if err != nil{
		panic("Unable to open file")
	}
	defer file.Close()

	var sum int64 = 0
	var id []FindId
	if err != nil {
		panic("Unable to open file")
	}
	if err := json.NewDecoder(file).Decode(&id); err != nil { 
		panic(err)
	}
	for _, elem := range id{
		sum += int64(elem.ID)
	}
	fmt.Print(sum)
}

func main(){
	ID_sum()
}
