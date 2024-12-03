package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.entry == nil {
		t.Error("cache is nil")
	}

}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cache.Add("key1", []byte("val1"))
	actual, ok := cache.Get("key1")
	if !ok {
		t.Error("key1 not found")
	}

	if string(actual) != "val1" {
		t.Error("key, val not match")
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cache.Add("key1", []byte("val1"))
	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get("key1")
	if ok {
		t.Error("key1 shoud have been reaped")
	}
}
