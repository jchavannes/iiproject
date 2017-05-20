package web

import (
	"github.com/jchavannes/jgo/web"
	"github.com/jchavannes/iiproject/app/db/auth"
	"net/http"
	"github.com/jchavannes/iiproject/eid/client"
	"github.com/jchavannes/iiproject/app/db/contact"
)

var contactAddSubmit = web.Route{
	Pattern: URL_CONTACT_ADD_SUBMIT,
	CsrfProtect: true,
	Handler: func(r *web.Response) {
		if ! auth.IsLoggedIn(r.Session.CookieId) {
			r.SetResponseCode(http.StatusUnauthorized)
			return
		}
		id := r.Request.GetFormValue("id")
		user, err := auth.GetSessionUser(r.Session.CookieId)
		if err != nil {
			r.Error(err, http.StatusUnprocessableEntity)
			return
		}
		idGetResponse, err := client.GetId(id)
		if err != nil {
			r.Error(err, http.StatusInternalServerError)
			return
		}
		err = contact.AddContact(id, idGetResponse.PublicKey, user.Id)
		if err != nil {
			r.Error(err, http.StatusInternalServerError)
			return
		}
	},
}
