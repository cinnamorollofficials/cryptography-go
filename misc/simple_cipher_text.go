package main

import (
	"fmt"
	"strings"
)

func shiftTune(r rune, shift int) rune {
	if r >= 'a' && r <= 'z' {
		return 'a' + (r-'a'+rune(shift)+26)%26
	}
	if r >= 'A' && r <= 'Z' {
		return 'A' + (r-'A'+rune(shift)+26)%26
	}
	return r
}

func transform(input string, shift int) string {
	var b strings.Builder
	b.Grow(len(input))

	for _, r := range input {
		b.WriteRune(shiftTune(r, shift))
	}

	return b.String()
}

func encrypt(plaintext string, key int) string {
	return transform(plaintext, key)
}

func decrypt(ciphertext string, key int) string {
	return transform(ciphertext, -key)
}

func main() {
	// define message and key
	plaintext := "Hello Hadi"
	shiftKey := 3

	// encrypt the messsage
	fmt.Printf("Original Text: %s\n", plaintext)
	fmt.Printf("Shift Key: %d\n", shiftKey)

	ciphertext := encrypt(plaintext, shiftKey)
	fmt.Printf("Encrypted Text: %s\n", ciphertext)

	// decrypt the message
	decryptedText := decrypt(ciphertext, shiftKey)
	fmt.Printf("Decrypted Text: %s\n", decryptedText)

}
