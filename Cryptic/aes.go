package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"log"
)

func AES_Encrypt(file string) {
	log.Print("Encrypting file with AES-256")

	plaintext, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// The key should be 16 bytes (AES-128), 24 bytes (AES-192) or
	// 32 bytes (AES-256)
	k := "9DGKmRcJoakBKGgbOJt4DoEfuUGEOMLA"

	block, err := aes.NewCipher([]byte(k))
	if err != nil {
		log.Panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Panic(err)
	}

	// Never use more than 2^32 random nonces with a given key
	// because of the risk of repeat.
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	// Save back to file
	err = ioutil.WriteFile("a_ciphertext.bin", ciphertext, 0777)
	if err != nil {
		log.Panic(err)
	}
}

func AES_Decrypt(file string) {
	log.Print("Decrypting file with AES-256")
	ciphertext, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	k := "9DGKmRcJoakBKGgbOJt4DoEfuUGEOMLA"
	block, err := aes.NewCipher([]byte(k))
	if err != nil {
		log.Panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Panic(err)
	}

	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile("decrypted.txt", plaintext, 0777)
	if err != nil {
		log.Panic(err)
	}
}
