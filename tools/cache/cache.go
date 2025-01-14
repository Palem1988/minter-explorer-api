package cache

import (
	"encoding/json"
	"github.com/MinterTeam/minter-explorer-api/blocks"
	"github.com/MinterTeam/minter-explorer-api/helpers"
	"github.com/centrifugal/centrifuge-go"
	"sync"
	"time"
)

type ExplorerCache struct {
	lastBlockId uint64
	items       *sync.Map
}

// cache constructor
func NewCache() *ExplorerCache {
	cache := &ExplorerCache{
		lastBlockId: uint64(0),
		items:       new(sync.Map),
	}

	return cache
}

// create new cache item
func (c *ExplorerCache) newCacheItem(value interface{}, ttl interface{}) *CacheItem {
	switch t := ttl.(type) {
	case time.Duration:
		ttl := time.Now().Add(t * time.Second)
		return &CacheItem{value: value, ttl: &ttl}
	case int:
		ttl := c.lastBlockId + uint64(t)
		return &CacheItem{value: value, btl: &ttl}
	}

	panic("Invalid cache ttl type.")
}

// get or store value from cache
func (c *ExplorerCache) Get(key interface{}, callback func() interface{}, ttl interface{}) interface{} {
	v, ok := c.items.Load(key)
	if ok {
		item := v.(*CacheItem)
		if !item.IsExpired(c.lastBlockId) {
			return item.value
		}
	}

	return c.Store(key, callback(), ttl)
}

// save value to cache
func (c *ExplorerCache) Store(key interface{}, value interface{}, ttl interface{}) interface{} {
	c.items.Store(key, c.newCacheItem(value, ttl))
	return value
}

// loop for checking items expiration
func (c *ExplorerCache) ExpirationCheck() {
	c.items.Range(func(key, value interface{}) bool {
		item := value.(*CacheItem)
		if item.IsExpired(c.lastBlockId) {
			c.items.Delete(key)
		}

		return true
	})
}

// set new last block id
func (c *ExplorerCache) SetBlockId(id uint64) {
	c.lastBlockId = id
	// clean expired items
	go c.ExpirationCheck()
}

// update last block id by ws data
func (c *ExplorerCache) OnPublish(sub *centrifuge.Subscription, e centrifuge.PublishEvent) {
	var block blocks.Resource
	err := json.Unmarshal(e.Data, &block)
	helpers.CheckErr(err)

	// update last block id
	c.SetBlockId(block.ID)
}
