package wsserviceclient

import (
	wsmessage "first_socket/pkg/ws_service/ws_message"
	wsrequest "first_socket/pkg/ws_service/ws_request"
	"sync"

	"github.com/gorilla/websocket"
)

type wsClient[
	WSMessage wsmessage.IWSMessage,
	WSRequest wsrequest.IWSRequest,
] struct {
	conn            *websocket.Conn
	ownerLogin      string
	connKey         string
	sendedMessage   chan WSMessage
	receivedMessage chan WSMessage
	isActive        bool
	mu              sync.Mutex
}

func (client *wsClient[WSMessage, WSRequest]) Run() {
	client.mu.Lock()
	client.isActive = true
	client.mu.Unlock()
	go client.readerRun()
	go client.writerRun()
}

func (client *wsClient[WSMessage, WSRequest]) readerRun() {
	defer func() {
		client.mu.Lock()
		if client.isActive {
			client.receivedMessage <- wsmessage.DisconnectedMessage(
				client.ownerLogin,
			).(WSMessage)
			client.isActive = false
		}
		client.mu.Unlock()
	}()

	client.receivedMessage <- wsmessage.ConnectedMessage(
		client.ownerLogin,
	).(WSMessage)

	for client.isActive {
		var req WSRequest

		if err := client.conn.ReadJSON(&req); err != nil {
			break
		}

		if message, err := req.ToMessage(); err == nil {
			client.receivedMessage <- message.(WSMessage)
		} else {
			break
		}
	}
}

func (client *wsClient[WSMessage, WSRequest]) writerRun() {
	for client.isActive {
		select {
		case message := <-client.sendedMessage:
			response, err := message.ToResponse()
			if err != nil {
				break
			}

			if innerErr := client.conn.WriteJSON(response); innerErr != nil {
				break
			}
		}
	}
}

func (client *wsClient[WSMessage, WSRequest]) GetReceivedChan() <-chan WSMessage {
	return client.receivedMessage
}

func (client *wsClient[WSMessage, WSRequest]) Send(message WSMessage) {
	select {
	case client.sendedMessage <- message:
	default:
	}
}

func (client *wsClient[WSMessage, WSRequest]) Close() {
	client.mu.Lock()
	defer client.mu.Unlock()

	client.isActive = false
	client.conn.Close()
	close(client.sendedMessage)
	close(client.receivedMessage)
}

func (client *wsClient[WSMessage, WSRequest]) GetOwnerLogin() string {
	return client.ownerLogin
}

func (client *wsClient[WSMessage, WSRequest]) GetConnKey() string {
	return client.connKey
}
