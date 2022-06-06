package decrypter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Decrypter(message string, password string, level uint8) string {

	// Making slice of the encrypt-code
	var encryptCode []string
	encryptCode = make([]string, 0, 26)
	encryptCode = append(encryptCode, "%", "$", "@", "/", "!", "?", "¿", "*", "{", "#", "}", ".", "_", "¡", "-", ",", "&", "]", "|", "[", "¬", "+", "-", "=", "~", "ˆ", "`")
	separator := "§"

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

	// Making slice of the message without separator
	var messageWoSep []string
	messageWoSep = make([]string, 0, 1500)
	messageWoSep = strings.Split(message, separator)

	// Making slice of the first encrypt
	var firstEncrypt string

	for i := 0; i < len(messageWoSep); i++ {
		for j := 0; j < len(messageWoSep[i]); j++ {
			switch string(messageWoSep[i][j]) {
			case encryptRange[0]:
				firstEncrypt += "0"
				break
			case encryptRange[1]:
				firstEncrypt += "1"
				break
			case encryptRange[2]:
				firstEncrypt += "2"
				break
			case encryptRange[3]:
				firstEncrypt += "3"
				break
			case encryptRange[4]:
				firstEncrypt += "4"
				break
			case encryptRange[5]:
				firstEncrypt += "5"
				break
			case encryptRange[6]:
				firstEncrypt += "6"
				break
			case encryptRange[7]:
				firstEncrypt += "7"
				break
			case encryptRange[8]:
				firstEncrypt += "8"
				break
			case encryptRange[9]:
				firstEncrypt += "9"
				break
			}
			if (len(messageWoSep[i]) - 1) == j {
				firstEncrypt += " "
			}
		}
	}

	var firstEncryptWoSpaces []string
	firstEncryptWoSpaces = strings.Fields(firstEncrypt)

	var asciiFirstEncrypt []int
	asciiFirstEncrypt = make([]int, 0, 1500)
	for i := 0; i < len(firstEncryptWoSpaces); i++ {
		firstEncryptWoSpacesInt, err := strconv.Atoi(firstEncryptWoSpaces[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		asciiFirstEncrypt = append(asciiFirstEncrypt, firstEncryptWoSpacesInt)
	}

	var asciiPassword []uint8
	asciiPassword = make([]uint8, 0, 500)
	for i := 0; i < len(password); i++ {
		asciiPassword = append(asciiPassword, password[i])
	}

	var adaptedPassword []uint8
	adaptedPassword = make([]uint8, 0, 500)
	if len(asciiFirstEncrypt) > len(asciiPassword) {
		for i := 0; i < len(asciiPassword); i++ {
			adaptedPassword = append(adaptedPassword, asciiPassword[i])
			if i == len(asciiPassword)-1 {
				i = -1
			}
			if len(asciiFirstEncrypt) == len(adaptedPassword) {
				break
			}
		}
	}

	if len(asciiFirstEncrypt) < len(asciiPassword) || len(asciiFirstEncrypt) == len(asciiPassword) {
		for i := 0; i < len(asciiFirstEncrypt); i++ {
			adaptedPassword = append(adaptedPassword, asciiPassword[i])
		}
	}

	var asciiDecryptedMessage []uint8
	asciiDecryptedMessage = make([]uint8, 0, 500)

	for i := 0; i < len(asciiFirstEncrypt); i++ {
		resta := asciiFirstEncrypt[i] - int(adaptedPassword[i])
		asciiDecryptedMessage = append(asciiDecryptedMessage, uint8(resta))
	}

	return string(asciiDecryptedMessage)
}
