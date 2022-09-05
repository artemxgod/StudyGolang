package main

import (
	"html/template"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Todo struct{
	Name	string
	Done	bool 
}

func IsNotDone(todo Todo) bool{
	return !todo.Done 
}

func main(){
	//создаем шаблон и передаем в него функции

	 tmpl, err := template.New("template.html").Funcs(
	 	template.FuncMap{"IsNotDone": IsNotDone}).ParseFiles("template.html")
	if err != nil{
		log.Fatal("Can not expand template ", err)
		return 
	}

	todos := []Todo{
		{"Выучить Go", false},
		{"посетить лекцию по вебу", false},
		{"...", false},
		{"Profit", false},
	}

	mux := new(http.ServeMux)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		
		if r.Method == http.MethodPost{
			//reading from urlencoded request
			param := r.FormValue("id")
			index, _ := strconv.ParseInt(param, 10, 0)
			todos[index].Done = true 
		}

		err := tmpl.Execute(w, todos)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("connect - http://127.0.0.1:8000")
	http.ListenAndServe(":8000", mux)
}

