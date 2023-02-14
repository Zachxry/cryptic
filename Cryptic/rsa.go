package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
)

// Generate a key pair (private and public) based on number of bits
func generateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	// This method requires a random number of bits.
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}

	// The public key is part of the PrivateKey struct
	return privateKey, &privateKey.PublicKey
}

func RsaEncrypt(file string) {
	privateKey, publicKey := generateKeyPair(2048)

	plaintext, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	message := []byte(plaintext)
	// This method ensures that a different cipher is generated each time
	cipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted: %v\n", cipherText)

	// write encrypted message to a file
	err = ioutil.WriteFile("r_ciphertext.enc", cipherText, 0777)
	if err != nil {
		log.Panic(err)
	}

	decMessage, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, cipherText, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Original: %s\n", string(decMessage))

	// write decrypted message to a file
	err = ioutil.WriteFile("decrypted.txt", decMessage, 0777)
	if err != nil {
		log.Panic(err)
	}

	// We actually sign the hashed message
	msgHash := sha256.New()
	_, err = msgHash.Write(message)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Signature: %v\n", signature)

	err = rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		fmt.Println("Verification failed: ", err)
	} else {
		fmt.Println("Message verified.")
	}
}
