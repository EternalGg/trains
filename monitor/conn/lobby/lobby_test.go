package lobby

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"testing"
	"time"
	"train/monitor/conn"
)

func TestConn(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := "ws://localhost:8080/ws"
	log.Printf("connecting to %s", u)

	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		time.Sleep(1000000000)
		session := conn.Session{}
		player := conn.Player{
			ID:   1,
			Name: "zzz",
		}
		p, _ := json.Marshal(player)
		session.Data = p
		session.Type = 1
		ss, _ := json.Marshal(session)
		err1 := c.WriteMessage(1, ss)
		if err != nil {
			log.Println("write:", t)
			fmt.Println(err1)
			return
		}
	}
}
