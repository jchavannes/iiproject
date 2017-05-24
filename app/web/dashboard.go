package web

import (
	"github.com/jchavannes/jgo/web"
	"github.com/jchavannes/iiproject/app/db/auth"
)

var dashboardRoute = web.Route{
	Pattern: URL_DASHBOARD,
	Handler: func(r *web.Response) {
		if ! auth.IsLoggedIn(r.Session.CookieId) {
			r.SetRedirect(getUrlWithBaseUrl(URL_INDEX, r))
			return
		}
		r.Render()
	},
}
