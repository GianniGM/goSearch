package main

//package esterno da scaricare
import (
	"fmt"
	"github.com/GianniGM/telebot"
	"googleSearch/gSearch"
	"strings"
)

func main() {

	var bot telebot.Bot
	conf := telebot.Configuration{
		BotName: "SearchBot",
		ApiKey:  "123456789:AAAAAAAPIKEYHEREEEEEEE",
	}

	// bot start è una funzione che prende le configurazioni per collegarsi
	// e una funziona (perché in go la funzione è un tipo)
	// ricordarsi che in go le funzioni sono dei tipi
	// comando defer (appena finisce il programma fa quello che gli dici di fare)

	//aggiungi un contatore che si incremente ogni volta che facciamo /next e cerchiamo il successivo
	//cerchiamo nell'array di sendgooglesearchrequest(mess, i) quando i arriva al massimo answer prende valore "finished"

	m := make(map[int64]*gSearch.MySearch)

	bot.Start(conf, func(recMess telebot.TeleMessage) (string, error) {
		var answer string
		var src *gSearch.MySearch
		var ok bool

		// if recMess.From.Uname == "" {
		// 	err := errors.New("Unable to find username")
		// 	return "", err
		// }

		id := recMess.Chat.Chatid
		mess := recMess.Text

		if src, ok = m[id]; !ok {
			//search not exists in the map
			s := gSearch.MySearch{ID: id}
			m[id] = &s
			src = &s
			fmt.Println("!!!NOT FOUND!!!")
		}
		//search exists in the map

		message := strings.SplitAfterN(mess, " ", 2)
		cmd := strings.TrimSpace(message[0])

		switch cmd {

		case "/start":

			answer = "Welcome to SearchGobot!\nType a term you want to search or /help"

		case "/help":

			answer = "It's simple, /search and type what you want to search using google"

		case "/prev":

			if str, first := src.GetPrev(); first != "" {
				answer = first + "\n/next ➡️\n."
			} else {
				answer = str + "\n⬅️ /prev ---- /next ➡️\n."
			}

		case "/next":

			if str, last := src.GetNext(); last != "" {
				answer = last + "⬅️ /prev"
			} else {
				answer = str + "\n⬅️ /prev ---- /next ➡️\n."
			}

		case "/search":
			fmt.Println("Searching", message[1])

			if len(message) > 1 && message[1] != "" {

				//do search of message[1]
				if str, err := src.Search(message[1]); err != nil {
					answer = "Sorry! Not Found :("
				} else {
					answer = "results: " + str + "\n/next ➡️\n."
				}

			} else {
				answer = "please, type\n /search [Term you wanna seaerch]"
			}

		default:
			fmt.Println("Searching", mess)

			if mess == "" {
				answer = "You typed nothing! :/"
			} else if str, err := src.Search(mess); err != nil {
				answer = "Sorry! Not Found :("
			} else {
				answer = "results: " + str + "\n/next ➡️\n."
			}
		}
		return answer, nil
	})
}
