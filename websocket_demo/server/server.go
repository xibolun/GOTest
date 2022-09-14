package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	"xibolun/gotest/websocket_demo"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{}

var globalChannel chan []byte

func exec(w http.ResponseWriter, r *http.Request) {
	task := &websocket_demo.Task{
		ID:       time.Now().Format("200601151545"),
		Content:  "uname -a",
		TaskType: "script",
		Params:   nil,
	}

	taskBinary, _ := json.Marshal(task)

	msgBinary := websocket_demo.FormatMsg(websocket_demo.TaskRouter, []byte("hello"), taskBinary)

	globalChannel <- msgBinary

	w.Write([]byte("invoke success"))
}

func conn(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		fmt.Printf("loop\n")
		select {
		case msg := <-globalChannel:
			if err = c.WriteMessage(websocket.BinaryMessage, msg); err != nil {
				log.Println("write:", err)
			}
		}
	}
}

func main() {
	flag.Parse()
	globalChannel = make(chan []byte)
	log.SetFlags(0)
	http.HandleFunc("/conn", conn)
	http.HandleFunc("/task", exec)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
