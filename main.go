package main

import (
	"cryptauri/encrypter"
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func main() {

	parser := argparse.NewParser("cryptauri", "Encrypt files and messages")

	var encrypt *bool = parser.Flag("e", "encrypt", &argparse.Options{Required: false, Help: "Encrypts a message"})
	var decrypt *bool = parser.Flag("d", "decrypt", &argparse.Options{Required: false, Help: "Decrypts a message"})

	message := parser.String("m", "message", &argparse.Options{Required: true, Help: "Message to encrypt/decrypt"})
	password := parser.String("p", "password", &argparse.Options{Required: true, Help: "Password for message"})
	level := parser.Int("l", "level", &argparse.Options{Required: false, Default: 1, Help: "Level of encryption"})

	err := parser.Parse(os.Args)

	if *encrypt == false && *decrypt == false {
		fmt.Println("\n[x] You must choose an option: [-e/--encrypt] [-d/--decrypt]\n")
		os.Exit(2)
	}

	if err != nil {
		fmt.Println(parser.Usage(err))
		os.Exit(2)
	}

	if *encrypt == true {
		messageEncrypted := encrypter.EncryptMessage(*message, *password, uint8(*level))
		fmt.Println()
		fmt.Println("[*] Encrypted Message:")
		fmt.Println(messageEncrypted)
		fmt.Println()
	}
}
