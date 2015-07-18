package main

/**
 *http://sdutlinux.org/t/1058
 *https://github.com/golang/groupcache/blob/master/lru/lru_test.go
 */

import (
	"container/list"
)

type Key interface{}

type Cache struct {
	MaxEntries int

	ll *list.List

	cache map[interface{}]*list.Element
}

type entry struct {
	key   Key
	value interface{}
}

func New(maxEntries int) *Cache {
	return &Cache{
		MaxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element)}
}

func (c *Cache) Add(key Key, value interface{}) (interface{}, bool) {

	if oldValue, ok := c.cache[key]; ok {
		c.ll.MoveToFront(oldValue)
		oldValue.Value.(*entry).value = value
		return oldValue.Value.(*entry).value, true
	}

	ele := c.ll.PushFront(&entry{key, value})
	c.cache[key] = ele

	for c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		c.RemoveOldest()
	}

	return nil, true
}

func (c *Cache) RemoveOldest() {
	if c.cache == nil {
		return
	}
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

func (c *Cache) Get(key Key) (interface{}, bool) {
	if c.cache == nil {
		return nil, true
	}

	if ele, hit := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}

	return nil, false
}

func (c *Cache) Remove(key Key) (interface{}, bool) {
	if c.cache == nil {
		return nil, true
	}

	if ele, hit := c.cache[key]; hit {
		c.RemoveOldest(ele)
	}
}

func (c *Cache) Len() int {
	if c.cache == nil {
		return 0
	}

	return c.ll.Len()
}
