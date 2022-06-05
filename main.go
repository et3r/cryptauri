package main

import (
	"cryptauri/encrypter"
	"fmt"
)

func main() {
	messageEncrypted := encrypter.EncryptMessage("undecipherable message", "{5up3rP@ssw0rdw0w!--$}", 7)
	fmt.Println("[*] Encrypted Message:")
	fmt.Println(messageEncrypted)
}
