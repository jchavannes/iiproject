package web

import (
	"github.com/jchavannes/iiproject/app/db"
	"github.com/jchavannes/iiproject/app/db/key"
	"github.com/jchavannes/iiproject/eid/api"
	"github.com/jchavannes/jgo/web"
	"encoding/json"
	"net/http"
)

var userIdRoute = web.Route{
	Pattern: "/u/{username}/id",
	Handler: func(r *web.Response) {
		username := r.Request.GetUrlNamedQueryVariable("username")
		user, _ := db.GetUserByUsername(username)
		if user == nil {
			r.SetResponseCode(http.StatusUnprocessableEntity)
			return
		}
		body := r.Request.GetBody()
		var idRequest api.IdRequest
		err := json.Unmarshal(body, &idRequest)
		if err != nil {
			r.Error(err, http.StatusBadRequest)
			return
		}
		switch idRequest.Name {
		case "/get":
			userKey, err := key.Get(user.Id)
			if err != nil {
				r.Error(err, http.StatusInternalServerError)
				return
			}
			resp := api.IdGetResponse{
				PublicKey: string(userKey.PublicKey),
			}
			r.WriteJson(resp, false)
		}
	},
}
