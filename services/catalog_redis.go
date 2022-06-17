package services

import (
	"context"
	"encoding/json"
	"fmt"
	"go_redis/repositories"
	"time"

	"github.com/go-redis/redis/v8"
)

type caltalogServiceRedis struct {
	productRepo repositories.ProductRepository
	redisClient *redis.Client
}

func NewCatalogServiceRedis(productRepo repositories.ProductRepository, redisClient *redis.Client) CatalogService {
	return caltalogServiceRedis{productRepo, redisClient}
}

func (s caltalogServiceRedis) GetProducts() (products []Product, err error) {

	key := "service::GetProducts"
	//Redis Get Products
	if productJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(productJson), &products) == nil {
			fmt.Println("redis service")
			return products, nil
		}
	}

	//Repository
	productDB, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	for _, p := range productDB {
		products = append(products, Product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	// Redis Set
	if data, err := json.Marshal(products); err == nil {
		s.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	fmt.Println("database service")

	return products, nil

}
