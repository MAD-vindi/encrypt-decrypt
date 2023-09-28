package decrypt

func Decrypt(s string) (decryptedStr string) {
	for _, ch := range s {
		asciiCode := int(ch)
		char := string(rune(asciiCode - 3))
		decryptedStr += char
	}
	return
}
