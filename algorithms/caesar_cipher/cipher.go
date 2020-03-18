package main

import (
	"fmt"
	"flag"
	"bytes"
	"strings"
)

func main() {
	cipherKey := flag.Int("c", 0, "Cipher Shifting Between (-26, 26)")

	input := flag.String("i", "", "Input")
	flag.Parse()

	if *cipherKey > 26 || *cipherKey < -26 {
		flag.PrintDefaults()
	} else {
		fmt.Println(caesarsCipher(*input, *cipherKey))
	}
}

func caesarsCipher(user_input string, key int) string {
	var outputBuffer bytes.Buffer

	for _, r := range strings.ToLower(user_input) {
		newByte := int(r)

		if newByte >= 'a' && newByte <= 'z' {
			newByte += key

			if newByte > 'z' {
				newByte -= 26
			} else if newByte < 'a' {
				newByte += 26
			}
		}
		outputBuffer.WriteString(string(newByte))
	}
	return outputBuffer.String()
}