package delivery

import (
	"net/http"

	"github.com/rssh-jp/api-develop/domain"
	"github.com/rssh-jp/api-develop/internal/http/echo/gen"

	"github.com/labstack/echo/v4"
)

type userDelivery struct {
	userUsecase domain.UserUsecase
}

func HandleUserHTTPDelivery(e *echo.Echo, userUsecase domain.UserUsecase) {
	handler := &userDelivery{
		userUsecase: userUsecase,
	}

	gen.RegisterHandlers(e, handler)
}

func (ud *userDelivery) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := ud.userUsecase.Fetch(ctx)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (ud *userDelivery) GetByID(c echo.Context, id int64) error {
	ctx := c.Request().Context()

	res, err := ud.userUsecase.GetByID(ctx, id)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (ud *userDelivery) Update(c echo.Context) error {
	ctx := c.Request().Context()

	r := new(gen.User)

	err := c.Bind(r)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	user := domain.User{
		ID:   r.Id,
		Name: r.Name,
		Age:  int(r.Age),
	}

	err = ud.userUsecase.Update(ctx, user)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "OK")
}
