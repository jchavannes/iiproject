package web

import (
	"github.com/jchavannes/iiproject/app/db"
	"github.com/jchavannes/iiproject/app/db/key"
	"github.com/jchavannes/iiproject/app/db/profile"
	"github.com/jchavannes/iiproject/eid/server"
	"github.com/jchavannes/jgo/web"
	"net/http"
	"strings"
)

var userProfileRoute = web.Route{
	Pattern: "/u/{username}/profile",
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

		profileString, err := profile.Get(user.Id)
		if err != nil {
			r.Error(err, http.StatusInternalServerError)
		}

		profileResponse, err := server.ProcessProfileRequest(
			r.Request.GetBody(),
			userKey.PublicKey,
			userKey.PrivateKey,
			strings.NewReader(profileString),
		)

		if err != nil {
			r.Error(err, http.StatusBadRequest)
			return
		}
		r.Write(string(profileResponse))
	},
}
