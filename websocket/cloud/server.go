package cloud

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lunny/tango"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for i := 0; true; i++ {
		time.Sleep(time.Second * 3)

		message := "Hello, client " + strconv.Itoa(i)
		err := c.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("write:", err)
			break
		}
		log.Println("echo ...")
	}
	log.Println("echo over")
}

func startTgWebsocket() {
	tg := tango.Classic()
	tg.Get("/echo", echo)

	tg.Run(":8080")
}
