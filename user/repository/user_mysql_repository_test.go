package repository

import (
	"context"
	"testing"

	"github.com/rssh-jp/api-develop/domain"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	query := "SELECT id, name, age FROM users"
	rows := sqlmock.NewRows([]string{"id", "name", "age"}).
		AddRow(1, "test1", 32).
		AddRow(2, "test2", 26)

	mock.ExpectQuery(query).WillReturnRows(rows)

	repo := NewUserMysqlRepository(db)

	users, err := repo.Fetch(context.TODO())
	if err != nil {
		t.Error(err)
	}

	t.Log(users)
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	query := "SELECT id, name, age FROM users WHERE id = ?"
	rows := sqlmock.NewRows([]string{"id", "name", "age"}).
		AddRow(1, "test1", 32)
	userID := int64(1)

	mock.ExpectQuery(query).
		WithArgs(userID).
		WillReturnRows(rows)

	repo := NewUserMysqlRepository(db)

	user, err := repo.GetByID(context.TODO(), userID)
	if err != nil {
		t.Error(err)
	}

	t.Log(user)
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	user := domain.User{
		ID:   1,
		Name: "change-test1",
		Age:  23,
	}

	query := "UPDATE users"

	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewUserMysqlRepository(db)

	err = repo.Update(context.TODO(), user)
	if err != nil {
		t.Error(err)
	}
}
