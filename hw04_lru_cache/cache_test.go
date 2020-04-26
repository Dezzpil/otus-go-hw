package hw04_lru_cache //nolint:golint,stylecheck

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("least recently", func(t *testing.T) {
		cache := NewCache(3)
		data := []struct {
			key string
			val int
		}{
			{"foo", 100},
			{"bar", 200},
			{"pew", 300},
			{"baz", 400},
		}
		keys := make([]string, 0, 3)
		for _, s := range data {
			keys = append(keys, s.key)
			cache.Set(s.key, s.val)
		}
		require.ElementsMatch(t, keys[1:], cache.Keys())
	})

	t.Run("purge logic", func(t *testing.T) {
		cache := NewCache(10)
		cache.Set("jhon", true)
		cache.Set("bob", false)
		var was bool

		_, was = cache.Get("jhon")
		require.True(t, was)
		cache.Clear()

		_, was = cache.Get("bob")
		require.False(t, was)
		require.Equal(t, 0, len(cache.Keys()))
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove if task with asterisk completed

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			//c.Set(Key(strconv.Itoa(i)), i)
			c.Set(strconv.Itoa(i), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			//c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
			c.Get(strconv.Itoa(rand.Intn(1_000_000)))
		}
	}()

	wg.Wait()
}
