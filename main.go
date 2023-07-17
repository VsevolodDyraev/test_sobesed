package main

import (
	"fmt"
	"log"
)

func main() {

	var bot_int InterfaceBot
	bot := Bot{}

	bot_int = &bot

	bot_int.init("2010595022:AAG1xMs2CgrcIRSkBJZHvsCgLp3ZnFkkHko")

	for {
		updates, err := bot.getUpdates()
		if err != nil {
			log.Fatal("Somefing wrong in get updates")
		}

		if len(updates) != 0 {
			fmt.Println(updates)
			for _, u := range updates {
				// bot.respond(u)
				if err != nil {
					log.Println(err)
					continue
				}
				bot.sendOne(u.Message.Text,u.Message.Chat.ChatId)
			}
		}

	}

}
