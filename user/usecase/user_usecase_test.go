package usecase

import (
	"context"
	"testing"

	"github.com/rssh-jp/api-develop/domain"

	"github.com/google/go-cmp/cmp"
)

type userRepository struct {
	users []domain.User
}

func (ur *userRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	return ur.users, nil
}
func (ur *userRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
	for _, item := range ur.users {
		if item.ID != id {
			continue
		}

		return item, nil
	}

	return domain.User{}, domain.ErrNotFound
}
func (ur *userRepository) Update(ctx context.Context, user domain.User) error {
	for index, item := range ur.users {
		if item.ID != user.ID {
			continue
		}

		ur.users[index].Name = user.Name
		ur.users[index].Age = user.Age

		return nil
	}

	return domain.ErrNotFound
}

func TestUserUsecase(t *testing.T) {
	data := []domain.User{
		domain.User{
			ID:   1,
			Name: "test-name-1",
			Age:  32,
		},
		domain.User{
			ID:   2,
			Name: "test-name-2",
			Age:  28,
		},
	}

	uu := NewUserUsecase(&userRepository{data})

	t.Run("Fetch", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			expect := data

			actual, err := uu.Fetch(context.Background())
			if err != nil {
				t.Error(err)
			}

			if !cmp.Equal(expect, actual) {
				t.Errorf("Not match.\nexpect: %+v\nactual: %+v\n", expect, actual)
			}
		})
	})

	t.Run("GetByID", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			expect := data[0]

			actual, err := uu.GetByID(context.Background(), 1)
			if err != nil {
				t.Error(err)
			}

			if !cmp.Equal(expect, actual) {
				t.Errorf("Not match.\nexpect: %+v\nactual: %+v\n", expect, actual)
			}
		})

		t.Run("Failure", func(t *testing.T) {
			_, err := uu.GetByID(context.Background(), 3)

			if err != domain.ErrNotFound {
				t.Errorf("Not error.\nexpect: %v\nactual: nil", domain.ErrNotFound)
			}
		})
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			err := uu.Update(context.Background(), domain.User{
				ID:   1,
				Name: "change-test-name-1",
				Age:  38,
			})
			if err != nil {
				t.Error(err)
			}
		})

		t.Run("Failure", func(t *testing.T) {
			err := uu.Update(context.Background(), domain.User{
				ID:   3,
				Name: "change-test-name-1",
				Age:  38,
			})

			if err != domain.ErrNotFound {
				t.Errorf("Not error.\nexpect: %v\nactual: nil", domain.ErrNotFound)
			}
		})
	})
}
