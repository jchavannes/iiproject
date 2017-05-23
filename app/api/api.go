package api

import (
	"github.com/jchavannes/jgo/web"
)

const (
	PATTERN_ID = "/{username}/id"
	PATTERN_PROFILE = "/{username}/profile"
	PATTERN_MESSAGE = "/{username}/message"
)

func Run() error {
	server := web.Server{
		Port: 8253,
		Routes: []web.Route{
			idRoute,
			profileRoute,
			messageRoute,
		},
	}
	return server.Run()
}
