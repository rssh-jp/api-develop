package repository

import (
	"context"
	"database/sql"
	"log"

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

func (ur *userMysqlRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]domain.User, error) {
	rows, err := ur.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println("Error rows.Close():", err)
		}
	}()

	ret := make([]domain.User, 0, 8)

	for rows.Next() {
		var id sql.NullInt64
		var name sql.NullString
		var age sql.NullInt32

		err := rows.Scan(&id, &name, &age)
		if err != nil {
			return nil, err
		}

		ret = append(ret, domain.User{
			ID:   id.Int64,
			Name: name.String,
			Age:  int(age.Int32),
		})
	}

	return ret, nil
}

func (ur *userMysqlRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	const query = "SELECT id, name, age FROM users"

	return ur.fetch(ctx, query)
}
func (ur *userMysqlRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
	const query = "SELECT id, name, age FROM users WHERE id = ?"

	args := []interface{}{
		id,
	}

	res, err := ur.fetch(ctx, query, args...)
	if err != nil {
		return domain.User{}, err
	}

	return res[0], nil
}
func (ur *userMysqlRepository) Update(ctx context.Context, user domain.User) error {
	const query = "UPDATE users SET name = ?, age = ? WHERE id = ?"
	args := []interface{}{
		user.Name,
		user.Age,
		user.ID,
	}

	_, err := ur.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
