package main

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/prysmaticlabs/go-bls"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var sec bls.SecretKey
	sec.SetByCSPRNG()
	pub := sec.GetPublicKey()
	m := []byte("super special message")

	sig := sec.Sign(m)
	log.Printf("Signature: 0x%x", sig.HexString())
	log.Printf("Size of signature in bytes: %d bytes", len(sig.Serialize()))

	startTime := time.Now()
	numSignatures := 10000

	var wg sync.WaitGroup
	wg.Add(numSignatures)
	// Multithreaded signature verification.
	for i := 0; i < numSignatures; i++ {
		go func() {
			defer wg.Done()
			if !sig.Verify(pub, m) {
				log.Fatal("Signature Does Not Verify")
			}
		}()
	}
	wg.Wait()
	endTime := time.Now()
	log.Printf("Time required to verify %d signatures on %d cores: %f seconds", numSignatures, runtime.NumCPU(), endTime.Sub(startTime).Seconds())
}
