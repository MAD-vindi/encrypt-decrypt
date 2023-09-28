package main

import (
	"fmt"

	"github.com/MAD-vindi/encrypt-decrypt/decrypt"
	"github.com/MAD-vindi/encrypt-decrypt/encrypt"
)

func main() {
	fmt.Println("Enter a string that you want to encrypt")
	var inputStr string
	fmt.Scanf("%s", &inputStr)
	temp := encrypt.Encrypt(inputStr)
	fmt.Printf("After encrypting %q you'll get %q\n", inputStr, temp)
	fmt.Printf("After decrypting %q you'll get %q\n", temp, decrypt.Decrypt(temp))
}
