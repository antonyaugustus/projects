package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"os"
	"time"

	b64 "encoding/base64"
)

func main() {
	argCount := len(os.Args[1:])

	if argCount != 1 {
		fmt.Printf("Usage: %s %s", os.Args[0], "<password to encrypt>")
		os.Exit(1)
	}

	token := make([]byte, 32)
	rand.Seed(time.Now().UnixNano())
	rand.Read(token)

	h := sha256.New()
	h.Write([]byte(os.Args[1]))

	fmt.Printf("%s %s",
		b64.StdEncoding.EncodeToString([]byte(token)),
		b64.StdEncoding.EncodeToString([]byte(h.Sum(nil))))
}
