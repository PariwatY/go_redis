package main

import (
	"fmt"
	"go_redis/repositories"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initialDB()
	redisClient := initRedis()

	productRepo := repositories.NewProductRepositoryRedis(db, redisClient)

	products, err := productRepo.GetProducts()
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(products)
}

func initialDB() *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/go_redis?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
