package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "productinfo/service/ecommerce/productinfo/service/ecommerce"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

type server struct {
	productMap map[string]*pb.Product
}

var orderMap = map[string]pb.Order{}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
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

func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}

func (s *server) SearchOrders(searchQuery *wrappers.StringValue,
	stream pb.OrderManagement_SearchOrdersServer) error {
	for key, order := range orderMap {
		log.Print(key, order)
		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(
				itemStr, searchQuery.Value) {
				// Send the matching orders in a stream
				err := stream.Send(&order)
				if err != nil {
					return fmt.Errorf(
						"error sending message to stream : %v",
						err)
				}
				log.Print("Matching Order Found : " + key)
				break
			}
		}
	}
	return nil
}

func (s *server) GetOrder(ctx context.Context,
	orderId *wrappers.StringValue) (*pb.Order, error) {
	// Service Implementation.
	ord := orderMap[orderId.Value]
	return &ord, nil
}
