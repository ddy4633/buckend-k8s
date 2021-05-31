package wscore

import (
	"github.com/gorilla/websocket"
	"log"
)

type WebShellClient struct {
	client *websocket.Conn
}

func NewWebShellClient(client *websocket.Conn) *WebShellClient {
	return &WebShellClient{client: client}
}

// 写websocket数据
func (wc *WebShellClient) Write(data []byte) (n int,err error) {
	// 调用websocket发送数据
	err = wc.client.WriteMessage(websocket.TextMessage,data)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return len(data),err
}

// 读websocket数据
func (wc *WebShellClient) Read(data []byte) (n int,err error)  {
	_,re,err := wc.client.ReadMessage()
	if err !=nil {
		return 0, err
	}
	return copy(data,string(re)),nil
}