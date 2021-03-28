package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {

	ChanellSecret := "c908a2772b8888770864cbfebe7fb19b"
	ChanellToken := "fvqbQTSR5zCUB2CFDaZK7bP6F4g77z30QBkn+3UArzE1NBaiXmXpINQKmTGWz463KfR4pUQWnvq/Q4eQTSdA3yUrymYgTxv6HvjaiI+Kc63lqQEToT8DovnKmlwfpDB3brXndxbCUEypmz8M/W50FQdB04t89/1O/w1cDnyilFU="
	bot, err := linebot.New(ChanellSecret, ChanellToken)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if message.Text == "Twitter" {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("twitterだお")).Do(); err != nil {
							fmt.Println(message.Text)
							log.Print(err)
						}
					} else {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("これはテストです")).Do(); err != nil {
							log.Print(err)
						}
					}
				case *linebot.StickerMessage:
					replyMessage := fmt.Sprintf(
						"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
