package web

import (
	"encoding/json"
	"github.com/jchavannes/iiproject/app/db"
	"github.com/jchavannes/iiproject/app/db/key"
	"github.com/jchavannes/iiproject/app/db/profile"
	"github.com/jchavannes/iiproject/eid/api"
	"github.com/jchavannes/go-pgp/pgp"
	"github.com/jchavannes/iiproject/eid/client"
	"github.com/jchavannes/jgo/web"
	"net/http"
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

		body := r.Request.GetBody()
		var profileRequest api.ProfileRequest
		err := json.Unmarshal(body, &profileRequest)
		if err != nil {
			r.Error(err, http.StatusBadRequest)
			return
		}

		switch profileRequest.Name {
		case "/get":
			userKey, err := key.Get(user.Id)
			if err != nil {
				r.Error(err, http.StatusInternalServerError)
				return
			}

			profileString, _ := profile.Get(user.Id)
			profileGetResponse := api.ProfileGetResponse{
				Body: profileString,
			}

			jsonResponse, err := json.Marshal(profileGetResponse)
			if err != nil {
				r.Error(err, http.StatusInternalServerError)
				return
			}

			privEntity, err := pgp.GetEntity(userKey.PublicKey, userKey.PrivateKey)
			if err != nil {
				r.Error(err, http.StatusInternalServerError)
				return
			}

			signature, err := pgp.Sign(privEntity, jsonResponse)
			if err != nil {
				r.Error(err, http.StatusInternalServerError)
				return
			}

			idGetResponse, err := client.GetId(profileRequest.Eid)
			if err != nil {
				r.Error(err, http.StatusInternalServerError)
				return
			}

			pubEntity, err := pgp.GetEntity([]byte(idGetResponse.PublicKey), nil)
			if err != nil {
				r.Error(err, http.StatusInternalServerError)
				return
			}

			message := append(jsonResponse, signature...)
			encrypted, err := pgp.Encrypt(pubEntity, message)
			if err != nil {
				r.Error(err, http.StatusInternalServerError)
				return
			}

			r.Write(string(encrypted))
		}
	},
}
