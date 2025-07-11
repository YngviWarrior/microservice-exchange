package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/YngviWarrior/microservice-exchange/infra/database"
	grpcserver "github.com/YngviWarrior/microservice-exchange/infra/grpcServer"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
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
	if err != nil {
		log.Fatal("failed to open log file:", err)
	}

	switch os.Getenv("ENVIROMENT") {
	case "local":
		log.SetOutput(os.Stdout)
		log.SetOutput(file)
	case "testnet", "demo":
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

	lis, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Running at port ", os.Getenv("PORT"))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
