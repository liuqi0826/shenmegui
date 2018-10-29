package global

import (
	"fmt"
	"syscall/js"

	"github.com/liuqi0826/seven/events"
)

var Config *Parameter
var Dispatch *events.EventDispatcher
var Window js.Value
var Document js.Value

func init() {
	Dispatch = new(events.EventDispatcher)
	Dispatch.EventDispatcher(nil)

	Window = js.Global().Get("window")
	Document = js.Global().Get("document")

	Config = new(Parameter)
	Config.WindowWidth = int32(Document.Get("body").Get("clientWidth").Int())
	Config.WindowHeight = int32(Document.Get("documentElement").Get("clientHeight").Int())
	fmt.Println(Config.WindowWidth, Config.WindowHeight)
}

//++++++++++++++++++++ Parameter ++++++++++++++++++++

type Parameter struct {
	WindowWidth  int32
	WindowHeight int32
}
