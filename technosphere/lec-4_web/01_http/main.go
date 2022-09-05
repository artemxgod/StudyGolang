package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"fmt"
)

func main(){
	mux := new(http.ServeMux)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		path := r.URL.Path

		//creating http client
		c := http.Client{}
		//appealing to host artii..
		resp, err := c.Get("http://artii.herokuapp.com/make?text=" + path)
		if err != nil{
			log.Fatal(err)
		}
		//have to close body after reading
		defer resp.Body.Close()

		//parsing host's body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil{
			log.Fatal(err)
		}
		//set status OK
		w.WriteHeader(http.StatusOK) 
		w.Write(body)

	})
	fmt.Println("connect - http://127.0.0.1:8000")
	http.ListenAndServe(":8000", mux)
	
}