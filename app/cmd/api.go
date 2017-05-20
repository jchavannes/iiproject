package cmd

import (
	"github.com/jchavannes/iiproject/app/api"
)

func CmdApi() error {
	return api.Run()
}
