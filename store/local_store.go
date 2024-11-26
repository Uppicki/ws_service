package wsservicestore

import (
	wsmessage "first_socket/pkg/ws_service/ws_message"

	"errors"
	"first_socket/pkg/ws_service/client"
)

type localStore[WSMessage wsmessage.IWSMessage] struct {
	clients map[string]map[string]wsserviceclient.IWSClient[WSMessage]
}

func (store *localStore[WSMessage]) AddClient(
	client wsserviceclient.IWSClient[WSMessage],
) error {
	login, connKey := client.GetOwnerLogin(), client.GetConnKey()

	if _, ok := store.clients[login][connKey]; ok {
		return errors.New("uncorrect connKey")
	}

	if _, ok := store.clients[login]; !ok {
		store.clients[login] = make(map[string]wsserviceclient.IWSClient[WSMessage])
	}

	store.clients[login][connKey] = client

	return nil
}

func (store *localStore[WSMessage]) RemoveUser(login string) {
	if user, ok := store.clients[login]; ok {
		for _, client := range user {
			client.Close()
		}
		delete(user, login)
	}
}

func (store *localStore[WSMessage]) RemoveClient(login string, connKey string) {
	if user, ok := store.clients[login]; ok {
		if client, ok := user[connKey]; ok {
			client.Close()
			delete(user, connKey)
		}

		if len(user) == 0 {
			delete(store.clients, login)
		}
	}
}

func (store *localStore[WSMessage]) GetUserClients(login string) []wsserviceclient.IWSClient[WSMessage] {
	clients := make([]wsserviceclient.IWSClient[WSMessage], 0)

	if user, ok := store.clients[login]; ok {
		for _, client := range user {
			clients = append(clients, client)
		}
	}

	return clients
}

func (store *localStore[WSMessage]) GetUserWithoutClient(
	login string,
	connKey string,
) []wsserviceclient.IWSClient[WSMessage] {
	clients := make([]wsserviceclient.IWSClient[WSMessage], 0)

	if user, ok := store.clients[login]; ok {
		for _, client := range user {
			if client.GetConnKey() != connKey {
				clients = append(clients, client)
			}
		}
	}

	return clients
}

func (store *localStore[WSMessage]) GetUsersClients(logins []string) []wsserviceclient.IWSClient[WSMessage] {
	clients := make([]wsserviceclient.IWSClient[WSMessage], 0)

	for _, login := range logins {
		if user, ok := store.clients[login]; ok {
			for _, client := range user {
				clients = append(clients, client)
			}
		}
	}

	return clients
}

func NewLocalStore[WSMessage wsmessage.IWSMessage]() IStore[WSMessage] {
	return &localStore[WSMessage]{
		clients: make(map[string]map[string]wsserviceclient.IWSClient[WSMessage]),
	}
}
