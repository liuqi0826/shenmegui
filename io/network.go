package io

import (
	"syscall/js"

	"github.com/liuqi0826/seven/events"
)

type WebSocket struct {
	events.EventDispatcher

	close     chan int
	websocket js.Value
}

func (this *WebSocket) WebSocket(addr string) {
	this.EventDispatcher.EventDispatcher(this)
	this.close = make(chan int)
	go func() {
		openEvt := js.NewEventCallback(0, func(e js.Value) {
			evt := new(events.Event)
			evt.Type = events.OPEN
			this.DispatchEvent(evt)
		})
		defer openEvt.Release()
		closeEvt := js.NewEventCallback(0, func(e js.Value) {
			evt := new(events.Event)
			evt.Type = events.CLOSE
			this.DispatchEvent(evt)
		})
		defer closeEvt.Release()
		messageEvt := js.NewEventCallback(0, func(e js.Value) {
			evt := new(events.Event)
			evt.Type = events.MESSAGE
			evt.Data = e.Get("data")
			this.DispatchEvent(evt)
		})
		defer messageEvt.Release()
		errorEvt := js.NewEventCallback(0, func(e js.Value) {
			evt := new(events.Event)
			evt.Type = events.ERROR
			evt.Data = e.Get("data")
			this.DispatchEvent(evt)
		})
		defer errorEvt.Release()
		this.websocket.Call("addEventListener", "onopen", openEvt)
		this.websocket.Call("addEventListener", "onclose", closeEvt)
		this.websocket.Call("addEventListener", "onmessage", messageEvt)
		this.websocket.Call("addEventListener", "onerror", errorEvt)
		<-this.close
	}()
	this.websocket = js.Global().Get("WebSocket").New(addr)
}
func (this *WebSocket) Send(value []byte) {
	valueTypeArray := js.TypedArrayOf(value)
	this.websocket.Call("send", valueTypeArray)
	valueTypeArray.Release()
}
func (this *WebSocket) Close() {
	this.websocket.Call("close")
	this.close <- 0
}
func (this *WebSocket) ReadyState() int {
	return this.websocket.Get("readyState").Int()
}
func (this *WebSocket) BufferedAmount() int {
	return this.websocket.Get("bufferedAmount").Int()
}
