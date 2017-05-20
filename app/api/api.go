package api

import (
	"github.com/jchavannes/jgo/web"
)

const (
	PATTERN_USER_ID = "/u/{username}/id"
	PATTERN_USER_PROFILE = "/u/{username}/profile"
)

func Run() error {

	server := web.Server{
		Port: 8253,
		Routes: []web.Route{
			userIdRoute,
			userProfileRoute,
		},
	}
	return server.Run()
}
