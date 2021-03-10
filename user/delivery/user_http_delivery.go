package delivery

import (
	"net/http"
	"strconv"

	"github.com/rssh-jp/api-develop/domain"

	"github.com/labstack/echo/v4"
)

type userDelivery struct {
	userUsecase domain.UserUsecase
}

func HandleUserHTTPDelivery(e *echo.Echo, userUsecase domain.UserUsecase) {
	handler := &userDelivery{
		userUsecase: userUsecase,
	}

	e.GET("/user", handler.GetUser)
	e.GET("/user/:id", handler.GetByID)
}

func (ud *userDelivery) GetUser(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := ud.userUsecase.Fetch(ctx)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
func (ud *userDelivery) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	res, err := ud.userUsecase.GetByID(ctx, int64(id))
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
