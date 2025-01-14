package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
 	"database/sql"
    _ "modernc.org/sqlite"
)


// TextMessage структура JSON полученных текстовых сообщений
type TextMessage struct {
    Ok     bool `json:"ok"`
    Result []struct {
        UpdateID int `json:"update_id"`
        Message  struct {
            MessageID int `json:"message_id"`
            From      struct {
                ID           int    `json:"id"`
                IsBot        bool   `json:"is_bot"`
                FirstName    string `json:"first_name"`
                LastName     string `json:"last_name"`
                LanguageCode string `json:"language_code"`
            } `json:"from"`
            Chat struct {
                ID        int    `json:"id"`
                FirstName string `json:"first_name"`
                LastName  string `json:"last_name"`
                Type      string `json:"type"`
            } `json:"chat"`
            Date int    `json:"date"`
            Text string `json:"text"`
        } `json:"message"`
    } `json:"result"`
}

func main() {

    url := "https://api.telegram.org/bot6971995963:AAH8uBq41VNlrm0BtKmk54s4_mN5ZcaksG0/getUpdates"

    getText, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }

    defer getText.Body.Close()
    text, err := ioutil.ReadAll(getText.Body)
    if err != nil {
        log.Fatal(err)
    }

    bodyGet := []byte(text)

    var app = TextMessage{}
    err1 := json.Unmarshal(bodyGet, &app)
    if err1 != nil {
        log.Fatal("error")
    }

    

  db, err := sql.Open("sqlite", "telegram.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

for _, row := range app.Result {   
    result, err := db.Exec("insert into telegram (text, chat_id) values ($1, $2);",
        row.Message.Text,
        row.Message.Chat.ID)
    fmt.Println(result.LastInsertId())  // id последнего добавленного объекта
    fmt.Println(result.RowsAffected())  // количество добавленных строк
  if err != nil{
        panic(err)
    }



for _, row := range app.Result {  
        fmt.Printf("UpdateID:%d  ChatID:%d  Text:%s\n", row.UpdateID, row.Message.Chat.ID, row.Message.Text)
    }


}
}