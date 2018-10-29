package sound

import (
	"sync"

	"github.com/liuqi0826/seven/events"
)

var ResourceManagerInstance *ResourceManager

type ResourceManager struct {
	sync.Mutex
	events.EventDispatcher
}

func (this *ResourceManager) Setup() error {
	var err error

	this.EventDispatcher.EventDispatcher(this)

	return err
}
