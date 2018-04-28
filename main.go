// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	dbot "github.com/line/line-bot-sdk-go/linebot"
)

var bot *dbot.Client

func main() {
	var err error
	bot, err = dbot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	// bot, err = dbot.New("31bd88471f3bd38c67dc5543fae67550", "0pOIgcRBERYWmJ5YnPrHZwp5TiG5ytxlYrdJSZYaU1x4kRl0YhNV32PAkSmksA8MpweBVSbFkrdlj8jTTw88E/MDQKPJMWHn8i0nmxmYZ+/VV8B4uHfpQctIyIWM+C5FwKCZlxoTCsMAMPl10Pfz+AdB04t89/1O/w1cDnyilFU=")
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	// port := "1234"
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == dbot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == dbot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *dbot.TextMessage:
				if event.Source.GroupID == "" {
					message.Text = strings.ToLower(message.Text)
					if message.Text == "bikin meme" {
						templatess := dbot.NewCarouselTemplate(
							dbot.NewCarouselColumn(
								"https://image.ibb.co/e3OaeH/roll_think.jpg", "Roll Safe Think About It & Drake", "Klik salah satu buat tau cara bikinnya",
								dbot.NewMessageTemplateAction("Think", "caranya bikin meme roll safe think about it"),
								dbot.NewMessageTemplateAction("Drake", "caranya bikin meme drake hotline bling"),
							),
							dbot.NewCarouselColumn(
								"https://image.ibb.co/hSE3Yc/odns.jpg", "One Doesnt Simply & Expanding Brain", "Klik salah satu buat tau cara bikinnya",
								dbot.NewMessageTemplateAction("ODNS", "caranya bikin meme one does not simply"),
								dbot.NewMessageTemplateAction("Brain", "caranya bikin meme expanding brain"),
							),
							dbot.NewCarouselColumn(
								"https://image.ibb.co/gGvZ8c/khaleesi.png", "Khaleesi & Distracted Boyfriend", "Klik salah satu buat tau cara bikinnya",
								dbot.NewMessageTemplateAction("Khaleesi", "caranya bikin meme khaleesi"),
								dbot.NewMessageTemplateAction("Distracted Boyfriend", "caranya bikin meme distracted boyfriend"),
							),
							dbot.NewCarouselColumn(
								"https://imgflip.com/s/meme/Two-Buttons.jpg", "Two buttons", "Two buttons",
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme two buttons"),
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme two buttons"),
							),
							dbot.NewCarouselColumn(
								"https://imgflip.com/s/meme/X-Everywhere.jpg", "Everywhere", "Everywhere",
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme toy story everywhere"),
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme toy story everywhere"),
							),
							dbot.NewCarouselColumn(
								"https://imgflip.com/s/meme/Bad-Luck-Brian.jpg", "Bad Luck Brian", "Bad Luck Brian",
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme bad luck brian"),
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme bad luck brian"),
							),
							dbot.NewCarouselColumn(
								"https://imgflip.com/s/meme/Captain-Picard-Facepalm.jpg", "Captain Picard Facepalm", "Facepalm",
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme captain picard facepalm"),
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme captain picard facepalm"),
							),
							dbot.NewCarouselColumn(
								"https://imgflip.com/s/meme/Socially-Awesome-Awkward-Penguin.jpg", "Penguin", "Penguin",
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme socially awesome awkward penguin"),
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme socially awesome awkward penguin"),
							),
							dbot.NewCarouselColumn(
								"https://imgflip.com/s/meme/Jackie-Chan-WTF.jpg", "Jackie Chan WTF", "Jackie Chan WTF",
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme jackie chan wtf"),
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme jackie chan wtf"),
							),
							dbot.NewCarouselColumn(
								"https://imgflip.com/s/meme/Batman-Slapping-Robin.jpg", "Batman", "Batman slapping robin",
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme batman"),
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme batman"),
							),
						)
						if _, err := bot.ReplyMessage(
							event.ReplyToken,
							dbot.NewTemplateMessage("milih template meme", templatess),
						).Do(); err != nil {
							log.Print(err)
						}

					} else if message.Text == "bikin meme kedua"{
						if _, err = bot.ReplyMessage(event.ReplyToken, dbot.NewTextMessage("hehe")).Do(); err != nil {
							log.Print(err)
						}
					} else if strings.HasPrefix(message.Text,"caranya") {
						text := CaraMeme(message.Text)
						if _, err = bot.ReplyMessage(event.ReplyToken, dbot.NewTextMessage(text)).Do(); err != nil {
							log.Print(err)
						}
					} else if message.Text == "cari barang online" {
						text := "hehe"
						if _, err = bot.ReplyMessage(event.ReplyToken, dbot.NewTextMessage(text)).Do(); err != nil {
							log.Print(err)
						}
					} else if message.Text == "maen" {
						if _, err = bot.ReplyMessage(event.ReplyToken, dbot.NewTextMessage("hehe")).Do(); err != nil {
							log.Print(err)
						}
					} else if message.Text == "info bot" {
						text := "hehe"
						if _, err = bot.ReplyMessage(event.ReplyToken, dbot.NewTextMessage(text)).Do(); err != nil {
							log.Print(err)
						}
					} else if message.Text == "kick just for fun" {
						if _, err = bot.ReplyMessage(event.ReplyToken, dbot.NewTextMessage("dadahh :*")).Do(); err != nil {
							log.Print(err)
						}
						if _, err = bot.LeaveRoom(event.Source.RoomID).Do(); err != nil {
							log.Print(err)
						}
					} else {
						imageUrl := GetMeme(message.Text,event.Source.UserID,event.Source.GroupID)
						if imageUrl == "bukan semuanya" {
							// if _, err = bot.ReplyMessage(event.ReplyToken, dbot.NewTextMessage("ngetik paan sih")).Do(); err != nil {
							// 	log.Print(err)
							// }	
						} else if imageUrl == "error"{
							if _, err = bot.ReplyMessage(event.ReplyToken, dbot.NewTextMessage("error mon maap, coba diulangin lagi hehe")).Do(); err != nil {
								log.Print(err)
							}
						} else {
							imageUrl = strings.Replace(imageUrl, "http://", "https://", -1)
							if _, err = bot.ReplyMessage(event.ReplyToken, dbot.NewImageMessage(imageUrl,imageUrl)).Do(); err != nil {
								log.Print(err)
							}							
						}
					}
						
				} else {
					message.Text = strings.ToLower(message.Text)
					if strings.HasPrefix(message.Text,"/cek") {
						filterSearch := strings.Split(message.Text, "/cek ")[1]
						splitFilter := strings.Split(filterSearch,"/")
						getDataBL,_ := BLScrape(splitFilter[0],splitFilter[1],splitFilter[2])

						var columns []*dbot.CarouselColumn
						for i := 0; i < len(getDataBL); i++ {
							var titleItem string
							if len(getDataBL[i].Title) > 40 {
								titleItem = getDataBL[i].Title[0:37] + "..."
							} else {
								titleItem = getDataBL[i].Title
							}
							
							column := dbot.NewCarouselColumn(
								getDataBL[i].Image, titleItem, getDataBL[i].Price,
								dbot.NewURITemplateAction("Go to seller", getDataBL[i].LinkSeller),
								dbot.NewURITemplateAction("Go to shop", getDataBL[i].Link),
							)	
							columns = append(columns, column)
						}
						template := dbot.NewCarouselTemplate(columns...)
						if _, err := bot.ReplyMessage(
							event.ReplyToken,
							dbot.NewTemplateMessage("nyari " + splitFilter[0], template),
						).Do(); err != nil {
							log.Print(err)
						}
					} else if message.Text == "bikin meme" {
						templatess := dbot.NewCarouselTemplate(
							dbot.NewCarouselColumn(
								"https://image.ibb.co/e3OaeH/roll_think.jpg", "Roll Safe Think About It & Drake", "Klik salah satu buat tau cara bikinnya",
								dbot.NewMessageTemplateAction("Think", "caranya bikin meme roll safe think about it"),
								dbot.NewMessageTemplateAction("Drake", "caranya bikin meme drake hotline bling"),
							),
							dbot.NewCarouselColumn(
								"https://image.ibb.co/hSE3Yc/odns.jpg", "One Doesnt Simply & Expanding Brain", "Klik salah satu buat tau cara bikinnya",
								dbot.NewMessageTemplateAction("ODNS", "caranya bikin meme one does not simply"),
								dbot.NewMessageTemplateAction("Brain", "caranya bikin meme expanding brain"),
							),
							dbot.NewCarouselColumn(
								"https://image.ibb.co/gGvZ8c/khaleesi.png", "Khaleesi & Distracted Boyfriend", "Klik salah satu buat tau cara bikinnya",
								dbot.NewMessageTemplateAction("Khaleesi", "caranya bikin meme khaleesi"),
								dbot.NewMessageTemplateAction("Distracted Boyfriend", "caranya bikin meme distracted boyfriend"),
							),
							dbot.NewCarouselColumn(
								"https://imgflip.com/s/meme/Two-Buttons.jpg", "Two buttons", "Two buttons",
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme two buttons"),
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme two buttons"),
							),
							dbot.NewCarouselColumn(
								"https://imgflip.com/s/meme/X-Everywhere.jpg", "Everywhere", "Everywhere",
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme toy story everywhere"),
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme toy story everywhere"),
							),
							dbot.NewCarouselColumn(
								"https://imgflip.com/s/meme/Bad-Luck-Brian.jpg", "Bad Luck Brian", "Bad Luck Brian",
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme bad luck brian"),
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme bad luck brian"),
							),
							dbot.NewCarouselColumn(
								"https://imgflip.com/s/meme/Captain-Picard-Facepalm.jpg", "Captain Picard Facepalm", "Facepalm",
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme captain picard facepalm"),
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme captain picard facepalm"),
							),
							dbot.NewCarouselColumn(
								"https://imgflip.com/s/meme/Socially-Awesome-Awkward-Penguin.jpg", "Penguin", "Penguin",
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme socially awesome awkward penguin"),
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme socially awesome awkward penguin"),
							),
							dbot.NewCarouselColumn(
								"https://imgflip.com/s/meme/Jackie-Chan-WTF.jpg", "Jackie Chan WTF", "Jackie Chan WTF",
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme jackie chan wtf"),
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme jackie chan wtf"),
							),
							dbot.NewCarouselColumn(
								"https://imgflip.com/s/meme/Batman-Slapping-Robin.jpg", "Batman", "Batman slapping robin",
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme batman"),
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme batman"),
							),
						)
						if _, err := bot.ReplyMessage(
							event.ReplyToken,
							dbot.NewTemplateMessage("milih template meme", templatess),
						).Do(); err != nil {
							log.Print(err)
						}

					} else if message.Text == "bikin meme kedua"{
						
						templatess := dbot.NewCarouselTemplate(
							dbot.NewCarouselColumn(
								"https://i.imgflip.com/t7px7.jpg", "Drake", "Drake hotline bling",
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme drake hotline bling"),
							),
							dbot.NewCarouselColumn(
								"https://imgflip.com/s/meme/Expanding-Brain.jpg", "expanding brain", "expanding brain",
								dbot.NewMessageTemplateAction("Caranya", "caranya bikin meme expanding brain"),
							),
						)
						if _, err := bot.ReplyMessage(
							event.ReplyToken,
							dbot.NewTemplateMessage("milih template memes", templatess),
						).Do(); err != nil {
							log.Print(err)
						}
					} else if strings.HasPrefix(message.Text,"caranya") {
						text := CaraMeme(message.Text)
						if _, err = bot.ReplyMessage(event.ReplyToken, dbot.NewTextMessage(text)).Do(); err != nil {
							log.Print(err)
						} 
					} else if message.Text == "kick just for fun" {
						if _, err = bot.ReplyMessage(event.ReplyToken, dbot.NewTextMessage("dadahh :*")).Do(); err != nil {
							log.Print(err)
						}
						if _, err = bot.LeaveGroup(event.Source.GroupID).Do(); err != nil {
							log.Print(err)
						}
					}else {
						imageUrl := GetMeme(message.Text,event.Source.UserID,event.Source.GroupID)
						if imageUrl == "bukan semuanya" {
							// if _, err = bot.ReplyMessage(event.ReplyToken, dbot.NewTextMessage("ngetik paan sih")).Do(); err != nil {
							// 	log.Print(err)
							// }	
						} else if imageUrl == "error"{
							if _, err = bot.ReplyMessage(event.ReplyToken, dbot.NewTextMessage("error mon maap, coba diulangin lagi hehe")).Do(); err != nil {
								log.Print(err)
							}
						} else {
							imageUrl = strings.Replace(imageUrl, "http://", "https://", -1)
							// log.Println(imageUrl)
							if _, err = bot.ReplyMessage(event.ReplyToken, dbot.NewImageMessage(imageUrl,imageUrl)).Do(); err != nil {
								log.Print(err)
							}							
						}
					} 
					
				}
			// default:
				// if _, err = bot.ReplyMessage(event.ReplyToken, dbot.NewTextMessage("maaf nih, aku cuman ngerti kalo kk ngomongnya make teks, bukan gambar,video dll")).Do(); err != nil {
				// 	log.Print(err)
				// }
			}
		} else if event.Type == dbot.EventTypeJoin {
			text := "halo hehe\n\nklo mo kick aq ngga perlu repot2, ketik ini ajah \"kick just for fun\"\nhave fun!"
			if _, err = bot.ReplyMessage(event.ReplyToken, dbot.NewTextMessage(text)).Do(); err != nil {
				log.Print(err)
			}
		}
	}
}