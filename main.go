package microserviceexchange

import (
	"log"
	"os"
	"time"

	"github.com/YngviWarrior/microservice-exchange/controller"
	"github.com/YngviWarrior/microservice-exchange/infra/server"
	"github.com/YngviWarrior/microservice-exchange/usecases"
	"github.com/joho/godotenv"
)

func InitServer() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(".env file is missing")
	}

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	// should return all usecases
	allUseCases := usecases.InitUseCases()
	// should return all controllers
	controllers := controller.InitController(&allUseCases)

	switch os.Getenv("ENVIROMENT") {
	case "local":
		log.SetOutput(os.Stdout)
		log.SetOutput(file)

	case "testnet":
		log.SetOutput(os.Stdout)

	case "server":
		loc, _ := time.LoadLocation("America/Sao_Paulo")
		time.Local = loc

		log.SetOutput(file)
	}

	if err != nil {
		log.Printf("%v", err)
	}

	server.NewServer().InitServer(controllers)
}
