package cache

import "sync"

type cacheSlice struct {
	items map[int]int
	sync.RWMutex
}

func NewCacheSlice(size int) *cacheSlice {
	return &cacheSlice{
		items: make(map[int]int),
	}
}

func (c *cacheSlice) Get(key int) (int, bool) {
	c.RLock()
	v, found := c.items[key]
	c.RUnlock()
	return v, found
}

func (c *cacheSlice) Set(key int, value int) {
	c.Lock()
	c.items[key] = value
	c.Unlock()
}

func (c *cacheSlice) Delete(key int) {
	c.Lock()
	delete(c.items, key)
	c.Unlock()
}

func (c *cacheSlice) Incr(key int) {
	v, found := c.items[key]
	if found {
		c.items[key] = v + 1
	} else {
		c.items[key] = 1
	}
	c.Unlock()
}
