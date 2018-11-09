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

	log.Printf("Signature 1: %s", sign1.GetHexString())
	sign1.Add(sign2)
	log.Printf("Signature 1 + 2 Aggregation: %s", sign1.GetHexString())
	pub1.Add(pub2)
	if !sign1.Verify(pub1, m) {
		log.Fatal("Aggregate signature does not verify")
	}
	log.Fatal("Aggregate signature verifies correctly!")
}
