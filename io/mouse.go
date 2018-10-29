package io

import (
	"syscall/js"

	"github.com/liuqi0826/seven/events"
	"github.com/liuqi0826/shenmegui/global"
)

var MouseInstance *MouseManeger

func init() {
	MouseInstance = new(MouseManeger)
	MouseInstance.MouseManeger()
}

type MouseManeger struct {
	events.EventDispatcher

	close chan int
}

func (this *MouseManeger) MouseManeger() {
	this.EventDispatcher.EventDispatcher(this)
	this.close = make(chan int)
	go func() {
		mouseDownEvt := js.NewEventCallback(0, func(e js.Value) {
			data := new(events.MouseEventData)
			data.AltKey = e.Get("altKey").Bool()
			data.CtrlKey = e.Get("ctrlKey").Bool()
			data.ShiftKey = e.Get("shiftKey").Bool()
			data.X = float32(e.Get("pageX").Float())
			data.Y = float32(e.Get("pageY").Float())
			evt := new(events.Event)
			if e.Get("button").Int() == 2 {
				evt.Type = events.RIGHT_MOUSE_DOWN
			} else {
				evt.Type = events.MOUSE_DOWN
			}
			evt.Data = data
			this.EventDispatcher.DispatchEvent(evt)
		})
		defer mouseDownEvt.Release()

		mouseUpEvt := js.NewEventCallback(0, func(e js.Value) {
			data := new(events.MouseEventData)
			data.AltKey = e.Get("altKey").Bool()
			data.CtrlKey = e.Get("ctrlKey").Bool()
			data.ShiftKey = e.Get("shiftKey").Bool()
			data.X = float32(e.Get("pageX").Float())
			data.Y = float32(e.Get("pageY").Float())
			evt := new(events.Event)
			if e.Get("button").Int() == 2 {
				evt.Type = events.RIGHT_MOUSE_UP
			} else {
				evt.Type = events.MOUSE_UP
			}
			evt.Data = data
			this.EventDispatcher.DispatchEvent(evt)
		})
		defer mouseUpEvt.Release()

		mouseClickEvt := js.NewEventCallback(0, func(e js.Value) {
			data := new(events.MouseEventData)
			data.AltKey = e.Get("altKey").Bool()
			data.CtrlKey = e.Get("ctrlKey").Bool()
			data.ShiftKey = e.Get("shiftKey").Bool()
			data.X = float32(e.Get("pageX").Float())
			data.Y = float32(e.Get("pageY").Float())
			evt := new(events.Event)
			evt.Type = events.CLICK
			evt.Data = data
			this.EventDispatcher.DispatchEvent(evt)
		})
		defer mouseClickEvt.Release()

		mouseDBClickEvt := js.NewEventCallback(0, func(e js.Value) {
			data := new(events.MouseEventData)
			data.AltKey = e.Get("altKey").Bool()
			data.CtrlKey = e.Get("ctrlKey").Bool()
			data.ShiftKey = e.Get("shiftKey").Bool()
			data.X = float32(e.Get("pageX").Float())
			data.Y = float32(e.Get("pageY").Float())
			evt := new(events.Event)
			evt.Type = events.DOUBLE_CLICK
			evt.Data = data
			this.EventDispatcher.DispatchEvent(evt)
		})
		defer mouseDBClickEvt.Release()

		mouseMoveEvt := js.NewEventCallback(0, func(e js.Value) {
			data := new(events.MouseEventData)
			data.AltKey = e.Get("altKey").Bool()
			data.CtrlKey = e.Get("ctrlKey").Bool()
			data.ShiftKey = e.Get("shiftKey").Bool()
			data.X = float32(e.Get("pageX").Float())
			data.Y = float32(e.Get("pageY").Float())
			data.MovementX = float32(e.Get("movementX").Float())
			data.MovementY = float32(e.Get("movementY").Float())
			evt := new(events.Event)
			evt.Type = events.MOUSE_MOVE
			evt.Data = data
			this.EventDispatcher.DispatchEvent(evt)
		})

		global.Document.Call("addEventListener", "mousemove", mouseMoveEvt)
		global.Document.Call("addEventListener", "mousedown", mouseDownEvt)
		global.Document.Call("addEventListener", "mouseup", mouseUpEvt)
		global.Document.Call("addEventListener", "click", mouseClickEvt)
		global.Document.Call("addEventListener", "dblclick", mouseDBClickEvt)

		<-this.close
	}()
}
