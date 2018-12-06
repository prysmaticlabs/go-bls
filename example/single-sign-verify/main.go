package main

import (
	"log"

	"github.com/prysmaticlabs/go-bls"
)

func main() {
	var sec1 bls.SecretKey
	var sec2 bls.SecretKey
	sec1.SetByCSPRNG()
	sec2.SetByCSPRNG()

	pub1 := sec1.GetPublicKey()
	pub2 := sec2.GetPublicKey()

	m := "super special message"
	sign1 := sec1.Sign(m)
	sign2 := sec2.Sign(m)

	log.Printf("Signature 1: 0x%x", sign1.HexString())
	log.Printf("Signature 2: 0x%x", sign2.HexString())
	sign1.Add(sign2)
	log.Printf("Signature 1 + 2 Aggregation: 0x%x", sign1.HexString())
	pub1.Add(pub2)
	if !sign1.Verify(pub1, m) {
		log.Fatal("Aggregate Signature Does Not Verify")
	}
	log.Println("Aggregate Signature Verifies Correctly!")
}
