package display

import (
	"github.com/liuqi0826/seven/events"
	"github.com/liuqi0826/shenmegui/global"
)

//++++++++++++++++++++ Stage ++++++++++++++++++++

type Stage struct {
	events.EventDispatcher

	Viewport *Viewport
}

func (this *Stage) AddChild(displayObject IDisplayObject) {
	this.Viewport.Scene.AddChild(displayObject)
}
func (this *Stage) RemoveChild(displayObject IDisplayObject) IDisplayObject {
	return this.Viewport.Scene.RemoveChild(displayObject)
}
func (this *Stage) StageWidth() int32 {
	return int32(global.Config.WindowWidth)
}
func (this *Stage) StageHeight() int32 {
	return int32(global.Config.WindowHeight)
}
func (this *Stage) Setup() {
	this.EventDispatcher.EventDispatcher(this)
	this.Viewport = new(Viewport)
	this.Viewport.Viewport(uint32(global.Config.WindowWidth), uint32(global.Config.WindowHeight), FORWARD)

	event := new(events.Event)
	event.Type = events.INIT
	this.DispatchEvent(event)
}
func (this *Stage) Frame() {
	event := new(events.Event)
	event.Type = events.ENTER_FRAME
	this.DispatchEvent(event)

	this.Viewport.Frame()
}
