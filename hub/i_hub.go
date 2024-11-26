package wsservicehub

import (
	wsmessage "first_socket/pkg/ws_service/ws_message"

	client "first_socket/pkg/ws_service/client"
	clientRepo "first_socket/pkg/ws_service/repository"

	"github.com/gorilla/websocket"
)

type IWSClientHub[WSMessage wsmessage.IWSMessage] interface {
	AddClient(string, string, *websocket.Conn) (
		client.IWSClient[WSMessage],
		error,
	)
	RemoveUser(string)
	RemoveUserClient(string, string)

	SendUser(string, WSMessage)
	SendUserWithoutClient(string, string, WSMessage)
	SendUsers([]string, WSMessage)
}

func NewWSHub[
	WSMessage wsmessage.IWSMessage,
](clientRepo clientRepo.IClientRepository[WSMessage]) IWSClientHub[WSMessage] {
	return &hub[WSMessage]{
		clientRepo: clientRepo,
	}
}
