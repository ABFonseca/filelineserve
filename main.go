package main

// @title           File Line Serve API
// @version         1.0
// @description     API receive requests for a text file line and returns it.

// @BasePath  /

import (
	"log"
	"os"

	"filelineserve/api"
	"filelineserve/entity"
	"filelineserve/service"

	"github.com/Netflix/go-env"
	"github.com/subosito/gotenv"
)

const defaultAPIPort = 8080

type Config struct {
	Port int `env:"API_PORT"`
}

func main() {
	cfg := &Config{}
	handleEnv(cfg)

	inputArgs := os.Args[1:]
	if len(inputArgs) < 1 {
		log.Fatal("invalid number of arguments, please input path/name of file to be served")
	}
	filePath := inputArgs[0]

	fileRepo := entity.NewFileRepository(filePath)
	fileService := service.NewFileService(fileRepo)

	FileLineServer := api.NewFileServeAPI(cfg.Port, fileService)

	log.Printf("Starting File Line Server on port:%d with file:%q", cfg.Port, filePath)

	FileLineServer.Run()

}

// handleEnv loads config from .env file and environment variable
// If Port is not found, the default port 8080 will be used
func handleEnv(cfg *Config) {
	if _, err := os.Stat(".env"); err == nil {
		err := gotenv.Load(".env")
		if err != nil {
			log.Fatal("error loading environment variables file to memory")
		}
	}

	if _, err := env.UnmarshalFromEnviron(cfg); err != nil {
		log.Fatal("error unmarshaling environment variables file")
	}
	if cfg.Port == 0 {
		cfg.Port = defaultAPIPort
	}
}
