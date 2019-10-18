package main

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, sync.RWMutex{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
	lock  sync.RWMutex
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	i.lock.Lock()
	defer i.lock.Unlock()
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
