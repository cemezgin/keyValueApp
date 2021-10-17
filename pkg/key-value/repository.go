package key_value

import (
	"encoding/json"
	"sync"
)

type Repository struct {
	Store *DataStore
}

type DataStore struct {
	M map[string]Item
	*sync.RWMutex
}

func NewRepository(store *DataStore) *Repository {
	return &Repository{ store}
}

func (rp *Repository) List() ([]byte, error){
	rp.Store.RLock()
	items := make([]Item, 0, len(rp.Store.M))

	for _, v := range rp.Store.M {
		items = append(items, v)
	}
	rp.Store.RUnlock()
	jsonBytes, err := json.Marshal(items)

	return jsonBytes, err
}

func (rp *Repository) Get(value string) (Item, bool) {

	rp.Store.RLock()
	u, ok := rp.Store.M[value]
	rp.Store.RUnlock()

	return u, ok
}

func (rp *Repository) Create(i Item) ([]byte, error){

	rp.Store.Lock()
	rp.Store.M[i.Key] = i
	rp.Store.Unlock()
	jsonBytes, err := json.Marshal(i)

	return jsonBytes, err
}
