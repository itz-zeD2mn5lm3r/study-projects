package server

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"sync"

	"less7/pkg/api"
)

var (
	clients      = make(map[string]net.Conn)
	clientsMutex sync.Mutex
)

func Handler(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	var nick string // Объявляем переменную nick в области видимости функции Handler

	if scanner.Scan() {
		nick = scanner.Text() // Присваиваем значение переменной nick
		clientsMutex.Lock()
		if _, exists := clients[nick]; exists {
			conn.Write([]byte("Nick is already taken. Please choose another one.\n"))
			clientsMutex.Unlock()
			return
		}
		clients[nick] = conn
		clientsMutex.Unlock()
		defer func() {
			clientsMutex.Lock()
			delete(clients, nick)
			clientsMutex.Unlock()
		}()
	}

	for scanner.Scan() {
		var msg api.MessageIn
		if err := json.Unmarshal(scanner.Bytes(), &msg); err != nil {
			fmt.Println("Error unmarshaling message:", err)
			continue
		}

		clientsMutex.Lock()
		receiverConn, exists := clients[msg.To]
		clientsMutex.Unlock()

		if !exists {
			conn.Write([]byte("User not found.\n"))
			continue
		}

		response := api.MessageOut{
			From: nick, // Используем переменную nick
			Body: msg.Body,
		}

		responseData, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error marshaling response:", err)
			continue
		}

		receiverConn.Write(append(responseData, '\n'))
	}
}
