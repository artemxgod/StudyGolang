package main

import (
	"html/template"
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)


func main(){

	tmpl := template.New("main")
	tmpl, _ = tmpl.Parse(
		`<div style="dislplay: inline-block; border: 1px solid #aaa;
	border-radius: 3px; padding: 30px; margin: 20px;">
		<pre>{{.}}</pre>
		</div>`)
	
	mux := new(http.ServeMux)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		path := r.URL.Path

		//creating http client
		c := http.Client{}
		//appealing to host artii..
		resp, err := c.Get("http://artii.herokuapp.com/make?text=" + path)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error"))
			return 
		}
		//have to close body after reading
		defer resp.Body.Close()

		//parsing host's body
		body, _ := ioutil.ReadAll(resp.Body)
		
		//giving template with data
		err = tmpl.Execute(w, string(body))
		if err != nil{
			fmt.Println("template executing a template")
		}

	})
	fmt.Println("connect - http://127.0.0.1:8000")
	http.ListenAndServe(":8000", mux)
}