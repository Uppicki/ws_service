package wsservicestore

import (
	wsmessage "first_socket/pkg/ws_service/ws_message"

	wsclient "first_socket/pkg/ws_service/client"
)

type IStore[WSMessage wsmessage.IWSMessage] interface {
	AddClient(wsclient.IWSClient[WSMessage]) error

	RemoveUser(string)
	RemoveClient(string, string)

	GetUserClients(string) []wsclient.IWSClient[WSMessage]
	GetUserWithoutClient(string, string) []wsclient.IWSClient[WSMessage]

	GetUsersClients([]string) []wsclient.IWSClient[WSMessage]
}
