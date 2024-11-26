package wsservicemessage

type WSMessage struct {
	MessageType WSMessageType
	Owner       string
}

func (m *WSMessage) Map() {

}

func (m *WSMessage) ToResponse() (any, error) {
	return nil, nil
}

func DisconnectedMessage(owner string) IWSMessage {
	return &WSMessage{
		MessageType: DisconnectedType,
		Owner:       owner,
	}
}

func ConnectedMessage(owner string) IWSMessage {
	return &WSMessage{
		MessageType: ConnectedType,
		Owner:       owner,
	}
}
