package wsservice

import "net/http"

type Payload struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Header  http.Header
}
