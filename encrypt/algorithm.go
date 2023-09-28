package encrypt

func Encrypt(s string) string {
	encryptedString := ""
	for _, ch := range s {
		asciiCode := int(ch)
		char := string(rune(asciiCode + 3))
		encryptedString += char
	}
	return encryptedString
}
