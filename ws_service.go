package wsservice

import (
	wsmessage "first_socket/pkg/ws_service/ws_message"

	wsserviceclient "first_socket/pkg/ws_service/client"
	wsservicehub "first_socket/pkg/ws_service/hub"
	"net/http"

	"github.com/gorilla/websocket"
)

type wsService[
	WSMessage wsmessage.IWSMessage,
] struct {
	hub wsservicehub.IWSClientHub[WSMessage]
}

func (service *wsService[WSMessage]) ServeWS(
	owner string,
	connKey string,
	payload Payload,
) error {
	conn, err := service.CreateConnection(
		payload.Writer,
		payload.Request,
		payload.Header,
	)
	if err != nil {
		return err
	}

	client, clientErr := service.hub.AddClient(owner, connKey, conn)
	if clientErr != nil {
		return err
	}

	go service.Listen(client)
	client.Run()

	return nil
}

func (service *wsService[WSMessage]) Listen(
	client wsserviceclient.IWSClient[WSMessage],
) {
	channel := client.GetReceivedChan()
	for {
		select {
		case message := <-channel:
			message.Map()
		}
	}
}

func (service *wsService[WSMessage]) CreateConnection(
	writer http.ResponseWriter,
	request *http.Request,
	header http.Header,
) (*websocket.Conn, error) {
	return Upgrader.Upgrade(writer, request, header)
}
