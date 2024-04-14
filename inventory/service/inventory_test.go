package service_test

import (
	"context"
	"testing"

	"github.com/burubur/fastcampus/inventory/service"
	"github.com/burubur/fastcampus/inventory/service/fake"
	"github.com/stretchr/testify/assert"
)

func TestInventory_AddStock(t *testing.T) {
	ctx := context.Background()
	fakeStore := fake.NewFakeStore()

	service := service.NewInventoryService(fakeStore)
	itemID := "item-1"
	initialStock := 10
	err := service.AddStock(ctx, itemID, initialStock)
	assert.NoError(t, err, "it should not return any error on the initial stage")

	currentStock, err := service.GetStock(ctx, itemID)
	assert.NoError(t, err, "it should not return error")

	assert.Equal(t, 10, currentStock, "it should return 5, after stock added on the previous step")
}
