package repository

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rssh-jp/api-develop/domain"
	"github.com/rssh-jp/api-develop/user/repository/mocks"

	"cloud.google.com/go/bigtable"
)

const (
	tableNameBigtable        = "user"
	columnNameBigtable       = "user-column"
	columnFamilyNameBigtable = "user-family"
)

type userBigtableRepository struct {
	client *bigtable.Client
}

func NewUserBigtableRepository(client *bigtable.Client, opts ...option) domain.UserRepository {
	c := new(config)

	for _, opt := range opts {
		opt(c)
	}

	if c.isMock {
		return mocks.NewUserMysqlRepository()
	}

	return &userBigtableRepository{
		client: client,
	}
}

func (ur *userBigtableRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	ret := make([]domain.User, 0, 8)

	tbl := ur.client.Open(tableNameBigtable)

	err := tbl.ReadRows(ctx, bigtable.PrefixRange(columnNameBigtable), func(row bigtable.Row) bool {
		item := row[columnFamilyNameBigtable][0]
		log.Printf("\t%s = %s\n", item.Row, string(item.Value))
		var user domain.User
		err := json.Unmarshal(item.Value, &user)
		if err != nil {
			log.Println("ERR json.Unmarshal:", err)
		}

		ret = append(ret, user)

		return true
	})
	if err != nil {
		return nil, err
	}

	return ret, nil
}
func (ur *userBigtableRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
	users, err := ur.Fetch(ctx)
	if err != nil {
		return domain.User{}, err
	}

	for _, user := range users {
		if user.ID != id {
			continue
		}

		return user, nil
	}

	return domain.User{}, domain.ErrNotFound
}
func (ur *userBigtableRepository) Update(ctx context.Context, user domain.User) error {
	return nil
}
