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

func (c *lruCache) Set(key Key, value interface{}) bool {
	var elementExists bool

	if _, exists := c.items[key]; exists {
		elementExists = true
		c.items[key].Value = value
		c.queue.MoveToFront(c.items[key])
	} else {
		newElement := &ListItem{
			Value: value,
		}
		c.items[key] = newElement
		c.queue.PushFront(key)
		if c.queue.Len() > c.capacity {
			delete(c.items, c.queue.Back().Value.(Key))
			c.queue.Remove(c.queue.Back())
		}
	}

	return elementExists
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if _, exists := c.items[key]; exists {
		c.queue.MoveToFront(c.items[key])
		return c.items[key].Value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {

}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
