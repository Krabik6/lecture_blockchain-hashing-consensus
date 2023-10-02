package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"time"
)

const (
	adjustmentStep = 10 * time.Second
	targetTime     = 15 * time.Second // Желаемое время для вычисления одного блока
)

var targetBits = 15

func main() {
	startPOW()
}

func startPOW() {
	for {
		target := big.NewInt(1)
		target.Lsh(target, uint(256-targetBits))
		nonce, hash := proofOfWork(
			[]byte("Hello World"),
			target)
		fmt.Printf("\nHash: %x\n", hash)
		fmt.Printf("Nonce: %d\n", nonce)
		time.Sleep(adjustmentStep)
	}
}

func prepareData(data []byte, nonce int, timestamp int64) []byte {
	nonceBytes := intToBytes(nonce)
	timestampBytes := int64ToBytes(timestamp)

	joinedData := bytes.Join([][]byte{
		data,
		nonceBytes,
		timestampBytes,
	}, []byte{})
	return joinedData
}

func proofOfWork(data []byte, target *big.Int) (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	startTime := time.Now()
	var duration time.Duration

	maxNonce := math.MaxInt64

	for nonce < maxNonce {
		timestamp := time.Now().Unix()
		data := prepareData(data, nonce, timestamp)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(target) == -1 {
			duration = time.Since(startTime)
			fmt.Printf("Difficulty: %d\n", targetBits)
			fmt.Printf("\nDuration: %s\n", duration)
			if duration < targetTime {
				// Уменьшаем сложность, блок нашелся быстрее, делаем задачу сложнее
				targetBits++
			} else {
				// Увеличиваем сложность, блок нашелся медленнее, делаем задачу проще
				targetBits--
			}
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

func int64ToBytes(n int64) []byte {
	result := make([]byte, 8)
	for i := 0; i < 8; i++ {
		result[i] = byte(n & 0xff)
		n >>= 8
	}
	return result
}
