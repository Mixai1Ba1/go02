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

// SimpleCache — простая реализация интерфейса Cache с использованием map.
type SimpleCache struct {
	data map[string]string
}

// NewSimpleCache — создаёт новый экземпляр SimpleCache.
func NewSimpleCache() *SimpleCache {
	return &SimpleCache{
		data: make(map[string]string),
	}
}

// Set добавляет значение в кеш по заданному ключу.
func (c *SimpleCache) Set(key, value string) error {
	c.data[key] = value
	return nil
}

// Get возвращает значение из кеша по ключу.
func (c *SimpleCache) Get(key string) (string, error) {
	value, ok := c.data[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

// Delete удаляет значение из кеша по ключу.
func (c *SimpleCache) Delete(key string) error {
	_, ok := c.data[key]
	if !ok {
		return ErrNotFound
	}
	delete(c.data, key)
	return nil
}

func main() {
	// Создаем кеш
	cache := NewSimpleCache()

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