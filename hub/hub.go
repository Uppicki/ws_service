package wsservicehub

import (
	wsserviceclient "first_socket/pkg/ws_service/client"
	wsservicerepository "first_socket/pkg/ws_service/repository"
	wsmessage "first_socket/pkg/ws_service/ws_message"

	"github.com/gorilla/websocket"
)

type hub[WSMessage wsmessage.IWSMessage] struct {
	clientRepo wsservicerepository.IClientRepository[WSMessage]
}

func (hub *hub[WSMessage]) AddClient(
	ownerLogin string,
	connKey string,
	conn *websocket.Conn,
) (
	wsserviceclient.IWSClient[WSMessage],
	error,
) {
	client := hub.clientRepo.CreateClient(ownerLogin, connKey, conn)

	if err := hub.clientRepo.AddClient(client); err != nil {
		return nil, err
	}

	return client, nil
}

func (hub *hub[WSMessage]) RemoveUser(ownerLogin string) {
	hub.clientRepo.RemoveUser(ownerLogin)
}

func (hub *hub[WSMessage]) RemoveUserClient(ownerLogin string, connKey string) {
	hub.clientRepo.RemoveClient(ownerLogin, connKey)
}

func (hub *hub[WSMessage]) SendUser(
	ownerLogin string,
	message WSMessage,
) {
	clients := hub.clientRepo.GetUserClients(ownerLogin)

	hub.sendClients(clients, message)
}

func (hub *hub[WSMessage]) SendUserWithoutClient(
	ownerLogin string,
	connKey string,
	message WSMessage,
) {
	clients := hub.clientRepo.GetUserWithoutClient(ownerLogin, connKey)

	hub.sendClients(clients, message)
}

func (hub *hub[WSMessage]) SendUsers(
	ownerLogins []string,
	message WSMessage,
) {
	clients := hub.clientRepo.GetUsersClients(ownerLogins)

	hub.sendClients(clients, message)
}

func (hub *hub[WSMessage]) sendClients(
	clients []wsserviceclient.IWSClient[WSMessage],
	message WSMessage,
) {
	for _, client := range clients {
		client.Send(message)
	}
}
