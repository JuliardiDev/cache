package src

import (
	"sync"
)

// Cache is entity
type Cache struct{}

var (
	hmap    map[string]map[string]string
	smember map[string][]string

	mux sync.Mutex
)

func init() {
	hmap = make(map[string]map[string]string)
}

// Hmset is to set hash map
func (c Cache) Hmset(key string, val map[string]string) {
	mux.Lock()
	hmap[key] = val
	mux.Unlock()
}

// Hgetall is to get value of hmset
func (c Cache) Hgetall(key string) map[string]string {
	return hmap[key]
}

// GetMap is to get all memory data
func (c Cache) GetMap() map[string]map[string]string {
	return hmap
}
