package main

import (
	pb "app-04/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// type server struct{}

/* var productList = []product.Product{
	{
		Id:          "1",
		Name:        "Book",
		Description: "Let Us C",
	},
	{
		Id:          "2",
		Name:        "Mobile",
		Description: "Samsung",
	},
} */

type server struct {
	productMap map[string]*pb.Product
}

func (s *server) AddProduct(ctx context.Context,
	in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"Error while generating Product ID", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

/* func (s server) AddProduct(context.Context, *product.Product) (*product.ProductID, error) {
	return nil, nil
}
func (s server) GetProduct(context.Context, *product.ProductID) (*product.Product, error) {
	return nil, nil
}
func (s server) mustEmbedUnimplementedProductInfoServer() {}
*/

/* func (s *server) AddProduct(ctx context.Context, in *product.Product) (*product.ProductID, error) {
	in.Id = strconv.Itoa(len(productList) + 1)
	productList = append(productList, *in)
	productID := &product.ProductID{
		Id: in.Id,
	}
	return productID, nil
}

func (s *server) GetProduct(ctx context.Context, in *product.ProductID) (*product.Product, error) {
	for _, p := range productList {
		if p.GetId() == in.Id {
			return &p, nil
		}
	}
	return nil, nil
}

func (s *server) mustEmbedUnimplementedProductInfoServer() {

} */
