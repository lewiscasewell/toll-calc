package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	sendInterval = 2
)

type OBUData struct {
	OBUID     int     `json:"obuID"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

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
		ids[i] = i + 1
	}
	return ids
}

func main() {
	for {
		fmt.Println(genLocation())
		time.Sleep(sendInterval * time.Second)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
