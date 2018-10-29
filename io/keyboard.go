package io

import (
	"syscall/js"

	"github.com/liuqi0826/seven/events"
	"github.com/liuqi0826/shenmegui/global"
)

var KeyboardInstance *KeyboardManeger

func init() {
	KeyboardInstance = new(KeyboardManeger)
	KeyboardInstance.KeyboardManeger()
}

type KeyboardManeger struct {
	events.EventDispatcher

	close chan int
}

func (this *KeyboardManeger) KeyboardManeger() {
	this.EventDispatcher.EventDispatcher(this)
	this.close = make(chan int)
	go func() {
		keyDownEvt := js.NewEventCallback(js.PreventDefault, func(e js.Value) {
			data := new(events.KeyboardEventData)
			data.AltKey = e.Get("altKey").Bool()
			data.CtrlKey = e.Get("ctrlKey").Bool()
			data.ShiftKey = e.Get("shiftKey").Bool()
			data.CharCode = uint32(e.Get("charCode").Int())
			data.KeyCode = uint32(e.Get("keyCode").Int())
			data.Key = e.Get("key").String()
			evt := new(events.Event)
			evt.Type = events.KEY_DOWN
			evt.Data = data
			this.EventDispatcher.DispatchEvent(evt)
		})
		defer keyDownEvt.Release()

		keyUpEvt := js.NewEventCallback(js.PreventDefault, func(e js.Value) {
			data := new(events.KeyboardEventData)
			data.AltKey = e.Get("altKey").Bool()
			data.CtrlKey = e.Get("ctrlKey").Bool()
			data.ShiftKey = e.Get("shiftKey").Bool()
			data.CharCode = uint32(e.Get("charCode").Int())
			data.KeyCode = uint32(e.Get("keyCode").Int())
			data.Key = e.Get("key").String()
			evt := new(events.Event)
			evt.Type = events.KEY_UP
			evt.Data = data
			this.EventDispatcher.DispatchEvent(evt)
		})
		defer keyUpEvt.Release()

		keyPressEvt := js.NewEventCallback(js.PreventDefault, func(e js.Value) {
			data := new(events.KeyboardEventData)
			data.AltKey = e.Get("altKey").Bool()
			data.CtrlKey = e.Get("ctrlKey").Bool()
			data.ShiftKey = e.Get("shiftKey").Bool()
			data.CharCode = uint32(e.Get("charCode").Int())
			data.KeyCode = uint32(e.Get("keyCode").Int())
			data.Key = e.Get("key").String()
			evt := new(events.Event)
			evt.Type = events.KEY_PRESS
			evt.Data = data
			this.EventDispatcher.DispatchEvent(evt)
		})
		defer keyPressEvt.Release()

		global.Document.Call("addEventListener", "keydown", keyDownEvt)
		global.Document.Call("addEventListener", "keyup", keyUpEvt)
		global.Document.Call("addEventListener", "keypress", keyPressEvt)

		<-this.close
	}()
}
