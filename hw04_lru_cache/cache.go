package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheElement struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	var elementExists bool

	if _, exists := c.items[key]; exists {
		elementExists = true
		c.items[key].Value.(*cacheElement).value = value
		c.queue.MoveToFront(c.items[key])
	} else {
		c.items[key] = c.queue.PushFront(&cacheElement{key: key, value: value})
		if c.queue.Len() > c.capacity {
			oldest := c.queue.Back()
			c.queue.Remove(c.queue.Back())
			delete(c.items, oldest.Value.(*cacheElement).key)
		}
	}

	return elementExists
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, exists := c.items[key]; exists {
		c.queue.MoveToFront(item)
		return item.Value.(*cacheElement).value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
