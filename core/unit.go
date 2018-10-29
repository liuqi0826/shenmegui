package core

import (
	"github.com/liuqi0826/seven/events"
)

type Unit struct {
	events.EventDispatcher

	id   uint32
	name string
}

func (this *Unit) Unit() {
	this.EventDispatcher.EventDispatcher(this)
	this.id = GetNextInstanceID()
}
func (this *Unit) GetID() uint32 {
	return this.id
}
func (this *Unit) GetName() string {
	return this.name
}
func (this *Unit) SetName(value string) {
	this.name = value
}
