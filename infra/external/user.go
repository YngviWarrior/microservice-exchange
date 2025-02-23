package external

import (
	"context"
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
)

func (u *external) GetUserById(id uint64) (out *pb.UserResponse) {
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
