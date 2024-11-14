package cache

import "errors"

type Cache struct {
	store map[string]interface{}
}

func New() *Cache {
	return &Cache{
		make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) error {
	if key != "" {
		c.store[key] = value
		return nil
	}
	return errors.New("key can not be empty")
}

func (c *Cache) Get(key string) (interface{}, bool) {
	value, exists := c.store[key]
	return value, exists
}

func (c *Cache) Delete(key string) error {
	_, exists := c.store[key]

	if exists {
		delete(c.store, key)
		return nil
	} else {
		return errors.New("invalid key")
	}
}
