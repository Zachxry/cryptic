package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

// methods
var aes_ bool
var des_ bool
var rsa_ bool
var random bool
var encrypt bool
var decrypt bool
var all bool

func init() {
	flag.BoolVar(&aes_, "aes", false, "encryption type")
	flag.BoolVar(&des_, "des", false, "encryption type")
	flag.BoolVar(&rsa_, "rsa", false, "encryption type")
	flag.BoolVar(&random, "random", false, "encryption type")
	flag.BoolVar(&all, "all", false, "encrypt all files in directory")
	flag.BoolVar(&encrypt, "encrypt", false, "encrypting file")
	flag.BoolVar(&decrypt, "decrypt", false, "decrypting file")
}

func main() {
	file := flag.String("file", "file.txt", "file specified")
	flag.Parse()

	if all {
		if aes_ {
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
					AES_Encrypt(file.Name())
				}
			}
		}

		if des_ {
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
					TripleDesEncrypt(file.Name())
				}
			}
		}

		if rsa_ {
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
					RsaEncrypt(file.Name())
				}
			}
		}
	}

	// aes encrypt and decrypt
	if aes_ {
		if encrypt {
			AES_Encrypt(*file)
		}

		if decrypt {
			AES_Decrypt(*file)
		}
	}

	// des encrypt and decrypt
	if des_ {
		if encrypt {
			TripleDesEncrypt(*file)
		}

		if decrypt {
			TripleDesDecrypt(*file)
		}
	}

	if rsa_ {
		if encrypt {
			RsaEncrypt(*file)
		}
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
