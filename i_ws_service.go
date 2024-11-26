package wsservice

import (
	wsmessage "first_socket/pkg/ws_service/ws_message"
	wsrequest "first_socket/pkg/ws_service/ws_request"

	client "first_socket/pkg/ws_service/client"
	wshub "first_socket/pkg/ws_service/hub"
	clientRepo "first_socket/pkg/ws_service/repository"

	"net/http"

	"github.com/gorilla/websocket"
)

type IWSService[WSMessage wsmessage.IWSMessage] interface {
	ServeWS(
		string,
		string,
		Payload,
	) error
	CreateConnection(
		http.ResponseWriter,
		*http.Request,
		http.Header,
	) (*websocket.Conn, error)
	Listen(client.IWSClient[WSMessage])
}

func DefaultWSService[
	WSMessage wsmessage.IWSMessage,
	WSRequest wsrequest.IWSRequest,
]() IWSService[WSMessage] {
	clientRepo := clientRepo.NewClientRepository[WSMessage, WSRequest]()
	hub := wshub.NewWSHub[WSMessage](clientRepo)

	return &wsService[WSMessage]{
		hub: hub,
	}
}
