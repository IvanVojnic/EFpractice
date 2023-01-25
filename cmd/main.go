package main

import "EFpractic2/pkg/repository"

type Config struct {
	Storage repository.StorageConfig `yaml:"storage"`
}

func main() {
	/*
		repos := repository.NewRepository(db)
		services := service.NewService(repos)
		handlers := handler.NewHandler(services)

		srv := new(EFpractic2.Server)
		if err := srv.Run("3000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error", err.Error())
		}*/
}
