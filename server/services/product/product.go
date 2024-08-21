package product

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/quocbang/grpc-gateway/pkg/pb"
	"github.com/quocbang/grpc-gateway/server/repositories"
	"github.com/quocbang/grpc-gateway/server/repositories/orm/models"
)

type productService struct {
	Repo repositories.Repositories
}

func NewProductService(repo repositories.Repositories) pb.ProductServiceServer {
	return productService{
		Repo: repo,
	}
}

func (p productService) CreateProduct(context.Context, *pb.CreateProductRequest) (*pb.CommonCreateResponse, error) {
	return &pb.CommonCreateResponse{
		AffectedRows: 1,
	}, nil
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

func (p productService) ClientStream(stream pb.ProductService_ClientStreamServer) error {
	counter := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			err := stream.SendAndClose(&pb.CommonCreateResponse{
				AffectedRows: int32(counter),
			})
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}

		// increase counter
		counter++
		log.Println(req)
	}
}

func (p productService) CreateProductsStream(stream pb.ProductService_CreateProductsStreamServer) error {
	affectedRows := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			err := stream.SendAndClose(&pb.CommonCreateResponse{
				AffectedRows: int32(affectedRows),
			})
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}

		// insert to pg
		products := make([]models.Product, len(req.Products))
		for idx, product := range req.Products {
			products[idx] = models.Product{
				ID:          uuid.New(),
				ProductCode: product.ID,
				Color:       product.Color,
				Producer:    product.Producer,
				Series:      product.Series,
				AdvanceInfo: &models.Advance{
					Rom: &models.Capacity{
						Size: product.AdvanceInfo.Rom.Size,
						Unit: product.AdvanceInfo.Rom.Unit,
					},
					Ram: &models.Capacity{
						Size: product.AdvanceInfo.Ram.Size,
						Unit: product.AdvanceInfo.Ram.Unit,
					},
					CPU: product.AdvanceInfo.Cpu,
				},
			}
		}
		if err := p.Repo.Product().Creates(stream.Context(), products); err != nil {
			err := stream.SendAndClose(&pb.CommonCreateResponse{
				AffectedRows: int32(affectedRows),
			})
			if err != nil {
				return err
			}
		}
	}
}
