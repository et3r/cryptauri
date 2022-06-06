package main

import (
	"cryptauri/decrypter"
	"cryptauri/encrypter"
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/fatih/color"
	"os"
)

func main() {

	// Output Bold Colors
	green := color.New(color.FgGreen, color.Bold)     // Green Color
	white := color.New(color.FgWhite, color.Bold)     // White Color
	magenta := color.New(color.FgMagenta, color.Bold) // Magenta color
	red := color.New(color.FgRed, color.Bold)         // Red Color

	// Args Controller
	// Defining the name of the app and its desc
	parser := argparse.NewParser("cryptauri", "Encrypt files and messages")
	// Defining flags to control the flow of the program, they both return a bool
	var encrypt *bool = parser.Flag("e", "encrypt", &argparse.Options{Required: false, Help: "Encrypts a message"})
	var decrypt *bool = parser.Flag("d", "decrypt", &argparse.Options{Required: false, Help: "Decrypts a message"})
	// Defining the verbose, it returns the number of verbose that the user uses (-vvv)
	var verbose *int = parser.FlagCounter("v", "verbose", &argparse.Options{Required: false, Help: "Show more information"})
	// Defining the flags with value, they all are required
	message := parser.String("m", "message", &argparse.Options{Required: true, Help: "Message to encrypt/decrypt"}) // Message to encrypt/decrypt
	password := parser.String("p", "password", &argparse.Options{Required: true, Help: "Password for message"})     // Password to merge/decrypt
	level := parser.Int("l", "level", &argparse.Options{Required: false, Default: 1, Help: "Level of encryption"})  // Level where the message is encrypted/is going to be encrypted
	// Parsing the args
	err := parser.Parse(os.Args)
	// If the user doesn't choose any option or if he chooses more than 1 option
	if (*encrypt == false && *decrypt == false) || (*encrypt == true && *decrypt == true) {
		red.Println("\n[x] You must choose one option: [-e/--encrypt] [-d/--decrypt]\n")
		os.Exit(2) // The application will exit with code status of 2
	}
	// If there's an error
	if err != nil {
		fmt.Println(parser.Usage(err)) // It will print the help panel
		os.Exit(2)                     // It will exit with code status of 2
	}
	// If the user chose the option of encrypt
	if *encrypt == true {
		messageEncrypted := encrypter.EncryptMessage(*message, *password, uint8(*level), *verbose) // Call the function of encryption
		green.Print("\n[*] Message Encrypted Successfully: \n\n")                                  // Informative Output
		white.Println(messageEncrypted[0])                                                         // Output the message encrypted
		if len(messageEncrypted) > 1 {                                                             // If the length of the slice returned is longer than 1, that means that there's verbose
			magenta.Print("\n[*] Password: ")   // Informative Output
			white.Println(messageEncrypted[1])  // Output the password to use to encrypt
			red.Print("[*] Encryption Level: ") // Informative Output
			white.Println(messageEncrypted[2])  // Output the level of encryption
		}
		fmt.Println() // line break
	}

	if *decrypt == true {
		messageDecrypted := decrypter.Decrypter(*message, *password, uint8(*level))
		green.Println("\n[*] Message Decrypted: \n")
		white.Println(messageDecrypted)
		fmt.Println()
	}
}
