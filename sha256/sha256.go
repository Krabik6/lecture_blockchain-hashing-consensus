package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	hashString := getHash("Hello, World!")
	fmt.Println("Hash:", hashString)
}

// Функция, которая получает строку и возвращает ее SHA-256 хеш
func getHash(str string) string {
	inputString := str
	inputBytes := []byte(inputString)
	sha256Hash := sha256.New()
	_, err := sha256Hash.Write(inputBytes)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return ""
	}
	hashBytes := sha256Hash.Sum(nil)
	hashString := fmt.Sprintf("%x", hashBytes)
	return hashString
}
