package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"io/ioutil"
	"log"
)

func TripleDesEncrypt(file string) {
	// because we are going to use TripleDES... therefore we Triple it!
	key := "12345678" + "12345678" + "12345678"
	plaintext, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	block, err := des.NewTripleDESCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	ciphertext := []byte(key)
	iv := ciphertext[:des.BlockSize]
	origData := PKCS5Padding(plaintext, block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(origData))
	mode.CryptBlocks(encrypted, origData)

	err = ioutil.WriteFile("d_ciphertext.bin", encrypted, 0777)
	if err != nil {
		log.Panic(err)
	}
}

func TripleDesDecrypt() {
	key := "12345678" + "12345678" + "12345678"
	data, err := ioutil.ReadFile("ciphertext.bin")
	if err != nil {
		log.Fatal(err)
	}

	block, err := des.NewTripleDESCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	ciphertext := []byte(key)
	iv := ciphertext[:des.BlockSize]

	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(data))
	decrypter.CryptBlocks(decrypted, data)
	decrypted = PKCS5UnPadding(decrypted)

	err = ioutil.WriteFile("decrypted.txt", decrypted, 0777)
	if err != nil {
		log.Panic(err)
	}
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
