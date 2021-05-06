package repository

import (
	"context"
	"log"

	"github.com/rssh-jp/api-develop/domain"
	"github.com/rssh-jp/api-develop/user/repository/mocks"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

const (
	tableNameSpanner = "user"
)

var (
	userColumnsSpanner = []string{"name", "age"}
)

type userSpannerRepository struct {
	client *spanner.Client
}

func NewUserSpannerRepository(client *spanner.Client, opts ...option) domain.UserRepository {
	c := new(config)

	for _, opt := range opts {
		opt(c)
	}

	if c.isMock {
		return mocks.NewUserMysqlRepository()
	}

	return &userSpannerRepository{
		client: client,
	}
}

func (ur *userSpannerRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	log.Println("+++++", 1)
	iter := ur.client.Single().Read(ctx, tableNameSpanner, spanner.AllKeys(), []string{"id", "name", "age"})

	defer iter.Stop()

	ret := make([]domain.User, 0, 8)
	log.Println("+++++", 2)

	for {
		log.Println("+++++", 3)
		row, err := iter.Next()
		log.Println("+++++", 4, err)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		log.Println("+++++", 5)
		var id, age int64
		var name string
		err = row.Columns(&id, &name, &age)
		log.Println("+++++", 6, err)
		if err != nil {
			return nil, err
		}

		user := domain.User{
			ID:   id,
			Name: name,
			Age:  int(age),
		}

		log.Println(user)

		ret = append(ret, user)
		log.Println("+++++", 7)
	}

	log.Println("+++++", 8)
	return ret, nil
}
func (ur *userSpannerRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
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
func (ur *userSpannerRepository) Update(ctx context.Context, user domain.User) error {

	mut, err := spanner.UpdateStruct(tableNameSpanner, user)
	if err != nil {
		return err
	}

	_, err = ur.client.Apply(ctx, []*spanner.Mutation{mut})
	if err != nil {
		return err
	}

	return nil
}
