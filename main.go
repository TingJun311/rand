package rand

import (
	r "crypto/rand"
	"math/big"
	"math/rand"
	"time"
)

func RandomNumN(number int64) (*big.Int) {

	num, err := r.Int(r.Reader, big.NewInt(number))
	if err != nil {
		panic(err)
	}
	return num
}


func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GenerateRandomBigInt(numBytes int) (*big.Int, error) {
    value := make([]byte, numBytes)
    _, err := rand.Read(value)
    if err != nil {
        return nil, err
    }

    for {
        if value[0] != 0 {
            break
        }
        firstByte := value[:1]
        _, err := rand.Read(firstByte)
        if err != nil {
            return nil, err
        }
    }

    nums := (&big.Int{}).SetBytes(value)
	return nums, nil
}