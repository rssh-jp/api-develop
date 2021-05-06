package delivery

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rssh-jp/api-develop/domain"

	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo/v4"
)

type testUserUsecase struct {
	users []domain.User
}

func (uu *testUserUsecase) Fetch(ctx context.Context) ([]domain.User, error) {
	return uu.users, nil
}
func (uu *testUserUsecase) GetByID(ctx context.Context, id int64) (domain.User, error) {
	for _, user := range uu.users {
		if user.ID != id {
			continue
		}

		return user, nil
	}

	return domain.User{}, domain.ErrNotFound
}
func (uu *testUserUsecase) Update(ctx context.Context, user domain.User) error {
	for index, wu := range uu.users {
		if wu.ID != user.ID {
			continue
		}

		uu.users[index].Name = user.Name
		uu.users[index].Age = user.Age

		return nil
	}

	return domain.ErrNotFound
}

func TestSuccess(t *testing.T) {
	e := echo.New()

	uu := &testUserUsecase{[]domain.User{
		domain.User{
			ID:   1,
			Name: "test-name1",
			Age:  32,
		},
		domain.User{
			ID:   2,
			Name: "test-name2",
			Age:  18,
		},
	}}

	HandleUserHTTPDelivery(e, uu)

	ts := httptest.NewServer(e)

	defer ts.Close()

	for _, c := range []struct {
		name   string
		url    string
		method string
		input  interface{}
		expect interface{}
	}{
		{
			name:   "全ユーザ取得",
			url:    ts.URL + "/user",
			method: http.MethodGet,
			expect: []domain.User{
				domain.User{
					ID:   1,
					Name: "test-name1",
					Age:  32,
				},
				domain.User{
					ID:   2,
					Name: "test-name2",
					Age:  18,
				},
			},
		},
		{
			name:   "ID1のユーザ取得",
			url:    ts.URL + "/user/1",
			method: http.MethodGet,
			expect: domain.User{
				ID:   1,
				Name: "test-name1",
				Age:  32,
			},
		},
		{
			name:   "ID1のユーザ変更",
			url:    ts.URL + "/user/update",
			method: http.MethodPost,
			input:  `{"id": 1, "name": "change-test-1", "age": 33}`,
			expect: "OK",
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			var res *http.Response
			var err error
			switch c.method {
			case http.MethodGet:
				res, err = http.Get(c.url)
			case http.MethodPost:
				res, err = http.Post(c.url, "application/json", strings.NewReader(c.input.(string)))
			default:
				t.Fatal("Not specified method")
			}
			if err != nil {
				t.Fatal(err)
			}

			defer res.Body.Close()

			buf, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			expect, err := json.Marshal(c.expect)
			if err != nil {
				t.Fatal(err, string(buf))
			}

			actual := strings.Trim(string(buf), "\n")

			if !cmp.Equal(string(expect), actual) {
				t.Errorf("Could not match testcases.\nexpect: %+v\nactual: %+v\n", string(expect), actual)
			}
		})
	}

	t.Run("Confirm", func(t *testing.T) {
		res, err := http.Get(ts.URL + "/user/1")

		defer res.Body.Close()

		buf, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		expect, err := json.Marshal(domain.User{
			ID:   1,
			Name: "change-test-1",
			Age:  33,
		})
		if err != nil {
			t.Fatal(err)
		}

		actual := strings.Trim(string(buf), "\n")

		if !cmp.Equal(string(expect), actual) {
			t.Errorf("Could not match testcases.\nexpect: %+v\nactual: %+v\n", string(expect), actual)
		}
	})
}
