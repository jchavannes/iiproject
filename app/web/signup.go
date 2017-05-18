package web

import (
	"github.com/jchavannes/jgo/web"
	"github.com/jchavannes/iiproject/app/db/auth"
	"net/http"
)

var signupRoute = web.Route{
	Pattern: URL_SIGNUP,
	Handler: func(r *web.Response) {
		if auth.IsLoggedIn(r.Session.CookieId) {
			r.SetRedirect(getUrlWithBaseUrl(URL_DASHBOARD, r))
			return
		}
		r.Render()
	},
}

var signupSubmitRoute = web.Route{
	Pattern: URL_SIGNUP_SUBMIT,
	CsrfProtect: true,
	Handler: func(r *web.Response) {
		if auth.IsLoggedIn(r.Session.CookieId) {
			r.SetRedirect(getUrlWithBaseUrl(URL_DASHBOARD, r))
			return
		}
		// Protects against some session hi-jacking attacks
		r.ResetOrCreateSession()
		username := r.Request.GetFormValue("username")
		password := r.Request.GetFormValue("password")

		err := auth.Signup(r.Session.CookieId, username, password)
		if err != nil {
			r.Error(err, http.StatusUnauthorized)
		}
	},
}
