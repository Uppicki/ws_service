package wsservicemessage

type WSMessageType string

const (
	DisconnectedType WSMessageType = "Disconnected"
	ConnectedType    WSMessageType = "Connected"
)
