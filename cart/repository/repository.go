package repository

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	cacher redis.Cmdable
}

func New(redis redis.Cmdable) Cache {
	return Cache{
		cacher: redis,
	}
}

type CartData struct {
	ProductID string
	Count     int
}

func (c Cache) AddToCart(ctx context.Context, userID string, productID string) error {
	cacheKey := fmt.Sprintf("cart-%s", userID)
	cacheValues := map[string]interface{}{
		"id":          productID,
		"name":        "Sepatu lokal dari UMKM",
		"description": "loremipsum",
		"color":       "black",
		"size":        "41",
	}
	err := c.cacher.HSet(ctx, cacheKey, cacheValues).Err()
	if err != nil {
		return err
	}

	return nil
}
