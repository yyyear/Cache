package Cache

import (
	"sync"
	"time"
)

type Cache struct {
	syncMap sync.Map
}

type cacheItem struct {
	value   interface{} // 值
	expired time.Time   // 过期时间
}

// Set 指针类型的方法接受者，才可以保证在方法内部访问和修改对象的数据。
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.syncMap.Store(key, cacheItem{
		value:   value,
		expired: time.Now().Add(duration),
	})
}

func (c *Cache) Get(key string) (interface{}, bool) {
	item, ok := c.syncMap.Load(key)
	if !ok {
		return nil, false
	}

	cacheItem := item.(cacheItem)
	if time.Now().After(cacheItem.expired) { // 判断是否过期
		c.syncMap.Delete(key)
		return nil, false
	}

	return cacheItem.value, true
}

// GetExpired 获取缓存的过期时间
func (c *Cache) GetExpired(key string) (time.Time, bool) {
	item, ok := c.syncMap.Load(key)
	if !ok {
		return time.Time{}, false
	}

	cacheItem := item.(cacheItem)
	return cacheItem.expired, true
}
