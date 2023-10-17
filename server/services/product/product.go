package product

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/quocbang/grpc-gateway/pkg/pb"
	"github.com/quocbang/grpc-gateway/server/repositories"
)

type productService struct {
	Repo repositories.Repositories
}

func NewProductService(repo repositories.Repositories) pb.ProductServiceServer {
	return productService{
		Repo: repo,
	}
}

func (p productService) CreateProduct(context.Context, *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	return nil, nil
}

func (p productService) SearchProduct(req *pb.SearchProductRequest, stream pb.ProductService_SearchProductServer) error {
	n := 0
	for {
		if n == 100 {
			return nil
		}
		err := stream.Send(&pb.SearchProductResponse{
			Product: []*pb.Product{
				{
					ID:       uuid.NewString(),
					Color:    "blue",
					Producer: "Apple",
					Series:   "Iphone 15",
					AdvanceInfo: &pb.AdvanceInfo{
						Rom: &pb.Rom{
							Unit: pb.SizeUnit_GIGABYTE,
							Size: 256,
						},
						Ram: &pb.Ram{
							Unit: pb.SizeUnit_GIGABYTE,
							Size: 8,
						},
						Cpu: 2.4,
					},
				},
			},
		})
		if err != nil {
			return status.Error(codes.Internal, fmt.Sprintf("failed to send product, error %v", err))
		}
		n++
	}
}
