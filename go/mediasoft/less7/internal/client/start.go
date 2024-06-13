package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"less7/pkg/api"
)

func Start() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			var message = new(api.MessageOut)

			text := scanner.Text()

			err := json.Unmarshal([]byte(text), message)
			if err != nil {
				fmt.Printf("System message: %s\n", text)
				continue
			}
			fmt.Printf("Message from %s: %s\n", message.From, message.Body)
		}
	}()

	inputScanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your unique nickname: ")
	if inputScanner.Scan() {
		conn.Write([]byte(inputScanner.Text() + "\n"))
	} else {
		return
	}

	fmt.Println("Type message to send, 'exit' to quit:")
	for {
		msg, msgOk := scanMessage(inputScanner) // вызов функции из другого файла
		if !msgOk {
			break
		}

		msgJson, err := json.Marshal(msg)
		if err != nil {
			fmt.Println(err)
			continue
		}

		conn.Write([]byte(string(msgJson) + "\n"))
	}

	conn.Close()
}
