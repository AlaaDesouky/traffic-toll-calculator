package main

import (
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"traffic-toll-calculator/types"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var (
	bufferSize = 1028
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	recv := NewDataReceiver()

	http.HandleFunc("/ws", recv.handleWS)

	ws_port := os.Getenv("WS_PORT")

	log.Printf("starting ws connection on port: %s\n", ws_port)
	http.ListenAndServe(ws_port, nil)
}

type DataReceiver struct {
	msgCh chan types.OBUData
	conn  *websocket.Conn
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgCh: make(chan types.OBUData, 128),
	}
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  bufferSize,
		WriteBufferSize: bufferSize,
	}

	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn

	go dr.wsReceiverLoop()
}

func (dr *DataReceiver) wsReceiverLoop() {
	log.Println("new OBU client connected")

	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error:", err)
			continue
		}

		data.RequestID = rand.Intn(math.MaxInt)
		log.Printf("received data: %+v", data)
	}
}
