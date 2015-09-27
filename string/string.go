package string

import (
	"bytes"
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
