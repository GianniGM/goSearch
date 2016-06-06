package main

//package esterno da scaricare
import (
	"errors"
	"fmt"
	"github.com/GianniGM/telebot"
	"googleSearch/user"
	"strings"
)

func main() {

	var bot telebot.Bot
	conf := telebot.Configuration{
		BotName: "SearchGobot",
		ApiKey:  "204011400:AAGasqp4x-XIupLiXSYGCLy7e0T7vVRAsZg",
	}

	// bot start è una funzione che prende le configurazioni per collegarsi
	// e una funziona (perché in go la funzione è un tipo)
	// ricordarsi che in go le funzioni sono dei tipi
	// comando defer (appena finisce il programma fa quello che gli dici di fare)

	//aggiungi un contatore che si incremente ogni volta che facciamo /next e cerchiamo il successivo
	//cerchiamo nell'array di sendgooglesearchrequest(mess, i) quando i arriva al massimo answer prende valore "finished"

	m := make(map[string]*user.User)

	bot.Start(conf, func(recMess telebot.TeleMessage) (string, error) {
		var answer string
		var usr *user.User
		var ok bool

		if recMess.From.Uname == "" {
			err := errors.New("Unable to find username")
			return "", err
		}

		userName := "@" + recMess.From.Uname
		mess := recMess.Text

		if usr, ok = m[userName]; !ok {
			//user not exists in the map
			u := user.User{Name: userName}
			m[userName] = &u
			usr = &u
			fmt.Println("!!!NOT FOUND!!!")
		}
		//user exists in the map

		message := strings.SplitAfterN(mess, " ", 2)
		cmd := strings.TrimSpace(message[0])

		switch cmd {

		case "/start":

			answer = "Welcome to SearchGobot!\nType a term you want to search or /help"

		case "/help":

			answer = "It's simple, /search and type what you want to search using google"

		case "/prev":

			if str, first := usr.GetPrev(); first != "" {
				answer = first + "\n/next ➡️\n."
			} else {
				answer = str + "\n⬅️ /prev ---- /next ➡️\n."
			}

		case "/next":

			if str, last := usr.GetNext(); last != "" {
				answer = last + "⬅️ /prev"
			} else {
				answer = str + "\n⬅️ /prev ---- /next ➡️\n."
			}

		case "/search":
			fmt.Println("Searching", message[1])

			if len(message) > 1 && message[1] != "" {

				//do search of message[1]
				if str, err := usr.Search(message[1]); err != nil {
					answer = "Sorry! Not Found :("
				} else {
					answer = usr.GetName() + " searched: " + str + "\n/next ➡️\n."
				}

			} else {
				answer = usr.GetName() + "please, type\n /search [Term you wanna seaerch]"
			}

		default:
			fmt.Println("Searching", mess)

			if mess == "" {
				answer = "You typed nothing! :/"
			} else if str, err := usr.Search(mess); err != nil {
				answer = "Sorry! Not Found :("
			} else {
				answer = usr.GetName() + " searched: " + str + "\n/next ➡️\n."
			}
		}
		return answer, nil
	})
}
