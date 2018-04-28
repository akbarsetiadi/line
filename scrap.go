package main

import (
    // import standard libraries
    // "fmt"
    "log"
    "encoding/json"
    "io/ioutil"
    "errors"

    "github.com/PuerkitoBio/goquery"
    "net/url"
    "net/http"
    "strings"
)

type ResponseGrammar struct {
    Result string `json:"result"`
    Errors []ErrorGrammar `json:"errors"`
    Score int `json:"score"`
}

type ErrorGrammar struct {
    Image string `json:"image"`
    Offset int `json:"offset"`
    Length int `json:"length"`
    Bad string `json:"bad"`
    Better []string `json:"better"`
}

type CarouselScrap struct {
    Title string
    Seller string
    Price string
    Link string
    LinkSeller string
    Image string
}

var replacer = strings.NewReplacer("a", "i", "u", "i", "e", "i", "o", "i")

func BLScrape(keyword string, min string, max string) ([]CarouselScrap, string){
    stringURL := "https://www.bukalapak.com/products?utf8=%E2%9C%93&search%5Bkeywords%5D="+ url.QueryEscape(keyword) +"&search%5Bpremium_seller%5D=0&search%5Bnew%5D=1&search%5Bused%5D=1&search%5Bfree_shipping_coverage%5D=&search%5Bprovince%5D=Jabodetabekkar&search%5Bcity%5D=&search%5Bcourier%5D=&search%5Bprice_min%5D="+ min +"&search%5Bprice_max%5D="+ max +"&search%5Brating_gte%5D=0&search%5Brating_lte%5D=5&search%5Btodays_deal%5D=0&search%5Binstallment%5D=0&search%5Bwholesale%5D=0&search%5Btop_seller%5D=0&search%5Bsort_by%5D=price%3Aasc"
    doc, err := goquery.NewDocument(stringURL)
    if err != nil {
        log.Fatal(err)
    }

    var dataShow []CarouselScrap

    doc.Find(".col-12--2 .product-display .product-description").Each(func(index int, itemd *goquery.Selection) {
        if index > 9 && index < 20 {
            var title string
            var seller string
            var price string
            var link string
            var linkSeller string
            var image string
            doc.Find(".col-12--2 .product-display .product-media a picture img").Each(func(inde int, ite *goquery.Selection) {
                if index == inde {
                    images,_ := ite.Attr("data-src")
                    image = images
                }
            })

            doc.Find(".col-12--2 .product-display .product-description h3 a").Each(func(indexs int, item *goquery.Selection) {
                if index == indexs {
                    title = item.Contents().Text()
                    links,_ := item.Attr("href")
                    link = "https://www.bukalapak.com" + links
                }
            })

            doc.Find(".col-12--2 .product-display .product-description .product-seller .user-display-ultra-compact h5 a").Each(func(indexss int, items *goquery.Selection) {
                if index == indexss {
                    seller = items.Contents().Text()
                    links,_ := items.Attr("href")
                    linkSeller = "https://www.bukalapak.com" + links
                }
            })
            doc.Find(".col-12--2 .product-display .product-description .product-price").Each(func(indexsss int, itemss *goquery.Selection) {
                if index == indexsss {
                    prices,_ := itemss.Attr("data-reduced-price")
                    price = "Rp " + prices
                }
            })
            dataShow = append(dataShow, CarouselScrap{Title: title, Seller: seller, Price: price, Link: link, LinkSeller: linkSeller, Image : image})
            // dataShow[index-10] = CarouselScrap{Title: title, Seller: seller, Price: price, Link: link, LinkSeller: linkSeller, Image : image}
        }
    })
    return dataShow,stringURL
}

