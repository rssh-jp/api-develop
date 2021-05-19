package main

import (
	"context"
	"log"
	"time"

	"github.com/rssh-jp/api-develop/internal/grpc/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewUsersClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	user := &pb.User{
		Id:   1,
		Name: "change-name1",
		Age:  33,
	}

	updateReply, err := c.Update(ctx, &pb.UpdateRequest{User: user})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(updateReply)

	users, err := c.Fetch(ctx, &pb.FetchRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users.Users {
		log.Println(user)
	}
}
