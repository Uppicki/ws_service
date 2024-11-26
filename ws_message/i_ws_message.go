package wsservicemessage



type IWSMessage interface {
	Map()
	ToResponse() (any, error)
}
