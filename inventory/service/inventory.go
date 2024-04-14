package service

import "context"

// inventory service
// data store interface

type StoreManager interface {
	AddItem(ctx context.Context, itemID string, quantity int) (err error)
	GetStock(ctx context.Context, itemID string) (numStock int, err error)
}

type Inventory struct {
	Store StoreManager
}

func NewInventoryService(store StoreManager) Inventory {
	return Inventory{
		Store: store,
	}
}

func (i Inventory) AddStock(ctx context.Context, itemID string, quantity int) error {
	// another complex business logic reside here
	return i.Store.AddItem(ctx, itemID, quantity)
}

func (i Inventory) GetStock(ctx context.Context, itemID string) (int, error) {
	return i.Store.GetStock(ctx, itemID)
}
