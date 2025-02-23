package grpcserver

import (
	"context"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-exchange/usecase"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (g *grpcServer) ListParity(ctx context.Context, in *pb.ListParityRequest) (out *pb.ListParityResponse, err error) {
	parityRepo := mysql.NewParityRepository(g.Db)
	usecase := usecase.NewUsecaseListParity(parityRepo)

	list, err := usecase.ListParity(&usecasedto.InputListParityDto{})
	if err != nil {
		return
	}

	out = &pb.ListParityResponse{
		Parities: []*pb.Parity{},
	}

	for _, v := range list {
		var o pb.Parity

		o.Active = v.Active
		o.Parity = v.Parity
		o.Symbol = v.Symbol

		out.Parities = append(out.Parities, &o)
	}

	return
}
