package app

import (
	"log"

	"github.com/Dmytro-yakymuk/travelwithme/internal/config"
	"github.com/Dmytro-yakymuk/travelwithme/internal/handler"
	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
	"github.com/Dmytro-yakymuk/travelwithme/internal/server"
	"github.com/Dmytro-yakymuk/travelwithme/internal/service"
	"github.com/Dmytro-yakymuk/travelwithme/pkg/database/mysql"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal(err)
	}

	db := mysql.ConnectDB(cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.Name)

	// Migrations
	db.AutoMigrate(&models.Category{}, &models.Region{}, &models.User{}, &models.Tour{}, &models.Image{}, &models.Comment{}, &models.Trip{}, &models.Order{})

	// Seeding
	// for _, seed := range seeds.All() {
	// 	if err := seed.Run(db); err != nil {
	// 		log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
	// 	}
	// }

	// Services, Repos & API Handlers
	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	// HTTP Server
	srv := server.NewServer(cfg, handlers.Init())
	if err := srv.Run(); err != nil {
		log.Printf("error occurred while running http server: %s\n", err.Error())
	}

	log.Print("Server started")
}
