package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Joke struct{
	ID uint32 `json:"id"`
	Joke string `json:"joke"`
}
type JokeResponse struct{
	Type string `json:"type"`
	Value Joke `json:"value"`
}
  
//parsing json GETting info
func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
		defer r.Body.Close()

		c := http.Client{}
		resp, err := c.Get("http://api.icndb.com/jokes/random?limitTo=[nerdy]")
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return 
		}

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		joke := JokeResponse{}

		err = json.Unmarshal(body, &joke)
		if err != nil {
			log.Fatal(err)
		}

		w.Write([]byte(joke.Value.Joke))
	})
	fmt.Println("Connect - http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}