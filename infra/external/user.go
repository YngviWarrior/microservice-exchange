package external

import (
	"context"
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"google.golang.org/grpc"
)

type userExternal struct {
	Conn *grpc.ClientConn
}

type UserExternalInterface interface {
	GetUserById(id uint64) (out *pb.UserResponse)
}

func NewUserExternal() UserExternalInterface {
	conn, err := grpc.Dial("ms-user:3002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("CMU: %v", err)
	}

	return &userExternal{
		Conn: conn,
	}
}

func (u *userExternal) GetUserById(id uint64) (out *pb.UserResponse) {
	out = &pb.UserResponse{}
	client := pb.NewUserServiceClient(u.Conn)

	out, err := client.GetUserById(context.Background(), &pb.GetUserByIdRequest{
		Id: id,
	})

	if err != nil {
		log.Fatalln("UI-GUBI 00: ", err)
	}

	return
}
