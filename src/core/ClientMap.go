package core

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// 提供外部公共的使用
var ClientMap *clientMapStruct

func init() {
	ClientMap = &clientMapStruct{}
}

type clientMapStruct struct {
	// key是客户端的IP Value就是websocket的连接对象
	data sync.Map
	lock sync.Mutex
}

// 存储通信的地址
func (c *clientMapStruct) Store(cli *websocket.Conn) {
	ws := NewWsClient(cli)
	c.data.Store(ws.conn.RemoteAddr().String(), ws)
	// 定时处理连接
	go ws.Ping(time.Second)
	// 处理读循环
	go ws.ReadLoop()
	// 处理总控制循环
	// go ws.HandleLoop()
}

// 移除过期了的客户端连接
func (c *clientMapStruct) Remove(cli *websocket.Conn) {
	c.data.Delete(cli.RemoteAddr().String())
}

// 对所有的客户端发送消息
func (c *clientMapStruct) Sendall(mes string) {
	// 遍历整个map对象拿到key和value(如果报错则移除这个客户端)
	c.data.Range(func(key, value interface{}) bool {
		c.lock.Lock()
		defer c.lock.Unlock()
		con := value.(WsClient).conn
		if err := value.(*WsClient).conn.WriteMessage(websocket.TextMessage, []byte(mes)); err != nil {
			c.Remove(con)
			log.Println(err)
		}
		return true
	})
}
