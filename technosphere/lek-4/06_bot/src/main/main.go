package main

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)


type Joke struct {
	ID 		uint32 	`json:"id"`
	Joke 	string	`json:"joke"`
}

type JokeResponse struct{
	Type	string	`json:"type"`
	Value	Joke	`json:"value"`	
}

func main(){
	 
}
