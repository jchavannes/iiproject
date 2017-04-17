package cmd

import (
	"github.com/jchavannes/iiproject/auth"
	"github.com/jchavannes/jgo/web"
	"net/http"
)

var (
	indexRoute = web.Route{
		Pattern: "/",
		Handler: func(r *web.Response) {
			r.Render()
		},
	}

	signupRoute = web.Route{
		Pattern: "/signup",
		Handler: func(r *web.Response) {
			r.Render()
		},
	}

	signupSubmitRoute = web.Route{
		Pattern: "/signup-submit",
		CsrfProtect: true,
		Handler: func(r *web.Response) {
			username := r.Request.GetFormValue("username")
			password := r.Request.GetFormValue("password")

			err := auth.Signup(r.Session.CookieId, username, password)

			if err != nil {
				r.Write(err.Error())
				r.SetResponseCode(http.StatusConflict)
				return
			}
		},
	}

	notFoundHandler = func(r *web.Response) {
		r.SetResponseCode(http.StatusNotFound)
		r.RenderTemplate("404")
	}

	preHandler = func(r *web.Response) {
		baseUrl := r.Request.GetHeader("AppPath")
		if baseUrl == "" {
			baseUrl = "/"
		}
		r.Helper["BaseUrl"] = baseUrl
	}
)

func runWeb() error {
	server := web.Server{
		NotFoundHandler: notFoundHandler,
		Port: 8252,
		PreHandler: preHandler,
		Routes: []web.Route{
			indexRoute,
			signupRoute,
			signupSubmitRoute,
		},
		StaticFilesDir: "pub",
		TemplatesDir: "templates",
		UseSessions: true,
	}
	return server.Run()
}
