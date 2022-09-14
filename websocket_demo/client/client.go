package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"xibolun/gotest/websocket_demo"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func ExecTask(msgBody []byte) (stdout, stderr string, err error) {
	var task websocket_demo.Task
	_ = json.Unmarshal(msgBody, &task)

	cmd := exec.Command("uname", "-a")
	output, err := cmd.CombinedOutput()
	stdout = string(output)
	log.Printf("exec cmd: %s success", task.Content)
	log.Printf(stdout)
	return
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/conn"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	// received message from server
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			continue
		}

		msg, err := websocket_demo.ParseMsg(message)
		if err != nil {
			log.Println("read:", err)
			continue
		}

		switch msg.Router {
		case websocket_demo.TaskRouter:
			_, _, _ = ExecTask(msg.Body)

		}
	}
}
