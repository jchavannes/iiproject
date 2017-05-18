package cmd

import (
	"github.com/jchavannes/iiproject/app/web"
)

func CmdWeb() error {
	return web.Run()
}
