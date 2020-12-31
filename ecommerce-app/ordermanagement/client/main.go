package main

import (
	"context"
	pb "ecommerce-app/ordermanagement/proto/ecommerce"
	"fmt"
	"io"
	"log"

	"time"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewOrderManagementClient(conn)
	fmt.Println(AddOrder(client))
	fmt.Println(GetOrder(client))
	fmt.Println(SearchOrders(client))
	UpdateOrders(client)
	fmt.Println(GetOrder(client))
}

func GetOrder(client pb.OrderManagementClient) (*pb.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	orderID := &wrapperspb.StringValue{Value: "1"}
	order, err := client.GetOrder(ctx, orderID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return order, nil
}

func AddOrder(client pb.OrderManagementClient) (*wrapperspb.StringValue, error) {
	// orders := []pb.Order{
	// 	pb.Order{Items: []string{"watch", "mobile"}, Description: "Electronics", Price: 100},
	// }
	order := &pb.Order{Items: []string{"watch", "mobile"}, Description: "Electronics", Price: 100}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return client.AddOrder(ctx, order)
}

func SearchOrders(client pb.OrderManagementClient) ([]*pb.Order, error) {
	orders := []*pb.Order{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := client.SearchOrders(ctx, &wrapperspb.StringValue{Value: "watch"})
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Print(order)
		orders = append(orders, order)
	}
	return orders, nil
}

func UpdateOrders(client pb.OrderManagementClient) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	updateStream, err := client.UpdateOrders(ctx)
	if err != nil {
		log.Fatalf("%v.UpdateOrders(_) = _, %v", client, err)
	}
	updOrder1 := pb.Order{
		Id:          "1",
		Items:       []string{"trimmer", "comb"},
		Price:       1200,
		Destination: "Mall",
	}
	// Updating order 1
	if err := updateStream.Send(&updOrder1); err != nil {
		log.Fatalf("%v.Send(%v) = %v",
			updateStream, updOrder1, err)
	}
	updOrder2 := pb.Order{
		Id:          "2",
		Items:       []string{"washing machine", "refrigerator"},
		Price:       15000,
		Destination: "Office",
	}
	// Updating order 2
	if err := updateStream.Send(&updOrder2); err != nil {
		log.Fatalf("%v.Send(%v) = %v",
			updateStream, updOrder2, err)
	}
	updOrder3 := pb.Order{
		Id:          "3",
		Items:       []string{"Copier", "Printer"},
		Price:       1080,
		Destination: "Head Office",
	}
	// Updating order 3
	if err := updateStream.Send(&updOrder3); err != nil {
		log.Fatalf("%v.Send(%v) = %v",
			updateStream, updOrder3, err)
	}
	updateRes, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v",
			updateStream, err, nil)
	}
	log.Printf("Update Orders Res : %s", updateRes)
}
