package web

import (
	"github.com/jchavannes/jgo/web"
	"github.com/jchavannes/iiproject/app/db/user"
	"github.com/jchavannes/iiproject/app/db/profile"
	"net/http"
)

var editProfileRoute = web.Route{
	Pattern: URL_EDIT_PROFILE,
	Handler: func(r *web.Response) {
		user := user.GetSessionUser(r.Session.CookieId)
		if user == nil {
			r.SetResponseCode(http.StatusUnauthorized)
			return
		}
		profileString, _ := profile.Get(user.Id)
		r.Helper["Profile"] = profileString
		r.Render()
	},
}

var editProfileSubmitRoute = web.Route{
	Pattern: URL_EDIT_PROFILE_SUBMIT,
	CsrfProtect: true,
	Handler: func(r *web.Response) {
		user := user.GetSessionUser(r.Session.CookieId)
		if user == nil {
			r.SetResponseCode(http.StatusUnauthorized)
			return
		}
		profileString := r.Request.GetFormValue("profile")
		err := profile.Edit(user.Id, profileString)
		if err != nil {
			r.Error(err, http.StatusInternalServerError)
		}
	},
}
