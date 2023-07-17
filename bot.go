package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)


type Bot struct{
	BotApi string
	offset int
}

type InterfaceBot interface{
	init(string)
	getUpdates() ([]Update,error)
	respond(update Update) error
	// sendAll(text string) error
	sendOne(text string, ip int) error
}


// инициализация бота
func (b *Bot) init(token string){
	
	b.BotApi = "https://api.telegram.org/bot" + token +"/"
	b.offset = 0
}

// получение обновлений
func (b *Bot) getUpdates() ([]Update,error) {
	// zapr := 

	resp, err := http.Get(b.BotApi + "getUpdates"  + "?offset=" + strconv.Itoa(b.offset))
	if err!=nil{
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil,err
	}

	var restResponse RestResponse

	err = json.Unmarshal(body, &restResponse)
	
	if err!= nil{
		return nil, err
	}

	if len(restResponse.Result)!=0{
		b.offset = restResponse.Result[len(restResponse.Result) - 1].UpdateId + 1
	}

	return restResponse.Result , err
}

// ответ на обновления
func (b *Bot) respond(update Update) error{
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = update.Message.Text

	buf,err := json.Marshal(botMessage)
	if err!=nil{
		return err
	}
	
	_, err = http.Post(b.BotApi + "sendMessage" ,"application/json" , bytes.NewBuffer(buf))
	if err!=nil{
		return err
	}
	return err
}

func (b *Bot) sendOne(text string, ip int) error{
	

	var botMessage BotMessage
	botMessage.ChatId = ip
	botMessage.Text = text
	botMessage.ReplyMarkup.Keyboard = [][]KeyboardButton{{{"button1"},{"button2"}},{{"button3"},{"button4"}}}


	buf, err := json.Marshal(botMessage)
	if err!= nil{
		return err
	}
	
	_, err = http.Post(b.BotApi + "sendMessage", "application/json" , bytes.NewBuffer(buf))
	
	if err!= nil{
		return err
	}

	return err
}