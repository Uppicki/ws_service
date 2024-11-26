package wsservicerepository

import (
	wsclient "first_socket/pkg/ws_service/client"
	store "first_socket/pkg/ws_service/store"
	wsmessage "first_socket/pkg/ws_service/ws_message"
	wsrequest "first_socket/pkg/ws_service/ws_request"

	"github.com/gorilla/websocket"
)

type IClientRepository[
	WSMessage wsmessage.IWSMessage,
] interface {
	CreateClient(string, string, *websocket.Conn) wsclient.IWSClient[WSMessage]
	AddClient(wsclient.IWSClient[WSMessage]) error

	RemoveUser(string)

	RemoveClient(string, string)

	GetUserClients(string) []wsclient.IWSClient[WSMessage]

	GetUserWithoutClient(string, string) []wsclient.IWSClient[WSMessage]

	GetUsersClients([]string) []wsclient.IWSClient[WSMessage]
}

func NewClientRepository[
	WSMessage wsmessage.IWSMessage,
	WSRequest wsrequest.IWSRequest,
]() IClientRepository[WSMessage] {
	return &clientRepository[WSMessage, WSRequest]{
		store: store.NewLocalStore[WSMessage](),
	}
}
