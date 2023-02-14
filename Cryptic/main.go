package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

// aes flag
var aes256 bool

// des flag
var desTriple bool

// rsa flag
var rsa_ bool

// random flag
var random bool

// enc & dec
var encrypt bool
var decrypt bool

func init() {
	// aes
	flag.BoolVar(&aes256, "aes256", false, "encryption type")

	// des
	flag.BoolVar(&desTriple, "des", false, "encryption type")

	// rsa
	flag.BoolVar(&rsa_, "rsa", false, "encryption type")

	// random method
	flag.BoolVar(&random, "random", false, "encryption type")

	// encrypt & decrypt
	flag.BoolVar(&encrypt, "encrypt", false, "encrypting file")
	flag.BoolVar(&decrypt, "decrypt", false, "decrypting file")
}

func main() {
	file := flag.String("file", "file.txt", "file specified")
	flag.Parse()

	// aes encrypt and decrypt
	if aes256 {
		if encrypt {
			AES_Encrypt(*file)
		}

		if decrypt {
			AES_Decrypt(*file)
		}
	}

	if desTriple {
		TripleDesEncrypt(*file)
		TripleDesDecrypt()
	}

	if random {
		files, err := ioutil.ReadDir(".")
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if file.Name() == "cryptic" {
				continue
			} else if file.Name() == "cryptic.exe" {
				continue
			} else {
				fmt.Println(file.Name(), file.IsDir())
				rand.Seed(time.Now().UnixNano())
				min := 1
				max := 3
				number := rand.Intn(max-min+1) + min
				if number == 1 {
					AES_Encrypt(file.Name())
				} else if number == 2 {
					TripleDesEncrypt(file.Name())
				} else {
					RsaEncrypt(file.Name())
				}
			}
		}
	}
}
