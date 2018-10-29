package sound

import (
	"syscall/js"

	"github.com/liuqi0826/seven/events"
	"github.com/liuqi0826/shenmegui/core"
	"github.com/liuqi0826/shenmegui/global"
)

type Sound struct {
	core.Unit
	events.EventDispatcher

	Audio js.Value
}

func (this *Sound) Sound() {
	this.Unit.Unit()
	this.Audio = global.Document.Call("createElement", "audio")
}
func (this *Sound) Load(url string, autoplay bool, loop bool) {
	this.Audio.Set("src", url)
	if autoplay {
		this.Audio.Set("autoplay", "autoplay")
	}
	if loop {
		this.Audio.Set("loop", "loop")
	}
}
func (this *Sound) Play() {
	this.Audio.Call("play")
}
func (this *Sound) IsPlay() bool {
	return this.Audio.Get("play").Bool()
}
func (this *Sound) Pause() {
	this.Audio.Call("pause")
}
func (this *Sound) IsPause() bool {
	return this.Audio.Get("pause").Bool()
}
func (this *Sound) SetVolume(value float32) {
	this.Audio.Set("volume", value)
}
func (this *Sound) CurrentTime() int {
	return this.Audio.Get("currentTime").Int()
}
func (this *Sound) Duration() int {
	return this.Audio.Get("duration").Int()
}
