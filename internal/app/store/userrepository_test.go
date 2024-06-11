package store_test

import (
	"testing"

	"github.com/s-cod/rest-api/internal/app/model"
	"github.com/s-cod/rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

var (
	databaseURL string = "host=localhost dbname=restapi_test user=restapi password=restapi sslmode=disable"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("users")

	s := store.New(db)

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUserRepository_FingByEmail(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)

	defer teardown("users")

	s := store.New(db)
	u := model.TestUser(t)

	_, err := s.User().FindByEmail(u.Email)
	assert.Error(t, err)

	s.User().Create(u)
	user, err := s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
