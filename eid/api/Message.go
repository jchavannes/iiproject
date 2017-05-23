package api

import "time"

type MessageSend struct {
	Name     string `json:"name"`
	Eid      string `json:"eid"`
	Message  string `json:"message"`
	SendTime time.Time `json:"send_time"`
}

type MessageSendResponse struct {
	Acknowledged bool `json:"acknowledged"`
}
