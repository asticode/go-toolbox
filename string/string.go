package string

import (
	"bytes"
	"math/rand"
)

func Concatenate(aStrings []string) string {
	var oBuffer bytes.Buffer

	// Loop through strings
	for _, sString := range aStrings {
		oBuffer.WriteString(sString)
	}

	// Return buffer
	return oBuffer.String()
}

func RandomString(iLength int) string {
	// Initialize
	sLetters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	sRandomString := make([]byte, iLength)

	// Loop through string letters
	for iIndex := range sRandomString {
		sRandomString[iIndex] = sLetters[rand.Int63n(int64(len(sLetters)))]
	}

	// Return
	return string(sRandomString)
}
