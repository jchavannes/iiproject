package web

import (
	"github.com/jchavannes/jgo/web"
)

var viewIdSubmitRoute = web.Route{
	Pattern: URL_VIEW_ID_SUBMIT,
	CsrfProtect: true,
	Handler: func(r *web.Response) {
		r.Write(r.Request.GetFormValue("id"))
	},
}
