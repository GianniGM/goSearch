# GoSearch 

A telegram bot written in go using [Telebot](https://github.com/cortinico/telebot) skeleton.
It does web Searches on Google for you.

## Usage

When deployed on a server just type your terms to search on Telegram Client and bot will reply you 
a list of suggested links.

## Custom Search configuration

1) Create your Google *Custom Search Engine* [here](https://cse.google.com/cse/all)

2) Get your *Search Engine ID*

3) Activate your *Custom Search API* [here](https://console.cloud.google.com/home/dashboard)

4) Paste in *search.go* your search ID and your API

```go
	
	SearchEngineID := "PASTE-YOUR-CUSTOM-SEARCH-ID"
	APIgSearchKey := "PASTE-YOUR-GOOGLECUSTOM-SEARCH-API"

```

## Bot configuration

1) download my [telebot](https://github.com/GianniGM/telebot) fork

2) create new bot using [botFather](https://telegram.me/BotFather)

3) add your API key and botname

```go

	conf := telebot.Configuration{
		BotName: "YourBotName_bot",
		ApiKey:  "162227600:!!!YOURAPIKEY!!!!BBBBCCCCCCCCCDDDDD"
	}

```

4) build, deploy and have fun! 

## Licence

The following software is released under the [GPL3 License](https://github.com/GianniGM/goSearch/blob/master/LICENSE)