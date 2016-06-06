package main

//package esterno da scaricare
import (
	"errors"
	"github.com/GianniGM/telebotgae"
	"googleSearch/user"
	"strings"
)

func main() {

	var bot telebotgae.Bot
	conf := telebotgae.Configuration{
		BotName: "EchoBot",
		ApiKey:  "123456789:AAAAAAAPIKEYHEREEEEEEE",
	}


	// bot start è una funzione che prende le configurazioni per collegarsi
	// e una funziona (perché in go la funzione è un tipo)
	// ricordarsi che in go le funzioni sono dei tipi
	// comando defer (appena finisce il programma fa quello che gli dici di fare)

	//aggiungi un contatore che si incremente ogni volta che facciamo /next e cerchiamo il successivo
	//cerchiamo nell'array di sendgooglesearchrequest(mess, i) quando i arriva al massimo answer prende valore "finished"

	// questa potrebbe diventare una struttura molto molto grossa
	// trovare un modo per pulire ciò che non serve
	m := make(map[string]user.User)

	bot.Startgae(conf, func(recMess telebot.TeleMessager, *http.Request) (string, error) {
		var answer string
		var usr user.User
		var ok bool

		if recMess.From.Uname == "" {
			err := errors.New("Unable to find username")
			return "", err
		}

		userName := "@" + recMess.From.Uname
		mess := recMess.Text

		if usr, ok = m[userName]; !ok {
			//user not exists in the map
			// create one 
			usr = user.User{Name: userName}
			m[userName] = usr
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
			answer = usr.GetPrev()
		case "/next":
			answer = usr.GetNext()
		case "/search":
			if len(message) > 1 && message[1] != "" {
				//do search of message[1]
				answer = usr.Search(message[1])
			} else {
				answer = usr.GetName() + "please, type\n /search [Term you wanna seaerch]"
			}

		default:
			if mess != "" {
				//do search of mess
				answer = usr.Search(mess)
			} else {
				answer = "you typed nothing!"
			}
		}
		return answer, nil
	})
}
