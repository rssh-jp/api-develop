package main

import (
	"log"

	userDelivery "github.com/rssh-jp/api-develop/user/delivery"
	userRepository "github.com/rssh-jp/api-develop/user/repository"
	userUsecase "github.com/rssh-jp/api-develop/user/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	const address = ":80"

	e := echo.New()

	ur := userRepository.NewUserMysqlRepository(nil, userRepository.OptionIsMock())
	uu := userUsecase.NewUserUsecase(ur)
	userDelivery.HandleUserHTTPDelivery(e, uu)

	log.Fatal(e.Start(address))
}
