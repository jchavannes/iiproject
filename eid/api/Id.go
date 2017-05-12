package api

type IdRequest struct {
	Name string `json:"name"`
}

type IdGetResponse struct {
	PublicKey string `json:"public_key"`
}
