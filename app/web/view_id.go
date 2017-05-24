package web

import (
	"github.com/jchavannes/jgo/web"
	"github.com/jchavannes/iiproject/eid/client"
	"github.com/jchavannes/iiproject/eid"
	"github.com/jchavannes/iiproject/app/db/auth"
	"net/http"
	"github.com/jchavannes/iiproject/app/db/key"
	"html"
)

var viewIdSubmitRoute = web.Route{
	Pattern: URL_VIEW_ID_SUBMIT,
	CsrfProtect: true,
	Handler: func(r *web.Response) {
		id := r.Request.GetFormValue("id")
		user, err := auth.GetSessionUser(r.Session.CookieId)
		if err != nil {
			r.Error(err, http.StatusUnprocessableEntity)
			return
		}
		userKey, err := key.Get(user.Id)
		if err != nil {
			r.Error(err, http.StatusUnprocessableEntity)
			return
		}
		profileResponse, err := client.GetProfile(id, r.Request.HttpRequest.Host + "/id/" + user.Username, eid.KeyPair{
			PublicKey: []byte(userKey.PublicKey),
			PrivateKey: []byte(userKey.PrivateKey),
		})
		if err != nil {
			r.Error(err, http.StatusUnprocessableEntity)
			return
		}
		r.Write(html.EscapeString(profileResponse.Body))
	},
}
