package web

import (
	"github.com/jchavannes/iiproject/app/db"
	"github.com/jchavannes/iiproject/app/db/key"
	"github.com/jchavannes/iiproject/eid/server"
	"github.com/jchavannes/jgo/web"
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
		userKey, err := key.Get(user.Id)
		if err != nil {
			r.Error(err, http.StatusInternalServerError)
			return
		}
		idResponse, err := server.ProcessIdRequest(r.Request.GetBody(), userKey.PublicKey)
		if err != nil {
			r.Error(err, http.StatusBadRequest)
			return
		}
		r.WriteJson(idResponse, false)
	},
}
