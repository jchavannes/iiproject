package cmd

import (
	"github.com/jchavannes/jgo/web"
	"net/http"
)

var (
	indexRoute = web.Route{
		Pattern: "/",
		Handler: func (r *web.Response) {
			r.Render()
		},
	}

	notFoundHandler = func (r *web.Response) {
		r.SetResponseCode(http.StatusNotFound)
		r.RenderTemplate("404")
	}
)

func Web() error {
	server := web.Server{
		NotFoundHandler: notFoundHandler,
		Port: 8252,
		Routes: []web.Route{
			indexRoute,
		},
		StaticFilesDir: "pub",
		TemplatesDir: "templates",
	}
	return server.Run()
}
