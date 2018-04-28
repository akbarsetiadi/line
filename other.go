package main

import (
	"encoding/json"
	// "errors"
	"log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Response struct {
	Result int `json:"result"`
	ID int `json:"id"`
	Response string `json:"response"`
	Message string `json:"msg"`
	Answer string `json:"answer"`
	Forced bool `json:"force"`
	Image string `json:"image"`
}

func CaraMeme(text string) string{
	if text == "caranya bikin meme roll safe think about it" {
		return "Ketik \"think {teks atas}/{teks bawah}\" . Kalo mau ada yang diilangin teksnya, dikasih \"-\" aja kek gini\n\"think {teks atas}/-\"";
	} else if text == "caranya bikin meme one does not simply" {
		return "Ketik \"odns {teks atas}/{teks bawah}\" . Kalo mau ada yang diilangin teksnya, dikasih \"-\" aja kek gini\n\"odns {teks atas}/-\""
	} else if text == "caranya bikin meme two buttons" {
		return "Ketik \"tombol {teks tombol kiri}/{teks tombol kanan}\" . Kalo mau ada yang diilangin teksnya, dikasih \"-\" aja kek gini\n\"tombol {teks tombol kiri}/-\""
	} else if text == "caranya bikin meme toy story everywhere" {
		return "Ketik \"everywhere {teks atas}/{teks bawah}\" . Kalo mau ada yang diilangin teksnya, dikasih \"-\" aja kek gini\n\"everywhere {teks atas}/-\""
	} else if text == "caranya bikin meme bad luck brian" {
		return "Ketik \"brian {teks atas}/{teks bawah}\". Kalo mau ada yang diilangin teksnya, dikasih \"-\" aja kek gini\n\"brian {teks atas}/-\""
	} else if text == "caranya bikin meme captain picard facepalm" {
		return "Ketik \"facepalm {teks atas}/{teks bawah} \" . Kalo mau ada yang diilangin teksnya, dikasih \"-\" aja kek gini\n\"facepalm {teks atas}/-\""
	} else if text == "caranya bikin meme socially awesome awkward penguin" {
		return "Ketik \"penguin {teks atas}/{teks bawah} \" . Kalo mau ada yang diilangin teksnya, dikasih \"-\" aja kek gini\n\"penguin {teks atas}/-\""
	} else if text == "caranya bikin meme jackie chan wtf" {
		return "Ketik \"wtf {teks atas}/{teks bawah} \" . Kalo mau ada yang diilangin teksnya, dikasih \"-\" aja kek gini \"wtf {teks atas}/-\""
	} else if text == "caranya bikin meme khaleesi" {
		return "Ketik \"ii {teks atas}/{teks bawah} \" . Kalo mau ada yang diilangin teksnya, dikasih \"-\" aja kek gini\n\"ii {teks atas}/-\"\n\n Teksnya ntar bakalan dibikin iii semua"
	} else if text == "caranya bikin meme batman" {
		return "Ketik \"batman {teks kiri}/{teks kanan} \" . Kalo mau ada yang diilangin teksnya, dikasih \"-\" aja kek gini\n\"batman {teks kiri}/-\""
	} else if text == "caranya bikin meme drake hotline bling" {
		return "Ketik \"drake {teks atas}/{teks bawah} \" . Kalo mau ada yang diilangin teksnya, dikasih \"-\" aja kek gini\n\"drake {teks atas}/-\""
	} else if text == "caranya bikin meme expanding brain" {
		return "Ketik \"brain {teks pertama}/{teks kedua}/{teks ketiga}/{teks keempat} \""
	} else if text == "caranya bikin meme distracted boyfriend" {
		return "Ketik \"dbf {teks kiri}/{teks tengah}/{teks kanan}\""
	} else {
		return "ngetik paan sih"
	}
}

func GetSimSimi(text string) string{
	// var Url *url.URL
	// Url, err := url.Parse("http://sandbox.api.simsimi.com")

	// if err != nil {
	// 	log.Println("error")
	// }

	// Url.Path += "/request.p?key=78881f25-f874-4b0a-b04d-d5dc9949c168&lc=id&ft=1.0&text=" + text

	res, err := http.Get("http://sandbox.api.simsimi.com/request.p?key=78881f25-f874-4b0a-b04d-d5dc9949c168&lc=id&ft=1.0&text=" + url.QueryEscape(text))

	if err != nil {
		log.Println("error")
	}

	if res.StatusCode == 200 {
		var body []byte
		body, _ = ioutil.ReadAll(res.Body)

		// if err != nil {
		// 	return err
		// }

		var r Response
		err = json.Unmarshal(body, &r)
		if err != nil {
			log.Println("error")
		}
		if r.Result != 100 {
			return "maaf nih, aku ngga ngerti maksud kk :("
		} else {
			return r.Response
		}
	} else {
		// log.Print("ini statuscodenyaa : ")
		log.Println(res.StatusCode)
		log.Println("hehe")
		return "maaf nih, aku ngga nangkep yg kk bilang :("
	}
	// return err
}

func GetTextFromSimsimi(text string) string{
	if strings.Contains(text,"buatan siapa"){
    	return "Salam kenal yaa! :)\nAku Simsimi buatannya kakak Akbar. Tapi aku lagi versi alpha, masih belom fix bangets"	
    }

    return GetSimSimi(text)
}

func GetYesOrNo() string{
	res, err := http.Get("https://yesno.wtf/api/")

	if err != nil {
		log.Println("error")
	}

	if res.StatusCode == 200 {
		var body []byte
		body, _ = ioutil.ReadAll(res.Body)

		// if err != nil {
		// 	return err
		// }

		var r Response
		err = json.Unmarshal(body, &r)
		if err != nil {
			log.Println("error")
		}
		return r.Image
	} else {
		// log.Print("ini statuscodenyaa : ")
		log.Println(res.StatusCode)
		log.Println("hehe")
		return "bingung :("
	}
}


//DUMP JANGAN DIBUANG

// if strings.HasPrefix(message.Text,"/cek") {
// 	filterSearch := strings.Split(message.Text, "/cek ")[1]
// 	splitFilter := strings.Split(filterSearch,"/")
// 	getDataBL,_ := BLScrape(splitFilter[0],splitFilter[1],splitFilter[2])

// 	var columns []*dbot.CarouselColumn
// 	for i := 0; i < len(getDataBL); i++ {
// 		var titleItem string
// 		if len(getDataBL[i].Title) > 40 {
// 			titleItem = getDataBL[i].Title[0:37] + "..."
// 		} else {
// 			titleItem = getDataBL[i].Title
// 		}
		
// 		column := dbot.NewCarouselColumn(
// 			getDataBL[i].Image, titleItem, getDataBL[i].Price,
// 			dbot.NewURITemplateAction("Go to seller", getDataBL[i].LinkSeller),
// 			dbot.NewURITemplateAction("Go to shop", getDataBL[i].Link),
// 		)	
// 		columns = append(columns, column)
// 	}
// 	template := dbot.NewCarouselTemplate(columns...)
// 	if _, err := bot.ReplyMessage(
// 		event.ReplyToken,
// 		dbot.NewTemplateMessage("nyari " + splitFilter[0], template),
// 	).Do(); err != nil {
// 		log.Print(err)
// 	}
// } else 