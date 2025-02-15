package main

import (
	"github.com/sirupsen/logrus"
	"log"

	"github.com/a-melchikov/avito-shop/internal/config"
	"github.com/a-melchikov/avito-shop/internal/db"
	"github.com/a-melchikov/avito-shop/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()

	pool := db.NewPostgresDB(cfg)
	defer pool.Close()

	logrus.Println("Подключение к БД установлено.")

	app := fiber.New()

	app.Get("/", handler.HelloHandler)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
