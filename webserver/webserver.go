package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func NewInMemoryServer() *PlayerServer {
	return &PlayerServer{NewInMemoryPlayerStore()}
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var re = regexp.MustCompile(`^/players/(.*)`)
	matches := re.FindStringSubmatch(r.URL.Path)

	if len(matches) != 2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	player := matches[1]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
