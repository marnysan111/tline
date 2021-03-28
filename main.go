package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {

	//	Secret := "60d4379d20eca6b2a701c111e5587258"
	//	Token := "7/f3ibMIbfOC8lzJSe42qcDiw6Hiwy2sYNVvppumQZ8k2bBzYWgEl8fz6nmlRC4n5JBC9oMIyMXk1UfgNtL3PM0xUvGlPmNkQuNXSF0/3z6ttrCmxTlpop4W7nZACBYI4dmPH9Pz339i3/7hGUAnfgdB04t89/1O/w1cDnyilFU="
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	Secret := os.Getenv("Secret")
	Token := os.Getenv("Token")
	bot, err := linebot.New(Secret, Token)
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
