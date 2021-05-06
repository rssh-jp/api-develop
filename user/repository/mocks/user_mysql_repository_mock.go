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
			domain.User{
				ID:   3,
				Name: "test3",
				Age:  37,
			},
			domain.User{
				ID:   4,
				Name: "test4",
				Age:  14,
			},
			domain.User{
				ID:   5,
				Name: "test5",
				Age:  53,
			},
			domain.User{
				ID:   6,
				Name: "test6",
				Age:  49,
			},
			domain.User{
				ID:   7,
				Name: "test7",
				Age:  24,
			},
			domain.User{
				ID:   8,
				Name: "test8",
				Age:  22,
			},
			domain.User{
				ID:   9,
				Name: "test9",
				Age:  35,
			},
			domain.User{
				ID:   10,
				Name: "test10",
				Age:  72,
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
