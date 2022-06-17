package main

import (
	"go_redis/handlers"
	"go_redis/repositories"
	"go_redis/services"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initialDB()
	redisClient := initRedis()
	_ = redisClient

	productRepo := repositories.NewProductRepositoryDB(db)
	productService := services.NewCatalogService(productRepo)
	productHandler := handlers.NewCatalogHandlerRedis(productService, redisClient)

	app := fiber.New()

	app.Get("products", productHandler.GetProducts)
	app.Listen(":8080")

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
