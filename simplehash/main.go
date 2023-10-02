package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
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

	hashValue := simpleHash(input)
	fmt.Println("Hash:", hashValue)
}

func simpleHash(input string) string {
	// Создаем хеш-сумму MD5 из строки
	hash := md5.New()
	io.WriteString(hash, input)
	md5sum := hash.Sum(nil)

	// Преобразуем хеш-сумму в строку из 32 шестнадцатеричных символов
	md5hex := hex.EncodeToString(md5sum)

	// Берем первые 5 символов хеша
	return md5hex[:5]
}
