package web

import (
	"github.com/jchavannes/jgo/web"
	"github.com/jchavannes/iiproject/app/db/auth"
	"net/http"
)

var logoutRoute = web.Route{
	Pattern: URL_LOGOUT,
	Handler: func(r *web.Response) {
		if auth.IsLoggedIn(r.Session.CookieId) {
			err := auth.Logout(r.Session.CookieId)
			if err != nil {
				r.Error(err, http.StatusInternalServerError)
				return
			}
		}
		r.SetRedirect(getUrlWithBaseUrl(URL_INDEX, r))
	},
}
