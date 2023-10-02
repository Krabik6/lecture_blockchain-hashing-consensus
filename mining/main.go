package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Print("Введите сложность (больше => сложнее => дольше): ")
	targetBitsStr := ""
	fmt.Scanln(&targetBitsStr)

	targetBits, err := strconv.Atoi(targetBitsStr)
	if err != nil {
		fmt.Println("Ошибка: введите корректное целое число для сложности.")
		return
	}

	startPOW(targetBits)
}

func startPOW(initialDifficulty int) {
	for {
		target := big.NewInt(1)
		target.Lsh(target, uint(256-initialDifficulty))
		nonce, hash := proofOfWork(
			[]byte("Hello World"),
			target,
		)
		fmt.Printf("\nHash: %x\n", hash)
		fmt.Printf("Nonce: %d\n", nonce)
	}
}

func prepareData(data []byte, nonce int) []byte {
	nonceBytes := intToBytes(nonce)

	joinedData := bytes.Join([][]byte{
		data,
		nonceBytes,
	}, []byte{})
	return joinedData
}

func proofOfWork(data []byte, target *big.Int) (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	maxNonce := 0xFFFFFFFF // Максимальное значение для nonce

	for nonce < maxNonce {
		data := prepareData(data, nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

func intToBytes(n int) []byte {
	result := make([]byte, 4)
	for i := 0; i < 4; i++ {
		result[i] = byte(n & 0xff)
		n >>= 8
	}
	return result
}
