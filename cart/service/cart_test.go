package cart_test

import (
	"context"
	"errors"
	"testing"

	cart "github.com/burubur/fastcampus/cart/service"
	"github.com/burubur/fastcampus/cart/service/mock"
	"github.com/burubur/fastcampus/cart/service/stub"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// Simulating interface on dependency with a Mock
func TestShoppingCart_AddItemToCart_ErrorOnRedis(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockRepositoryManager(ctrl)
	repositoryMock.EXPECT().AddToCart(context.Background(), "user-1", "product-a").Return(errors.New("failing on cache system"))

	scs := cart.New(repositoryMock)
	err := scs.AddItemToCart(context.Background(), "user-1", "product-a")

	assert.Error(t, err, "it should return an error")
}

// Simulating interface on dependency with a Mock
func TestShoppingCart_AddItemToCart_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockRepositoryManager(ctrl)
	repositoryMock.EXPECT().AddToCart(context.Background(), "user-2", "product-b").Return(nil)

	scs := cart.New(repositoryMock)
	err := scs.AddItemToCart(context.Background(), "user-2", "product-b")
	assert.NoError(t, err, "it should not return any error")
}

// Simulating interface on dependency with a Stub
func TestShoppingCart_AddItemToCart_Success_WithStub(t *testing.T) {
	repositoryStub := stub.NewRepositoryStub()
	scs := cart.New(repositoryStub)
	err := scs.AddItemToCart(context.Background(), "user-3", "product-c")
	assert.NoError(t, err, "it should not return any error")
}
