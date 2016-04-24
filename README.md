# goSearch 

A telegram bot written in go using [Telebot](https://github.com/cortinico/telebot) skeleton.

It does web Searches on Google for you.

## Usage

When deployed on a server just type your terms to search on Telegram Client and bot will reply you 
a list of suggested links.

## Configuration

1) download [telebot](https://github.com/cortinico/telebot)

2) create new bot using [botFather](https://telegram.me/BotFather)

3) add your API key and botname

```go

	conf := telebot.Configuration{
		BotName: "YourBotName_bot",
		ApiKey:  "162227600:!!!YOURAPIKEY!!!!BBBBCCCCCCCCCDDDDD"}

```

4) build, deploy and have fun! 

## Licence

The following software is released under the [GPL3](https://github.com/GianniGM/goSearch/blob/master/LICENSE)