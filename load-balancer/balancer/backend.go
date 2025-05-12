package balancer

import (
	"net/url"
	"sync"
)

type Backend struct {
	URL   *url.URL
	Alive bool
	Mutex sync.RWMutex
}

func (b *Backend) SetAlive(alive bool) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	b.Alive = alive
}

func (b *Backend) IsAlive() bool {
	b.Mutex.RLock()
	defer b.Mutex.RUnlock()
	return b.Alive
}
