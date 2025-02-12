package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/YngviWarrior/microservice-exchange/infra/database"
	grpcserver "github.com/YngviWarrior/microservice-exchange/infra/grpcServer"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(".env file is missing")
	}

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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

	database := database.NewDatabase()

	exchangeService := grpcserver.NewGrpcServer(database)

	grpcServer := grpc.NewServer()

	pb.RegisterExchangeServiceServer(grpcServer, exchangeService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", os.Getenv("port"))
	if err != nil {
		panic(err)
	}

	log.Println("Running at port ", os.Getenv("port"))
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

	// exchangeRepo := mysql.NewExchangeRepository(database)
	// tradeConfigRepo := mysql.NewTradeConfigRepository(database)

	// // should return all usecases
	// allUseCases := usecase.NewUsecase(
	// 	exchangeRepo,
	// 	tradeConfigRepo,
	// )
	// // should return all controllers
	// controllers := controller.NewController(allUseCases)

	// if err != nil {
	// 	log.Printf("%v", err)
	// }

	// server.NewServer().InitServer(controllers)
}
