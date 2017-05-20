
package web

import (
	"github.com/jchavannes/jgo/web"
	"github.com/jchavannes/iiproject/app/db/auth"
	"net/http"
	"github.com/jchavannes/iiproject/app/db/contact"
)

var contactList = web.Route{
	Pattern: URL_CONTACT_LIST,
	Handler: func(r *web.Response) {
		if ! auth.IsLoggedIn(r.Session.CookieId) {
			r.SetResponseCode(http.StatusUnauthorized)
			return
		}
		user, err := auth.GetSessionUser(r.Session.CookieId)
		if err != nil {
			r.Error(err, http.StatusUnprocessableEntity)
			return
		}
		contacts, err := contact.GetUserContacts(user.Id)
		if err != nil {
			r.Error(err, http.StatusInternalServerError)
			return
		}
		r.WriteJson(contacts, false)
	},
}
