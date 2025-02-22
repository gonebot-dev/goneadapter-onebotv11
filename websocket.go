package onebotv11

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

func networkConn(ws *websocket.Conn) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go SendingMessage(ws)
	go ReadingMessage(ws)
	waitGroup.Wait()
}

// The main thread to receive messages.
func socketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	var err error
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error creating connection:\n%s\n", err)
	}
	log.Printf("Connection Established.\n")
	defer ws.Close()
	networkConn(ws)
}

func Connector() {
	BackendHostAddress := os.Getenv("ONEBOTV11_HOST")
	if BackendHostAddress == "" {
		BackendHostAddress = "127.0.0.1:21390"
	}

	log.Println("Trying to establish connection with onebot11.")
	http.HandleFunc("/onebot/v11/ws", socketHandler)
	log.Fatal(http.ListenAndServe(BackendHostAddress, nil))
}
