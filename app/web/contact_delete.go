package web

import (
	"github.com/jchavannes/jgo/web"
	"github.com/jchavannes/iiproject/app/db/auth"
	"net/http"
	"github.com/jchavannes/iiproject/app/db/contact"
	"strconv"
)

var contactDeleteSubmit = web.Route{
	Pattern: URL_CONTACT_DELETE_SUBMIT,
	CsrfProtect: true,
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
		contactIdString := r.Request.GetFormValue("contactId")
		contactId, err := strconv.ParseInt(contactIdString, 10, 32)
		if err != nil {
			r.Error(err, http.StatusUnprocessableEntity)
			return
		}
		err = contact.DeleteUserContact(user.Id, uint(contactId))
		if err != nil {
			r.Error(err, http.StatusInternalServerError)
			return
		}
	},
}
