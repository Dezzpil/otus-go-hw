package hw04_lru_cache //nolint:golint,stylecheck
import "sync"

type Key string

type Cache interface {
	Set(key string, value interface{}) bool // Добавить значение в кэш по ключу
	Get(key string) (interface{}, bool)     // Получить значение из кэша по ключу
	Clear()                                 // Очистить кэш
	Keys() []string
}

type lruCache struct {
	sync.Mutex

	capacity int
	queue    List
	items    map[Key]*cacheItem
}

type cacheItem struct {
	value     interface{}
	queueItem *ListItem
}

func (l *lruCache) Set(key string, value interface{}) bool {
	k := Key(key)
	l.Lock()
	defer l.Unlock()
	if ci, ok := l.items[k]; ok {
		ci.value = value
		l.queue.MoveToFront(ci.queueItem)
		return ok
	}

	if l.queue.Len() >= l.capacity {
		li := l.queue.Back()
		delete(l.items, li.Value.(Key))
		l.queue.Remove(li)
	}

	li := l.queue.PushFront(k)
	ci := cacheItem{value, li}
	l.items[k] = &ci
	return false
}

func (l *lruCache) Get(key string) (interface{}, bool) {
	l.Lock()
	defer l.Unlock()
	if ci, ok := l.items[Key(key)]; ok {
		l.queue.MoveToFront(ci.queueItem)
		return ci.value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.Lock()
	l.items = map[Key]*cacheItem{}
	l.Unlock()
}

func (l *lruCache) Keys() []string {
	keys := make([]string, 0, l.queue.Len())
	for k := range l.items {
		keys = append(keys, string(k))
	}
	return keys
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    map[Key]*cacheItem{},
	}
}
