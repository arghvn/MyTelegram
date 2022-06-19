package main

// We install the package https://github.com/adnanh/webhook (description in readme)

// We use go telegram github to communicate between telegram and coding
//https://github.com/go-telegram-bot-api/telegram-bot-api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(Handler))

// Create a struct that mimics the webhook response body
// https://core.telegram.org/bots/api#update

type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

// This handler is called everytime telegram sends us a webhook event

func Handler(res http.ResponseWriter, req *http.Request) {

	// First, decode the JSON response body

	body := &webhookReqBody{}

	//See readme
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	// Check if the message contains the word "arghvn"
	// if not, return without doing anything
	if !strings.Contains(strings.ToLower(body.Message.Text), "arghvn") {
		return
	}

	// If the text contains arghvn, call the `sayPolo` function, which
	// is defined below
	if err := sayPolo(body.Message.Chat.ID); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	fmt.Println("reply sent")
}

//The below code deals with the process of sending a response message
// to the user

// Create a struct to conform to the JSON body
// of the send message request
// https://core.telegram.org/bots/api#sendmessage

type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// sayPolo takes a chatID and sends "polo" to them
func sayPolo(chatID int64) error {
	// Create the request body struct
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   "Polo!!",
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	// Send a post request with your token
	res, err := http.Post("https://api.telegram.org/bot777845702:AAFdPS_taJ3pTecEFv2jXkmbQfeOqVZGERw/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}

// FInally, the main funtion starts our server on port 3000

}
