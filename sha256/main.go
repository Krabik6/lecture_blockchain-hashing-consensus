package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Введите строку:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Ошибка при чтении ввода: %v\n", err)
		os.Exit(1)
	}

	input = strings.TrimSpace(input) // Удаление лишних пробелов и символов новой строки

	hashString := getHash(input)
	fmt.Println("Hash:", hashString)
}

// Функция, которая получает строку и возвращает ее SHA-256 хеш
func getHash(str string) string {
	inputString := str
	inputBytes := []byte(inputString)
	sha256Hash := sha256.New()
	_, err := sha256Hash.Write(inputBytes)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return ""
	}
	hashBytes := sha256Hash.Sum(nil)
	hashString := fmt.Sprintf("%x", hashBytes)
	return hashString
}
