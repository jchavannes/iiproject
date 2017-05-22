package web

import (
	"github.com/jchavannes/jgo/web"
	"github.com/jchavannes/iiproject/app/db/auth"
)

var messagesRoute = web.Route{
	Pattern: URL_MESSAGES,
	Handler: func(r *web.Response) {
		if auth.IsLoggedIn(r.Session.CookieId) {
			r.Render()
		} else {
			r.SetRedirect(getUrlWithBaseUrl(URL_INDEX, r))
		}
	},
}
