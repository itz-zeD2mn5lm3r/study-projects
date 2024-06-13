package client

import (
	"bufio"
	"fmt"

	"less7/pkg/api"
)

func scanMessage(inputScanner *bufio.Scanner) (*api.MessageIn, bool) {
	result := new(api.MessageIn)

	fmt.Print("\nMessage body: ")
	if inputScanner.Scan() {
		text := inputScanner.Text()
		if text == "exit" {
			return nil, false
		}
		result.Body = text
	}

	fmt.Print("\nReceiver nick: ")
	if inputScanner.Scan() {
		text := inputScanner.Text()
		if text == "exit" {
			return nil, false
		}
		result.To = text
	}

	return result, true
}
