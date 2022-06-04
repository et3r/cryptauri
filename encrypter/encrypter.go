package encrypter

import "fmt"

func EncryptMessage(message string, password string, level uint8) {

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

	// Making slice of the message and the password joined
	var firstEncrypt []uint
	firstEncrypt = make([]uint, 0, 500)
	for i := 0; i < len(asciiSentence); i++ {
		suma := asciiSentence[i] + adaptedPassword[i]
		firstEncrypt = append(firstEncrypt, uint(suma))
	}

	fmt.Println("Message:")
	fmt.Println(asciiSentence)

	fmt.Println("Password:")
	fmt.Println(asciiPassword)

	fmt.Println("Adapted Password")
	fmt.Println(adaptedPassword)

	fmt.Println("First Encrypt:")
	fmt.Println(firstEncrypt)
}
