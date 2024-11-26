package wsservicerequest

import wsmessage "first_socket/pkg/ws_service/ws_message"

type IWSRequest interface {
	ToMessage() (wsmessage.IWSMessage, error)
}
