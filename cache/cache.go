package main

import (
	"errors"
	"fmt"
)

// ErrNotFound — ошибка, которая возвращается, если ключ не найден.
var ErrNotFound = errors.New("value not found")

// Cache — интерфейс кеша.
type Cache interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

// simpleCache — простая реализация интерфейса Cache с использованием map.
type simpleCache struct {
	storage map[string]string
}

// NewCache — создаёт новый экземпляр simpleCache.
func NewCache() Cache {
	return &simpleCache{
		storage: make(map[string]string),
	}
}

// Set добавляет значение в кеш по заданному ключу.
func (c *simpleCache) Set(key, value string) error {
	c.storage[key] = value
	return nil
}

// Get возвращает значение из кеша по ключу.
func (c *simpleCache) Get(key string) (string, error) {
	value, ok := c.storage[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

// Delete удаляет значение из кеша по ключу.
func (c *simpleCache) Delete(key string) error {
	_, ok := c.storage[key]
	if !ok {
		return ErrNotFound
	}
	delete(c.storage, key)
	return nil
}

func main() {
	// Создаем кеш
	cache := NewCache()

	// Добавляем данные
	cache.Set("name", "John Doe")
	cache.Set("age", "30")

	// Получаем данные
	name, err := cache.Get("name")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Name:", name)
	}

	// Удаляем данные
	err = cache.Delete("name")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Name deleted")
	}

	// Пытаемся получить удалённый ключ
	_, err = cache.Get("name")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
