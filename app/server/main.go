package main

import (
	"context"
	"fmt"
	"log"
	"net"

	userDelivery "github.com/rssh-jp/api-develop/user/delivery"
	userRepository "github.com/rssh-jp/api-develop/user/repository"
	userUsecase "github.com/rssh-jp/api-develop/user/usecase"

	"cloud.google.com/go/bigtable"
	"cloud.google.com/go/spanner"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	const address = ":80"
	const grpcAddress = ":50051"

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	//ctx := context.Background()

	//spannerClient, err := makeSpannerClient(ctx, "test-project", "test-api-spanner", "test-db")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//bigtableClient, err := makeBigtableClient(ctx, "test-project", "test-api-bigtable")
	//if err != nil {
	//	log.Fatal(err)
	//}

	ur := userRepository.NewUserMysqlRepository(nil, userRepository.OptionIsMock())
	//ur := userRepository.NewUserSpannerRepository(spannerClient)
	//ur := userRepository.NewUserBigtableRepository(bigtableClient)
	uu := userUsecase.NewUserUsecase(ur)
	userDelivery.HandleUserHTTPDelivery(e, uu)
	grpcServer := userDelivery.HandleUserGRPCDelivery(uu)

	chErr := make(chan error)
	defer close(chErr)

	go func() {
		chErr <- e.Start(address)
	}()

	go func() {
		log.Println("gRPC Server START", grpcAddress)
		listener, err := net.Listen("tcp", grpcAddress)
		if err != nil {
			log.Fatal(err)
		}

		chErr <- grpcServer.Serve(listener)
	}()

	select {
	case err := <-chErr:
		log.Fatal(err)
	}
}

func makeSpannerClient(ctx context.Context, project, instance, db string) (*spanner.Client, error) {
	dsn := fmt.Sprintf("projects/%s/instances/%s/databases/%s", project, instance, db)
	log.Println("-----------------", 1)
	client, err := spanner.NewClient(ctx, dsn)
	log.Println("-----------------", 2, err)
	if err != nil {
		return nil, err
	}

	log.Println("-----------------", 4)

	return client, nil

}
func makeBigtableClient(ctx context.Context, project, instance string) (*bigtable.Client, error) {
	log.Println("-----------------", 1)
	client, err := bigtable.NewClient(ctx, project, instance)
	log.Println("-----------------", 2, err)
	if err != nil {
		return nil, err
	}

	return client, nil

}
