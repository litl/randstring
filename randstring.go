package randstring

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"math/big"
)

/*
   Use mixed-case alphanumeric characters, minus vowels so we don't
   get naughty words. This leaves us with 10+21+21=52 possibilities
   per character, or 5.7 bits (-log(1/52, 2)) of information.

   Thus a random string of length 15 gives us 85 bits of information
*/
const chars = "0123456789bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"

func randFromString(n int, charSet string) (string, error) {
	max := big.NewInt(int64(len(charSet)))

	bytes := make([]byte, n)
	for i := range bytes {
		j, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		bytes[i] = charSet[int(j.Int64())]
	}

	return string(bytes), nil
}

func AlphaNum(n int) (string, error) {
	return randFromString(n, chars)
}

func Numeric(n int) (string, error) {
	return randFromString(n, chars[0:10])
}

func Hex(n int) (string, error) {
	bytes := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
