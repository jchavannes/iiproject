package eid

type ProfileRequest struct {
	Name string `json:"name"`
	Eid  string `json:"eid"`
}

type ProfileGetResponse struct {
	Profile string `json:"profile"`
}
