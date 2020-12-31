package main

import (
	"context"
	pb "ecommerce-app/ordermanagement/proto/ecommerce"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type server struct{}

var orderMap = map[string]pb.Order{}

const (
	port = ":50051"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &server{})

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//AddOrder Unary
func (s *server) AddOrder(ctx context.Context, in *pb.Order) (*wrapperspb.StringValue, error) {
	orderId, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		return &wrapperspb.StringValue{Value: ""}, err
	}
	in.Id = orderId.String()
	orderMap[orderId.String()] = *in
	return &wrapperspb.StringValue{Value: orderId.String()}, nil
}

//GetOrder Unary
func (s *server) GetOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order, error) {
	ord := orderMap[orderId.Value]
	return &ord, nil
}

//SearchOrders server streaming
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

func (s *server) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {
	ordersStr := "Updated Order IDs : "
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			// Finished reading the order stream.
			return stream.SendAndClose(
				&wrappers.StringValue{Value: "Orders processed " + ordersStr})
		}
		// Update order
		orderMap[order.Id] = *order
		log.Print("Order ID ", order.Id, ": Updated")
		ordersStr += order.Id + ", "
	}
}
