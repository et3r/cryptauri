package encrypter

import (
	"fmt"
	"os"
	"strconv"
)

func EncryptMessage(message string, password string, level uint8) string {

	// Making slice of the encrypt-code
	var encryptCode []string
	encryptCode = make([]string, 0, 26)
	encryptCode = append(encryptCode, "%", "$", "@", "/", "!", "?", "¿", "*", "{", "#", "}", ".", "_", "¡", "-", ",", "&", "]", "|", "[", "¬", "+", "-", "=", "~", "ˆ", "`")

	// Making slice of the second encrypt range
	var encryptRange []string
	encryptRange = make([]string, 0, 10)
	for i := level - 1; i < (level + 9); i++ {
		if level >= 19 || level <= 0 {
			fmt.Println("Level must be in range 1-18")
			os.Exit(2)
		}
		encryptRange = append(encryptRange, encryptCode[i])
	}

	// Making slice of each character of the message
	var asciiSentence []uint8
	asciiSentence = make([]uint8, 0, 500) // Type: []uint8, Max Msg Characters: 500
	for i := 0; i < len(message); i++ {
		asciiSentence = append(asciiSentence, message[i])
	}

	// Making slice of each character of the password
	var asciiPassword []uint8
	asciiPassword = make([]uint8, 0, 500)
	for i := 0; i < len(password); i++ {
		asciiPassword = append(asciiPassword, password[i])
	}

	// Making the slice of the adapted password
	var adaptedPassword []uint8
	adaptedPassword = make([]uint8, 0, 500)

	// If the message is longer than the password
	if len(asciiSentence) > len(asciiPassword) {
		for i := 0; i < len(asciiPassword); i++ {
			adaptedPassword = append(adaptedPassword, asciiPassword[i])
			if i == len(asciiPassword)-1 {
				i = -1
			}
			if len(asciiSentence) == len(adaptedPassword) {
				break
			}
		}
	}

	// If the password is longer than the message
	if len(asciiPassword) > len(asciiSentence) {
		for i := 0; i < len(asciiSentence); i++ {
			adaptedPassword = append(adaptedPassword, asciiPassword[i])
		}
	}

	// If the message has the same length of the password
	if len(asciiSentence) == len(asciiPassword) {
		for i := 0; i < len(asciiSentence); i++ {
			adaptedPassword = append(adaptedPassword, asciiPassword[i])
		}
	}

	// Making slice of the first encrypt
	var firstEncrypt []int
	firstEncrypt = make([]int, 0, 500)
	for i := 0; i < len(asciiSentence); i++ {
		suma := asciiSentence[i] + adaptedPassword[i]
		firstEncrypt = append(firstEncrypt, int(suma))
	}

	// Making slice of the second encrypt
	var secondEncrypt []string
	secondEncrypt = make([]string, 0, 500)
	for i := 0; i < len(firstEncrypt); i++ {
		firstEncryptChar := strconv.Itoa(firstEncrypt[i])
		for j := 0; j < len(firstEncryptChar); j++ {

			switch string(firstEncryptChar[j]) {
			case "0":
				secondEncrypt = append(secondEncrypt, encryptRange[0])
				break
			case "1":
				secondEncrypt = append(secondEncrypt, encryptRange[1])
				break
			case "2":
				secondEncrypt = append(secondEncrypt, encryptRange[2])
				break
			case "3":
				secondEncrypt = append(secondEncrypt, encryptRange[3])
				break
			case "4":
				secondEncrypt = append(secondEncrypt, encryptRange[4])
				break
			case "5":
				secondEncrypt = append(secondEncrypt, encryptRange[5])
				break
			case "6":
				secondEncrypt = append(secondEncrypt, encryptRange[6])
				break
			case "7":
				secondEncrypt = append(secondEncrypt, encryptRange[7])
				break
			case "8":
				secondEncrypt = append(secondEncrypt, encryptRange[8])
				break
			case "9":
				secondEncrypt = append(secondEncrypt, encryptRange[9])
				break
			}
			if (len(firstEncryptChar) - 1) == j {
				secondEncrypt = append(secondEncrypt, "§")
			}
		}
	}

	var finalMessage string
	for i := 0; i < len(secondEncrypt); i++ {
		finalMessage += secondEncrypt[i]
	}

	return finalMessage
}
