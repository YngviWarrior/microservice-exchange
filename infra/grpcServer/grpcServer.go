package grpcserver

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
)

type grpcServer struct {
	pb.UnimplementedExchangeServiceServer
	Db database.DatabaseInterface
}

func NewGrpcServer(db database.DatabaseInterface) *grpcServer {
	return &grpcServer{
		Db: db,
	}
}