func CreateMeme(templateId, topText, bottomText string) (string, error) {
    values := url.Values{}
    values.Set("template_id", templateId)
    values.Set("username", "akbarstd")
    values.Set("password", "akbar200697")
    if topText != "-" {
        values.Set("text0", topText)
    }
    if bottomText != "-" {
        values.Set("text1", bottomText)   
    }
    
    resp, err := http.PostForm("https://api.imgflip.com/caption_image", values)

    if err != nil {
        return "", err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var data map[string]interface{}
    err = json.Unmarshal(body, &data)
    if err != nil {
        return "", err
    }

    if !data["success"].(bool) {
        return "", errors.New(data["error_message"].(string))
    }

    url := data["data"].(map[string]interface{})["url"].(string)

    return url, nil
}

func CreateMemeBrain(templateId, text1,text2,text3,text4 string) (string, error) {
    values := url.Values{}
    values.Set("template_id", templateId)
    values.Set("username", "akbarstd")
    values.Set("password", "akbar200697")
    values.Set("boxes[0][text]",text1)
    values.Set("boxes[0][x]","10")
    values.Set("boxes[0][y]","20")
    values.Set("boxes[0][width]","400")
    values.Set("boxes[0][height]","200")
    values.Set("boxes[0][color]","#ffffff")
    values.Set("boxes[0][outline_color]","#000000")
    values.Set("boxes[1][text]",text2)
    values.Set("boxes[1][x]","10")
    values.Set("boxes[1][y]","320")
    values.Set("boxes[1][width]","400")
    values.Set("boxes[1][height]","200")
    values.Set("boxes[1][color]","#ffffff")
    values.Set("boxes[1][outline_color]","#000000")
    values.Set("boxes[2][text]",text3)
    values.Set("boxes[2][x]","10")
    values.Set("boxes[2][y]","620")
    values.Set("boxes[2][width]","400")
    values.Set("boxes[2][height]","200")
    values.Set("boxes[2][color]","#ffffff")
    values.Set("boxes[2][outline_color]","#000000")
    values.Set("boxes[3][text]",text4)
    values.Set("boxes[3][x]","10")
    values.Set("boxes[3][y]","920")
    values.Set("boxes[3][width]","400")
    values.Set("boxes[3][height]","200")
    values.Set("boxes[3][color]","#ffffff")
    values.Set("boxes[3][outline_color]","#000000")
    
    resp, err := http.PostForm("https://api.imgflip.com/caption_image", values)

    if err != nil {
        return "", err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var data map[string]interface{}
    err = json.Unmarshal(body, &data)
    if err != nil {
        return "", err
    }

    if !data["success"].(bool) {
        return "", errors.New(data["error_message"].(string))
    }

    url := data["data"].(map[string]interface{})["url"].(string)

    return url, nil
}

func CreateMemeDB(templateId, text1,text2,text3 string) (string, error) {
    values := url.Values{}
    values.Set("template_id", templateId)
    values.Set("username", "akbarstd")
    values.Set("password", "akbar200697")
    values.Set("boxes[0][text]",text1)
    values.Set("boxes[0][x]","130")
    values.Set("boxes[0][y]","400")
    values.Set("boxes[0][width]","400")
    values.Set("boxes[0][height]","200")
    values.Set("boxes[0][color]","#ffffff")
    values.Set("boxes[0][outline_color]","#000000")
    values.Set("boxes[1][text]",text2)
    values.Set("boxes[1][x]","580")
    values.Set("boxes[1][y]","300")
    values.Set("boxes[1][width]","300")
    values.Set("boxes[1][height]","100")
    values.Set("boxes[1][color]","#ffffff")
    values.Set("boxes[1][outline_color]","#000000")
    values.Set("boxes[2][text]",text3)
    values.Set("boxes[2][x]","850")
    values.Set("boxes[2][y]","440")
    values.Set("boxes[2][width]","300")
    values.Set("boxes[2][height]","100")
    values.Set("boxes[2][color]","#ffffff")
    values.Set("boxes[2][outline_color]","#000000")
    
    resp, err := http.PostForm("https://api.imgflip.com/caption_image", values)

    if err != nil {
        return "", err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var data map[string]interface{}
    err = json.Unmarshal(body, &data)
    if err != nil {
        return "", err
    }

    if !data["success"].(bool) {
        return "", errors.New(data["error_message"].(string))
    }

    url := data["data"].(map[string]interface{})["url"].(string)

    return url, nil
}

func GetMeme(text,userID,groupID string) string{
    command := strings.Split(text, " ")
    if command[0] == "think" {
        command = strings.Split(text, "think ")
        tulisan := strings.Split(command[1],"/")
        memeURL,err := CreateMeme("89370399",tulisan[0],tulisan[1])
        if err != nil {
            return "error"
        } else {
            InsertData(userID, "1",memeURL,groupID)
            return memeURL           
        }
    } else if command[0] == "odns" {
        command = strings.Split(text, "odns ")
        tulisan := strings.Split(command[1],"/")
        memeURL,err := CreateMeme("61579",tulisan[0],tulisan[1])
        if err != nil {
            return "error"
        } else {
            InsertData(userID, "2",memeURL,groupID)
            return memeURL           
        }
    } else if command[0] == "tombol" {
        command = strings.Split(text, "tombol ")
        tulisan := strings.Split(command[1],"/")
        memeURL,err := CreateMeme("87743020",tulisan[0],tulisan[1])  
        if err != nil {
            return "error"
        } else {
            InsertData(userID, "4",memeURL,groupID)
            return memeURL           
        }      
    }else if command[0] == "everywhere" {
        command = strings.Split(text, "everywhere ")
        tulisan := strings.Split(command[1],"/")
        memeURL,err := CreateMeme("347390",tulisan[0],tulisan[1])
        if err != nil {
            return "error"
        } else {
            InsertData(userID, "5",memeURL,groupID)
            return memeURL           
        }
    }else if command[0] == "brian" {
        command = strings.Split(text, "brian ")
        tulisan := strings.Split(command[1],"/")
        memeURL,err := CreateMeme("61585",tulisan[0],tulisan[1])
        if err != nil {
            return "error"
        } else {
            InsertData(userID, "6",memeURL,groupID)
            return memeURL           
        }
    }else if command[0] == "facepalm" {
        command = strings.Split(text, "facepalm ")
        tulisan := strings.Split(command[1],"/")
        memeURL,err := CreateMeme("1509839",tulisan[0],tulisan[1])
        if err != nil {
            return "error"
        } else {
            InsertData(userID, "7",memeURL,groupID)
            return memeURL           
        }
    }else if command[0] == "penguin" {
        command = strings.Split(text, "penguin ")
        tulisan := strings.Split(command[1],"/")
        memeURL,err := CreateMeme("61584",tulisan[0],tulisan[1])
        if err != nil {
            return "error"
        } else {
            InsertData(userID, "8",memeURL,groupID)
            return memeURL           
        }
    }else if command[0] == "wtf" {
        command = strings.Split(text, "wtf ")
        tulisan := strings.Split(command[1],"/")
        memeURL,err := CreateMeme("412211",tulisan[0],tulisan[1])
        if err != nil {
            return "error"
        } else {
            InsertData(userID, "9",memeURL,groupID)
            return memeURL           
        }
    }else if command[0] == "ii" {
        command = strings.Split(text, "ii ")
        tulisan := strings.Split(command[1],"/")
        tulisanAtas := replacer.Replace(tulisan[0])
        tulisanBawah := replacer.Replace(tulisan[1])
        log.Println(tulisanAtas)
        log.Println(tulisanBawah)
        memeURL,err := CreateMeme("47877343",tulisanAtas,tulisanBawah)
        if err != nil {
            log.Println(err)
            return "error"
        } else {
            InsertData(userID, "3",memeURL,groupID)
            return memeURL           
        }
    }else if command[0] == "batman" {
        command = strings.Split(text, "batman ")
        tulisan := strings.Split(command[1],"/")
        memeURL,err := CreateMeme("438680",tulisan[0],tulisan[1])
        if err != nil {
            return "error"
        } else {
            InsertData(userID, "10",memeURL,groupID)
            return memeURL           
        }
    } else if command[0] == "drake" {
        command = strings.Split(text, "drake ")
        tulisan := strings.Split(command[1],"/")
        memeURL,err := CreateMeme("71491297",tulisan[0],tulisan[1])
        if err != nil {
            return "error"
        } else {
            InsertData(userID, "11",memeURL,groupID)
            return memeURL           
        }
    } else if command[0] == "brain" {
        text = strings.ToUpper(text)
        command = strings.Split(text, "BRAIN ")
        tulisan := strings.Split(command[1],"/")
        if len(tulisan) == 4 {
            memeURL,err := CreateMemeBrain("93895088",tulisan[0],tulisan[1],tulisan[2],tulisan[3])
            if err != nil {
                return "error"
            } else {
                InsertData(userID, "12",memeURL,groupID)
                return memeURL           
            }
        } else {
            return "bukan semuanya"
        }
    } else if command[0] == "dbf" {
        text = strings.ToUpper(text)
        command = strings.Split(text, "DBF ")
        tulisan := strings.Split(command[1],"/")
        if len(tulisan) == 3 {
            memeURL,err := CreateMemeDB("112126428",tulisan[0],tulisan[1],tulisan[2])
            if err != nil {
                return "error"
            } else {
                InsertData(userID, "13",memeURL,groupID)
                return memeURL           
            }
        } else {
            return "bukan semuanya"
        }
    } else {
        return "bukan semuanya"
    }
}

func InsertData(userID, meme,urlImg,groupID string) {
    values := url.Values{}
    values.Set("user_id", userID)
    values.Set("meme", meme)
    values.Set("url_img", urlImg)
    values.Set("group_id", groupID)
    
    resp, err := http.PostForm("https://kakekdenik.com/portal-line.php", values)

    if err != nil {
       log.Println(err)
    } else {
        log.Println(resp)
    }
}





// func TokenScrape() string{
//     doc, err := goquery.NewDocument("https://textgears.com/api/signup.php?givemethatgoddamnkey=please&plan=10")
//     if err != nil {
//         log.Fatal(err)
//     }

//     var token string

//     // use CSS selector found with the browser inspector
//     // for each, use index and item

//     doc.Find("span").Each(func(index int, item *goquery.Selection) {
//         if( item.AttrOr("class","") == "green") {
//             token = item.Contents().Text()
//         }
//     })
//     return token
// }

// func CheckGrammar(text string) []string{
//     // textUrl := url.QueryEscape(text)
//     token := TokenScrape()
//     res, err := http.Get("https://api.textgears.com/check.php?text="+ textUrl)

//     if err != nil {
//         log.Println("error")
//     }

//     if res.StatusCode == 200 {
//         var body []byte
//         body, _ = ioutil.ReadAll(res.Body)

//         // if err != nil {
//         //  return err
//         // }

//         var r ResponseGrammar
//         err = json.Unmarshal(body, &r)
//         if err != nil {
//             log.Println("error")
//         }
//         if r.Result == false {

//             return "maaf nih, aku ngga ngerti maksud kk :("
//         } else if r.Score < 100 {
            
//         }
//     } else {
//         // log.Print("ini statuscodenyaa : ")
//         log.Println(res.StatusCode)
//         log.Println("hehe")
//         return "maaf nih, aku ngga nangkep yg kk bilang :("
//     }
// }