package onebotv11

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

// The main thread to receive messages.
func socketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	var err error
	ws, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[ONEBOTV11] | Error creating connection: %s\n", err)
	}
	log.Printf("[ONEBOTV11] | Connection Established.\n")
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(3)
	go receiveHandler()
	go sendHandler()
	go actionHandler()
	waitGroup.Wait()
}

// Start the adapter.
func start() {
	BackendHostAddress := os.Getenv("ONEBOTV11_HOST")
	if BackendHostAddress == "" {
		BackendHostAddress = "127.0.0.1:21390"
	}

	log.Println("[ONEBOTV11] | Waiting for Onebot V11 connection...")
	http.HandleFunc("/onebot/v11/ws", socketHandler)
	log.Fatal(http.ListenAndServe(BackendHostAddress, nil))
}

func finalize() {
	if ws != nil {
		ws.Close()
	}
	log.Println("[ONEBOTV11] | Shutting down.")
}
