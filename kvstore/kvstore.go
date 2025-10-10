package kvstore

import (
	"errors"
	"sync"
)

type KVStore interface {
	Set(k string, v string)
	Get(k string) (string, bool)
	Del(k string) error
}

type Store struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewKVStore() *Store {
	return &Store{
		data: make(map[string]string),
		mu:   sync.RWMutex{},
	}
}

func (s *Store) Set(k string, v string) {
	s.mu.Lock() // we are writing so we use Lock (write lock)
	defer s.mu.Unlock()

	s.data[k] = v
}

func (s *Store) Get(k string) (string, bool) {
	s.mu.RLock() // we aren't writing so we use RLock (read lock)
	defer s.mu.RUnlock()

	v, ok := s.data[k]
	return v, ok
}

func (s *Store) Del(k string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[k]; !ok {
		return errors.New("key not found")
	}

	delete(s.data, k)
	return nil
}
