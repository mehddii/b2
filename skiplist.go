package main

import (
	"cmp"
	"fmt"
	"math/rand/v2"
	"sync"
)

const (
	MaxLevel = 1000
	P        = 0.5
)

type SkipList[K cmp.Ordered, V any] interface {
	Get(K) (V, error)
	Insert(K, V)
	Delete(K)
}

type SkiptListNode[K cmp.Ordered, V any] struct {
	Key     K
	Value   V
	Forward []*SkiptListNode[K, V]
}

type SkipMap[K cmp.Ordered, V any] struct {
	Header *SkiptListNode[K, V]
	Level  int
	Mu     sync.RWMutex
}

func NewSkipMap[K cmp.Ordered, V any]() *SkipMap[K, V] {
	return &SkipMap[K, V]{
		Header: &SkiptListNode[K, V]{
			Forward: make([]*SkiptListNode[K, V], 1),
		},
		Level: 1,
	}
}

func (s *SkipMap[K, V]) Get(key K) (V, error) {
	current := s.Header
	for i := s.Level - 1; i >= 0; i-- {
		for current.Forward[i] != nil && current.Forward[i].Key < key {
			current = current.Forward[i]
		}
	}

	current = current.Forward[0]
	if current.Key != key {
		return *new(V), fmt.Errorf("Key %v does not exist", key)

	}

	return current.Value, nil
}

func randomLevel() int {
	level := 1
	for rand.Float32() < P && level < MaxLevel {
		level++
	}

	return level
}

func (s *SkipMap[K, V]) Insert(key K, value V) {
	update := make([]*SkiptListNode[K, V], s.Level)

	current := s.Header
	for i := s.Level - 1; i >= 0; i-- {
		for current.Forward[i] != nil && current.Forward[i].Key < key {
			current = current.Forward[i]
		}

		update[i] = current
	}

	current = current.Forward[0]
	if current != nil && current.Key == key {
		current.Value = value
	} else {
		level := randomLevel()
		if level > s.Level {
			for i := s.Level + 1; i <= level; i++ {
				s.Header.Forward = append(s.Header.Forward, nil)
			}

			for i := s.Level + 1; i <= level; i++ {
				update = append(update, s.Header)
			}

			s.Level = level
		}

		node := &SkiptListNode[K, V]{
			Key:     key,
			Value:   value,
			Forward: make([]*SkiptListNode[K, V], level),
		}

		for i := range level {
			node.Forward[i] = update[i].Forward[i]
			update[i].Forward[i] = node
		}
	}
}

func (s *SkipMap[K, V]) Delete(key K) {
	update := make([]*SkiptListNode[K, V], s.Level)

	current := s.Header
	for i := s.Level - 1; i >= 0; i-- {
		for current.Forward[i] != nil && current.Forward[i].Key < key {
			current = current.Forward[i]
		}

		update[i] = current
	}

	current = current.Forward[0]
	if current != nil && current.Key == key {
		for i := range s.Level {
			if update[i].Forward[i] != current {
				break
			}
			update[i].Forward[i] = current.Forward[i]
		}

		for s.Level > 0 && s.Header.Forward[s.Level-1] == nil {
			s.Level--
			s.Header.Forward = s.Header.Forward[:s.Level]
		}
	}
}
