package websocket

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	URL := "ws://echo.websocket.org/"

	var dialer *websocket.Dialer

	conn, _, err := dialer.Dial(URL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	go timeWriter(conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return
		}

		fmt.Printf("received: %s\n", message)
	}
}

func timeWriter(conn *websocket.Conn) {
	for {
		time.Sleep(time.Second * 2)
		conn.WriteMessage(websocket.TextMessage, []byte(time.Now().Format("2006-01-02 15:04:05")))
	}
}
