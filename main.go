package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)


type Consumer struct {
	Name string
	Host string
}

type Message struct {
	action string
	data   string
	queue  string

}



func send(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Сообщение отправлено"))
}

func recieve(w http.ResponseWriter, r *http.Request) {

}

func register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	var message Message
	log.Println(body)
	err = json.Unmarshal(body, &message)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	log.Println(message)
}


func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", register)
	mux.HandleFunc("/send", send)
	mux.HandleFunc("/recieve", recieve)
	log.Println("Запуск веб-сервера на http://localhost:7777")
	err := http.ListenAndServe(":7777", mux)
    log.Fatal(err)
}