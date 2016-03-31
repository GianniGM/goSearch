package main

//package esterno da scaricare
import (
	"github.com/cortinico/telebot"
)

func main() {

	var bot telebot.Bot
	conf := telebot.Configuration{
		BotName: "gSearchgobot",
		ApiKey:  "184967413:AAHTy5Piw1G4ovbVXTIy-gtOgVVdlEE1K6c",
	}

	// bot start è una funzione che prende le configurazioni per collegarsi
	// e una funziona (perché in go la funzione è un tipo)
	// ricordarsi che in go le funzioni sono dei tipi
	// comando defer (appena finisce il programma fa quello che gli dici di fare)

	bot.Start(conf, func(mess string) (string, error) {
		var answer string
		switch mess {
		case "/help":
			answer = "Welcome to SearchGobot type a term you want to search"
		default:
			answer = SendGoogleSearchRequest(mess)
		}
		return answer, nil
	})
}
