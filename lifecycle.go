package onebotv11

import (
	"net/http"
	"os"
	"sync"

	"github.com/gonebot-dev/gonebot/logging"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
)

// The main thread to receive messages.
func socketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	var err error
	ws, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		logging.Logf(zerolog.ErrorLevel, "OneBotV11", "Error creating connection: %s\n", err)
	}
	logging.Logf(zerolog.InfoLevel, "OneBotV11", "Connection Established.\n")
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

	logging.Log(zerolog.InfoLevel, "OneBotV11", "Waiting for Onebot V11 connection...")
	http.HandleFunc("/onebot/v11/ws", socketHandler)
	err := http.ListenAndServe(BackendHostAddress, nil)
	if err != nil {
		logging.Logf(zerolog.FatalLevel, "OneBotV11", "Error establishing connection: %s\n", err)
	}
}

func finalize() {
	if ws != nil {
		ws.Close()
	}
	logging.Log(zerolog.InfoLevel, "OneBotV11", "Shutting down.")
}
