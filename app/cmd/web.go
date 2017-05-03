package cmd

import (
	"github.com/jchavannes/iiproject/app/auth"
	"github.com/jchavannes/jgo/web"
	"net/http"
	"github.com/jchavannes/iiproject/app/profile"
	"github.com/jchavannes/iiproject/app/db"
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
	indexRoute = web.Route{
		Pattern: URL_INDEX,
		Handler: func(r *web.Response) {
			if auth.IsLoggedIn(r.Session.CookieId) {
				r.SetRedirect(getUrlWithBaseUrl(URL_DASHBOARD, r))
				return
			}
			r.Render()
		},
	}

	dashboardRoute = web.Route{
		Pattern: URL_DASHBOARD,
		Handler: func(r *web.Response) {
			if auth.IsLoggedIn(r.Session.CookieId) {
				r.Render()
			} else {
				r.SetRedirect(getUrlWithBaseUrl(URL_INDEX, r))
			}
		},
	}

	editProfileRoute = web.Route{
		Pattern: URL_EDIT_PROFILE,
		Handler: func(r *web.Response) {
			user := auth.GetSessionUser(r.Session.CookieId)
			if user == nil {
				r.SetResponseCode(http.StatusUnauthorized)
				return
			}
			profileString, _ := profile.Get(user.Id)
			r.Helper["Profile"] = profileString
			r.Render()
		},
	}

	editProfileSubmitRoute = web.Route{
		Pattern: URL_EDIT_PROFILE_SUBMIT,
		CsrfProtect: true,
		Handler: func(r *web.Response) {
			user := auth.GetSessionUser(r.Session.CookieId)
			if user == nil {
				r.SetResponseCode(http.StatusUnauthorized)
				return
			}
			profileString := r.Request.GetFormValue("profile")
			err := profile.Edit(user.Id, profileString)
			if err != nil {
				r.SetResponseCode(http.StatusInternalServerError)
			}
		},
	}

	viewIdSubmitRoute = web.Route{
		Pattern: URL_VIEW_ID_SUBMIT,
		CsrfProtect: true,
		Handler: func(r *web.Response) {
			r.Write(r.Request.GetFormValue("id"))
		},
	}

	signupRoute = web.Route{
		Pattern: URL_SIGNUP,
		Handler: func(r *web.Response) {
			if auth.IsLoggedIn(r.Session.CookieId) {
				r.SetRedirect(getUrlWithBaseUrl(URL_DASHBOARD, r))
				return
			}
			r.Render()
		},
	}

	signupSubmitRoute = web.Route{
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
				r.SetResponseCode(http.StatusUnauthorized)
				r.Write(err.Error())
			}
		},
	}

	loginRoute = web.Route{
		Pattern: URL_LOGIN,
		Handler: func(r *web.Response) {
			if auth.IsLoggedIn(r.Session.CookieId) {
				r.SetRedirect(getUrlWithBaseUrl(URL_DASHBOARD, r))
				return
			}
			r.Render()
		},
	}

	loginSubmitRoute = web.Route{
		Pattern: URL_LOGIN_SUBMIT,
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

			err := auth.Login(r.Session.CookieId, username, password)
			if err != nil {
				r.SetResponseCode(http.StatusUnauthorized)
				r.Write(err.Error())
			}
		},
	}

	logoutRoute = web.Route{
		Pattern: URL_LOGOUT,
		Handler: func(r *web.Response) {
			if auth.IsLoggedIn(r.Session.CookieId) {
				err := auth.Logout(r.Session.CookieId)
				if err != nil {
					r.SetResponseCode(http.StatusInternalServerError)
					r.Write(err.Error())
					return
				}
			}
			r.SetRedirect(getUrlWithBaseUrl(URL_INDEX, r))
		},
	}

	userProfileRoute = web.Route{
		Pattern: "/u/{username}/profile",
		Handler: func(r *web.Response) {
			username := r.Request.GetUrlNamedQueryVariable("username")
			user, _ := db.GetUserByUsername(username)
			if user == nil {
				r.SetResponseCode(http.StatusUnprocessableEntity)
				return
			}
			profileString, _ := profile.Get(user.Id)
			r.Write(profileString)
		},
	}

	notFoundHandler = func(r *web.Response) {
		r.SetResponseCode(http.StatusNotFound)
		r.RenderTemplate("404")
	}

	preHandler = func(r *web.Response) {
		baseUrl := getBaseUrl(r)
		r.Helper["BaseUrl"] = baseUrl
		if auth.IsLoggedIn(r.Session.CookieId) {
			r.Helper["Username"] = auth.GetSessionUser(r.Session.CookieId).Username
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

func runWeb() error {
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
		},
		StaticFilesDir: "pub",
		TemplatesDir: "templates",
		UseSessions: true,
	}
	return server.Run()
}
