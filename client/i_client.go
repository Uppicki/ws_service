package wsserviceclient

import (
	wsmessage "first_socket/pkg/ws_service/ws_message"
	wsrequest "first_socket/pkg/ws_service/ws_request"

	"github.com/gorilla/websocket"
)

type IWSClient[WSMessage wsmessage.IWSMessage] interface {
	Run()
	Close()

	GetReceivedChan() <-chan WSMessage
	Send(WSMessage)

	GetOwnerLogin() string
	GetConnKey() string
}

func NewWSClient[
	WSMessage wsmessage.IWSMessage,
	WSRequest wsrequest.IWSRequest,
](
	ownerLogin string,
	connKey string,
	conn *websocket.Conn,
) IWSClient[WSMessage] {
	return &wsClient[WSMessage, WSRequest]{
		ownerLogin:      ownerLogin,
		connKey:         connKey,
		conn:            conn,
		sendedMessage:   make(chan WSMessage),
		receivedMessage: make(chan WSMessage),
		isActive:        false,
	}
}
