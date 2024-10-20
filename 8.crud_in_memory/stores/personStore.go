package stores

import (
	"crud_http/models"
)

type PersonStore struct {
	people map[string]models.Person
}

func NewPersonStore() *PersonStore {
	return &PersonStore{
		people: make(map[string]models.Person),
	}
}

func (p PersonStore) GetByName(name string) (person models.Person, found bool, err error) {
	val, ok := p.people[name]
	return val, ok, nil
}

func (p *PersonStore) Upsert(person models.Person) (err error) {
	p.people[person.Name] = person
	return nil
}
