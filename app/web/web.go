package web

import (
	"github.com/jchavannes/iiproject/app/db/user"
	"github.com/jchavannes/jgo/web"
	"net/http"
)

const (
	URL_INDEX = "/"
	URL_DASHBOARD = "/dashboard"
	URL_EDIT_PROFILE = "/edit-profile"
	URL_EDIT_PROFILE_SUBMIT = "/edit-profile-submit"
	URL_VIEW_ID_SUBMIT = "/view-id-submit"
	URL_SIGNUP = "/signup"
	URL_SIGNUP_SUBMIT = "/signup-submit"
	URL_LOGIN = "/login"
	URL_LOGIN_SUBMIT = "/login-submit"
	URL_LOGOUT = "/logout"
)

var (
	notFoundHandler = func(r *web.Response) {
		r.SetResponseCode(http.StatusNotFound)
		r.RenderTemplate("404")
	}

	preHandler = func(r *web.Response) {
		baseUrl := getBaseUrl(r)
		r.Helper["BaseUrl"] = baseUrl
		if user.IsLoggedIn(r.Session.CookieId) {
			r.Helper["Username"] = user.GetSessionUser(r.Session.CookieId).Username
		}
	}

	getBaseUrl = func(r *web.Response) string {
		baseUrl := r.Request.GetHeader("AppPath")
		if baseUrl == "" {
			baseUrl = "/"
		}
		return baseUrl
	}

	getUrlWithBaseUrl = func(url string, r *web.Response) string {
		baseUrl := getBaseUrl(r)
		baseUrl = baseUrl[:len(baseUrl) - 1]
		return baseUrl + url
	}
)

func Run() error {
	server := web.Server{
		NotFoundHandler: notFoundHandler,
		Port: 8252,
		PreHandler: preHandler,
		Routes: []web.Route{
			indexRoute,
			dashboardRoute,
			editProfileRoute,
			editProfileSubmitRoute,
			viewIdSubmitRoute,
			signupRoute,
			signupSubmitRoute,
			loginRoute,
			loginSubmitRoute,
			logoutRoute,
			userProfileRoute,
			userIdRoute,
		},
		StaticFilesDir: "pub",
		TemplatesDir: "templates",
		UseSessions: true,
	}
	return server.Run()
}
