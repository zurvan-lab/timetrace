package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomNumber(min, max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return 0, err
	}

	return int(n.Int64()) + min, nil
}

func BytesToString(data []byte) string {
	return string(data)
}

func StringToBytes(s string) []byte {
	return []byte(s)
}
