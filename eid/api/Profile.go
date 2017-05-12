package api

type ProfileRequest struct {
	Name string `json:"name"`
	Eid  string `json:"eid"`
}

type ProfileGetResponse struct {
	Body string `json:"body"`
}
