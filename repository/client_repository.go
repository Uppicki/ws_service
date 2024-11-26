package wsservicerepository

import (
	wsmessage "first_socket/pkg/ws_service/ws_message"
	wsrequest "first_socket/pkg/ws_service/ws_request"

	client "first_socket/pkg/ws_service/client"
	wsservicestore "first_socket/pkg/ws_service/store"

	"github.com/gorilla/websocket"
)

type clientRepository[
	WSMessage wsmessage.IWSMessage,
	WSRequest wsrequest.IWSRequest,
] struct {
	store wsservicestore.IStore[WSMessage]
}

func (repo *clientRepository[WSMessage, WSRequest]) CreateClient(
	ownerLogin string,
	connKey string,
	conn *websocket.Conn,
) client.IWSClient[WSMessage] {
	return client.NewWSClient[WSMessage, WSRequest](ownerLogin, connKey, conn)
}

func (repo *clientRepository[WSMessage, WSRequest]) AddClient(
	client client.IWSClient[WSMessage],
) error {
	return repo.store.AddClient(client)
}

func (repo *clientRepository[WSMessage, WSRequest]) RemoveUser(
	ownerLogin string,
) {
	repo.store.RemoveUser(ownerLogin)
}

func (repo *clientRepository[WSMessage, WSRequest]) RemoveClient(
	ownerLogin string,
	connKey string,
) {
	repo.store.RemoveClient(ownerLogin, connKey)
}

func (repo *clientRepository[WSMessage, WSRequest]) GetUserClients(
	ownerLogin string,
) []client.IWSClient[WSMessage] {
	return repo.store.GetUserClients(ownerLogin)
}

func (repo *clientRepository[WSMessage, WSRequest]) GetUserWithoutClient(
	ownerLogin string,
	connKey string,
) []client.IWSClient[WSMessage] {
	return repo.store.GetUserWithoutClient(ownerLogin, connKey)
}

func (repo *clientRepository[WSMessage, WSRequest]) GetUsersClients(
	logins []string,
) []client.IWSClient[WSMessage] {
	return repo.store.GetUsersClients(logins)
}
