package main

import (
	"fmt"

	"github.com/otiai10/gosseract"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	// client.SetLanguage("eng", "spa")
	client.SetImage("/Users/ed/go/src/ocr-contract/captcha.jpg")
	text, _ := client.Text()
	fmt.Println(text)
	// Hello, World!
}
