package main

import (
	"log"
	"time"

	bls "github.com/prysmaticlabs/go-bls"
)

func main() {
	m := []byte("super special message")
	startTime := time.Now()
	bls.HashAndMapToG2(m)
	endTime := time.Now()
	log.Printf("HashAndMapToG2 took %f seconds to complete", endTime.Sub(startTime).Seconds())
}
