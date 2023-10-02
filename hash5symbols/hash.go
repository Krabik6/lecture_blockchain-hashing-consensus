package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	input := "Hello, World!"
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
