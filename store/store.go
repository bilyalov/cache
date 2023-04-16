package store

import (
	"sync"
	"time"
)

type Store struct {
	data map[string]any
	mu   *sync.RWMutex
}

func New() Store {
	return Store{
		data: make(map[string]any),
		mu:   new(sync.RWMutex),
	}
}

func (s *Store) expire(name string, ttl time.Duration) {
	if ttl != 0 {
		go func() {
			time.Sleep(ttl)

			s.Delete(name)
		}()
	}
}

func (s *Store) Set(name string, value any, ttl time.Duration) {
	defer s.expire(name, ttl)

	s.mu.Lock()
	s.data[name] = value
	s.mu.Unlock()
}

func (s *Store) Get(name string) (val any, ok bool) {
	s.mu.RLock()
	val, ok = s.data[name]
	s.mu.RUnlock()

	return
}

func (s *Store) Delete(name string) {
	s.mu.Lock()
	delete(s.data, name)
	s.mu.Unlock()
}
