package external

import (
	"log"
	"os"

	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"google.golang.org/grpc"
)

type external struct {
	Conn *grpc.ClientConn
}

type ExternalInterface interface {
	GetUserById(id uint64) (out *pb.UserResponse)
}

func NewUserExternal() ExternalInterface {
	conn, err := grpc.Dial(os.Getenv("MS_USER"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("NUE: %v", err)
	}

	return &external{
		Conn: conn,
	}
}
