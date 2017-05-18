package web

import (
	"github.com/jchavannes/jgo/web"
	"github.com/jchavannes/iiproject/app/db/user"
)

var indexRoute = web.Route{
	Pattern: URL_INDEX,
	Handler: func(r *web.Response) {
		if user.IsLoggedIn(r.Session.CookieId) {
			r.SetRedirect(getUrlWithBaseUrl(URL_DASHBOARD, r))
			return
		}
		r.Render()
	},
}
