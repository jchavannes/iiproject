package web

import (
	"github.com/jchavannes/jgo/web"
	"github.com/jchavannes/iiproject/app/db/user"
	"net/http"
)

var logoutRoute = web.Route{
	Pattern: URL_LOGOUT,
	Handler: func(r *web.Response) {
		if user.IsLoggedIn(r.Session.CookieId) {
			err := user.Logout(r.Session.CookieId)
			if err != nil {
				r.Error(err, http.StatusInternalServerError)
				return
			}
		}
		r.SetRedirect(getUrlWithBaseUrl(URL_INDEX, r))
	},
}
