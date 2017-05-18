package web

import (
	"github.com/jchavannes/jgo/web"
	"github.com/jchavannes/iiproject/app/db/user"
)

var dashboardRoute = web.Route{
	Pattern: URL_DASHBOARD,
	Handler: func(r *web.Response) {
		if user.IsLoggedIn(r.Session.CookieId) {
			r.Render()
		} else {
			r.SetRedirect(getUrlWithBaseUrl(URL_INDEX, r))
		}
	},
}
