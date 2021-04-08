package main

import (
	"github.com/purnaresa/bulwark/crypto"
	"log"
)

func main() {
	// provision key pair
	privateKey, publicKey, err := crypto.GenerateKeyPair()
	if err != nil {
		log.Fatalln(err)
	}
	// create dummy plaintext
	plaintext := []byte("hello, my name is plaintext")
	log.Printf("plaintext : %s\n\n", string(plaintext))
	// create signature
	log.Println("creating signature...")
	signature, err := crypto.SignDefault(plaintext, privateKey)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("signature : %s\n\n", string(signature))
	}
	// verify signature
	log.Println("verifying signature and plaintext...")
	errVerify := crypto.VerifyDefault(plaintext, publicKey, signature)
	if errVerify != nil {
		log.Fatalln(errVerify)
	} else {
		log.Println("verification success!")
	}
}
