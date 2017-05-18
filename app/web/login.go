package web

import (
	"github.com/jchavannes/jgo/web"
	"github.com/jchavannes/iiproject/app/db/user"
	"net/http"
)

var loginRoute = web.Route{
	Pattern: URL_LOGIN,
	Handler: func(r *web.Response) {
		if user.IsLoggedIn(r.Session.CookieId) {
			r.SetRedirect(getUrlWithBaseUrl(URL_DASHBOARD, r))
			return
		}
		r.Render()
	},
}

var loginSubmitRoute = web.Route{
	Pattern: URL_LOGIN_SUBMIT,
	CsrfProtect: true,
	Handler: func(r *web.Response) {
		if user.IsLoggedIn(r.Session.CookieId) {
			r.SetRedirect(getUrlWithBaseUrl(URL_DASHBOARD, r))
			return
		}
		// Protects against some session hi-jacking attacks
		r.ResetOrCreateSession()
		username := r.Request.GetFormValue("username")
		password := r.Request.GetFormValue("password")

		err := user.Login(r.Session.CookieId, username, password)
		if err != nil {
			r.Error(err, http.StatusUnauthorized)
		}
	},
}
