package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lewiscasewell/tolling/types"
)

const (
	sendInterval = time.Second * 5
	wsEndpoint   = "ws://localhost:8080/ws"
)

// func sendOBUData(conn *websocket.Conn,data OBUData) {
// 	err := conn.WriteJSON(data)
// 	if err != nil {
// 		log.Println("write:", err)
// 		return
// 	}

// }

func genLocation() (float64, float64) {
	return genCoordinates(), genCoordinates()
}

func genCoordinates() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}

func genOBUIDs(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func main() {
	obuIds := genOBUIDs(10)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	for {
		for _, id := range obuIds {
			lat, lng := genLocation()
			data := types.OBUData{
				OBUID:     id,
				Latitude:  lat,
				Longitude: lng,
			}
			fmt.Println(data)
			err := conn.WriteJSON(data)
			if err != nil {
				log.Fatal("write:", err)
				return
			}
		}
		time.Sleep(sendInterval)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
