package wscore

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type WsClient struct {
	conn      *websocket.Conn // 连接队列
	readChan  chan *WsMessage // 读队列
	closeChan chan int        // 关闭队列
	locl      sync.Mutex      // 加锁
}

func NewWsClient(conn *websocket.Conn) *WsClient {
	return &WsClient{conn: conn, readChan: make(chan *WsMessage), closeChan: make(chan int, 1)}
}

// 心跳检查，如果出错则自动移除
func (w *WsClient) Ping(wait time.Duration) {
	for {
		time.Sleep(wait)
		func() {
			w.locl.Lock()
			defer w.locl.Unlock()
			if err := w.conn.WriteMessage(websocket.TextMessage, []byte("ping")); err != nil {
				ClientMap.Remove(w.conn)
				return
			}
		}()
	}
}

// 批处理数据
func (w *WsClient) ReadLoop() {
	for {
		ty, data, err := w.conn.ReadMessage()
		if err != nil {
			// 出错则移除连接并关闭连接
			w.conn.Close()
			ClientMap.Remove(w.conn)
			w.closeChan <- 1
			break
		}
		// 写数据到队列中去
		w.readChan <- NewWsMessage(ty, data)
	}
}

//func (w *WsClient) HandleLoop() {
//loop:
//	for {
//		select {
//		case msg := <-w.readChan:
//			log.Println("Type:", string(msg.MessageType), string(msg.MessageData))
//		case <-w.closeChan:
//			log.Println("已经被关闭")
//			break loop
//		}
//	}
//}
