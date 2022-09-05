package main

import (
	"log"
	"os"

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
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT-TOKEN"))
	if err != nil {
		log.Fatal(err)
	} 
	log.Printf("Authorized on account %s", bot.Self.UserName)

}
