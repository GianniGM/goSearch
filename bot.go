package main

//package esterno da scaricare
import (
	"github.com/cortinico/telebot"
)

func main() {

	var bot telebot.Bot
	conf := telebot.Configuration{
		BotName: "SearchGobot",
		ApiKey:  "99999999:AAjjjjjjjjjjjjjjjjjjjjjj",
	}

	// bot start è una funzione che prende le configurazioni per collegarsi
	// e una funziona (perché in go la funzione è un tipo)
	// ricordarsi che in go le funzioni sono dei tipi
	// comando defer (appena finisce il programma fa quello che gli dici di fare)

	//aggiungi un contatore che si incremente ogni volta che facciamo /next e cerchiamo il successivo
	//cerchiamo nell'array di sendgooglesearchrequest(mess, i) quando i arriva al massimo answer prende valore "finished"
	index := 0

	var search []string
	bot.Start(conf, func(mess string) (string, error) {
		var answer string
		switch mess {
		case "":
			answer = "Welcome to SearchGobot!\nType a term you want to search or /help"
		case "/start":
			answer = "Welcome to SearchGobot!\nType a term you want to search or /help"
		case "/help":
			answer = "It's simple, just type what you want to search using google"
		case "/prev":
			index--
			if index <= 0 {
				index = 0
				answer = search[index] + "/next ➡️\n."
			} else {
				answer = search[index] + "⬅️ /prev ---- /next ➡️\n."
			}

		case "/next":
			index++
			if index < len(search) {
				answer = search[index] + "⬅️ /prev ---- /next ➡️\n."
			} else {
				answer = "no more results! :)\n⬅️ /prev"
			}
		default:
			index = 0
			search = SendGoogleSearchRequest(mess)
			if search == nil {
				answer = "Sorry! Not found :("
			} else {
				answer = search[index] + "/next ➡️\n."
			}
		}
		return answer, nil
	})
}
