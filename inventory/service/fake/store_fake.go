package fake

import (
	"context"
	"errors"
)

type FakeStore struct {
	inventory map[string]int
}

func NewFakeStore() *FakeStore {
	return &FakeStore{inventory: make(map[string]int)}
}

func (f *FakeStore) AddItem(ctx context.Context, itemID string, quantity int) (err error) {
	if quantity < 0 {
		return errors.New("quantity can't be negative")
	}

	f.inventory[itemID] += quantity
	return nil
}

func (f *FakeStore) GetStock(ctx context.Context, itemID string) (numStock int, err error) {
	numStock, exist := f.inventory[itemID]
	if !exist {
		return 0, errors.New("item not found")
	}

	return numStock, nil
}
