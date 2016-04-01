package main

//package esterno da scaricare
import (
	"github.com/cortinico/telebot"
)

func main() {

	var bot telebot.Bot
	conf := telebot.Configuration{
		BotName: "NameSearchbot",
		ApiKey:  "aaaaaaaaaaaaaabbccidifdfkjdflajdlfa",
	}

	// bot start è una funzione che prende le configurazioni per collegarsi
	// e una funziona (perché in go la funzione è un tipo)
	// ricordarsi che in go le funzioni sono dei tipi
	// comando defer (appena finisce il programma fa quello che gli dici di fare)

	bot.Start(conf, func(mess string) (string, error) {
		var answer string
		switch mess {
		case "/start":
			answer = "Welcome to SearchGobot!\nType a term you want to search or /test or /help"
		case "/help":
			answer = "It's simple, just type what you want to search using google"
		case "":
			answer = "Welcome to SearchGobot type a term you want to search"
		default:
			answer = SendGoogleSearchRequest(mess)
		}
		return answer, nil
	})
}
