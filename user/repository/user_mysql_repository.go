package repository

import (
	"context"
	"database/sql"

	"github.com/rssh-jp/api-develop/domain"
	"github.com/rssh-jp/api-develop/user/repository/mocks"
)

type userMysqlRepository struct {
	db *sql.DB
}

func NewUserMysqlRepository(db *sql.DB, opts ...option) domain.UserRepository {
	c := new(config)

	for _, opt := range opts {
		opt(c)
	}

	if c.isMock {
		return mocks.NewUserMysqlRepository()
	}

	return &userMysqlRepository{
		db: db,
	}
}

func (ur *userMysqlRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	return []domain.User{
		domain.User{
			ID:   1,
			Name: "test",
			Age:  25,
		},
	}, nil
	return nil, nil
}
func (ur *userMysqlRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
	return domain.User{}, nil
}
func (ur *userMysqlRepository) Update(ctx context.Context, user domain.User) error {
	return nil
}
