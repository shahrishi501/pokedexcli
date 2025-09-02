package pokecache

import (
	"testing"
	"time"
)


func TestCacheStoresAndRetrieves(t *testing.T) {
    c := NewCache(2 * time.Second)
    key := "test"
    val := []byte("hello")

    c.Add(key, val)
    got, ok := c.Get(key)
    if !ok || string(got) != "hello" {
        t.Fatalf("expected 'hello', got %s", got)
    }
}


func TestCacheExpiration(t *testing.T) {
    c := NewCache(1 * time.Second)
    c.Add("expire", []byte("bye"))

    time.Sleep(2 * time.Second) // wait longer than interval
    _, ok := c.Get("expire")
    if ok {
        t.Fatal("expected cache entry to expire, but it was found")
    }
}
