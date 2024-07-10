package main

import (
	"log"
	"math"
	"math/rand"
	"os"
	"time"
	"traffic-toll-calculator/types"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var (
	obuIDSCount = 20
	sendInterval = time.Second * 5
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	obuIDS := generateOBUIDS(obuIDSCount)

	ws_endpoint := os.Getenv("WS_ENDPOINT")
	conn, _, err := websocket.DefaultDialer.Dial(ws_endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		for _, obuID := range obuIDS {
			lat, lng := generateLatLng()
			data := types.OBUData{
				OBUID: obuID,
				Lat: lat,
				Lng: lng,
			}
			log.Println("sending data over ws...")
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		} 
		time.Sleep(sendInterval)
	}
}

func generateCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()

	return n + f
}

func generateLatLng() (float64, float64) {
	return generateCoord(), generateCoord()
}

func generateOBUIDS(n int) [] int {
	ids := make([]int, n)
	for i := range n {
		ids[i] = rand.Intn(math.MaxInt)
	}

	return ids
}