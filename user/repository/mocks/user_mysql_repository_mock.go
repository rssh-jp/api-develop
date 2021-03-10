package mocks

import (
	"context"

	"github.com/rssh-jp/api-develop/domain"
)

type userMysqlRepository struct {
	data []domain.User
}

func NewUserMysqlRepository() domain.UserRepository {
	return &userMysqlRepository{
		data: []domain.User{
			domain.User{
				ID:   1,
				Name: "test",
				Age:  25,
			},
			domain.User{
				ID:   2,
				Name: "test2",
				Age:  12,
			},
		},
	}
}

func (ur *userMysqlRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	return ur.data, nil
}
func (ur *userMysqlRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
	for _, d := range ur.data {
		if d.ID != id {
			continue
		}

		return d, nil
	}

	return domain.User{}, domain.ErrNotFound
}
func (ur *userMysqlRepository) Update(ctx context.Context, user domain.User) error {
	for index, d := range ur.data {
		if d.ID != user.ID {
			continue
		}

		ur.data[index] = user
		return nil
	}

	return domain.ErrNotFound
}
