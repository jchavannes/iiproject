package web

import (
	"github.com/jchavannes/jgo/web"
	"github.com/jchavannes/iiproject/app/db/auth"
)

var indexRoute = web.Route{
	Pattern: URL_INDEX,
	Handler: func(r *web.Response) {
		if auth.IsLoggedIn(r.Session.CookieId) {
			r.SetRedirect(getUrlWithBaseUrl(URL_DASHBOARD, r))
			return
		}
		r.Render()
	},
}
