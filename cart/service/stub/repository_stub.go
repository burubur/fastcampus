package stub

import "context"

type RepositoryStub struct{}

func (r RepositoryStub) AddToCart(ctx context.Context, userID string, productID string) (err error) {
	return nil
}

func NewRepositoryStub() RepositoryStub {
	return RepositoryStub{}
}
