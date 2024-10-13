package main

type InMemoryPlayerStore struct {
}

func (inMemoryStore *InMemoryPlayerStore) GetPlayerScore(name string) int {
	if name == "Pepper" {
		return 20
	}
	return 10
}
