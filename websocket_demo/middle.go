package websocket_demo

import (
	"encoding/binary"
	"fmt"
	"github.com/gorilla/websocket"
)

const (
	TaskRouter = "task_exec"
)

type Task struct {
	ID       string            `json:"id"`
	Content  string            `json:"content"`
	TaskType string            `json:"taskType"`
	Params   map[string]string `json:"params"`
}

type Message struct {
	Version int
	Router  string
	Msg     []byte
	Body    []byte
}

func ParseMsg(b []byte) (*Message, error) {
	if len(b) < 6 {
		return nil, fmt.Errorf("invalid message ")
	}
	version := int(b[0])
	b = b[1:]
	routerLen := int(b[0])
	b = b[1:]
	msgLen := binary.LittleEndian.Uint32(b[:4])
	b = b[4:]
	if int64(routerLen)+int64(msgLen) > int64(len(b)) {
		return nil, fmt.Errorf("invalid message ")
	}
	router := b[:routerLen]
	b = b[routerLen:]
	msg := b[:msgLen]
	b = b[msgLen:]
	body := b
	return &Message{
		Version: version,
		Router:  string(router),
		Msg:     msg,
		Body:    body,
	}, nil
}

// version(1byte) | router name length(1byte) | message length (4byte) | router name(router name length) | message | body
func FormatMsg(router string, msg []byte, body []byte) []byte {
	routerLen := len(router)
	msgLen := uint32(len(msg))
	msgLenBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(msgLenBytes, msgLen)

	res := make([]byte, 0, 6+len(msg)+len(body))
	res = append(res, byte(1), byte(routerLen)) // version | router name len
	res = append(res, msgLenBytes...)           // msgLen
	res = append(res, []byte(router)...)        // router
	if msg != nil {
		res = append(res, msg...) // msg
	}
	if body != nil {
		res = append(res, body...) // body
	}
	return res
}


type WScoket struct{
	conn *websocket.Conn
}
