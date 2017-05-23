package api

import (
	"github.com/jchavannes/iiproject/app/db"
	"github.com/jchavannes/iiproject/app/db/key"
	"github.com/jchavannes/iiproject/app/db/message"
	"github.com/jchavannes/iiproject/eid/server"
	"github.com/jchavannes/jgo/web"
	"net/http"
)

var messageRoute = web.Route{
	Pattern: PATTERN_MESSAGE,
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

		messageSend, messageResponse, err := server.ProcessMessageRequest(
			r.Request.GetBody(),
			userKey.GetKeyPair(),
		)
		if err != nil {
			r.Error(err, http.StatusInternalServerError)
			return
		}

		err = message.Add(user.Id, messageSend.Eid, messageSend, false)
		if err != nil {
			r.Error(err, http.StatusInternalServerError)
			return
		}

		r.Write(string(messageResponse))
	},
}
