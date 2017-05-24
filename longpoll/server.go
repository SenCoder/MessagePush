package poll

import (
	"coding.net/sencoder/Push/common"

	"github.com/lunny/tango"
)

type PollAction struct {
	tango.Ctx
	tango.Json
}

func (ctx *PollAction) Get() interface{} {

	return common.Message{Debug: "PollAction.Get"}
}

// start http server
func StartServer() {
	tg := tango.Classic()

	tg.Any("/poll/", new(PollAction))

	tg.Run(":8080")
}
