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

func UniqueId(iLength int) string {
	// Initialize
	aLetters := []string("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	sUniqueId := make([]string, iLength)

	// Loop through string letters
	for iIndex := range iLength {
		sUniqueId[iIndex] = aLetters[rand.Intn(len(aLetters))]
	}

	return string(sUniqueId)
}
