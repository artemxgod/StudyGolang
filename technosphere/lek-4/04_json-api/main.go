package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	Name	string 	`json:"name"`
	Done	bool	`json:"done"`
}

func main(){
	todos:= []Todo{
		{"learn Go", false},
	}

	http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request){
		//giving a static file which will take API from browser
		//open file
		fileContent, err := ioutil.ReadFile("index.html")
		if err != nil{
			log.Println("index.html - not found",err)
			w.WriteHeader(http.StatusNotFound)
			return 
		}
		//output file
		w.Write(fileContent)
	})

	http.HandleFunc("/todos/", func(w http.ResponseWriter, r * http.Request){
		fmt.Println("Request ", r.URL.Path)
		defer r.Body.Close()

		switch r.Method{
		//GET - getting data
		case http.MethodGet: 
			//turn structure into json
			productsJson, err := json.Marshal(todos)
			if err != nil{
				log.Fatal("issue incrypting to json")
			}
			//we must let browser know that we are sending a json
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			//output json
			w.Write(productsJson)
			
		//POST - add smth
		case http.MethodPost:
			decoder := json.NewDecoder(r.Body)
			todo := Todo{}
			//turn json request into a structure
			err := decoder.Decode(&todo)
			if err != nil{
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			todos = append(todos, todo)
		//PUT - to update existing info
		case http.MethodPut: 
			id := r.URL.Path[len("/todos/"):]
			index, _ := strconv.ParseInt(id, 10, 0)
			todos[index].Done = true 

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	fmt.Println("connect - http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}