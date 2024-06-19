package sqlstore_test

import (
	"testing"

	"github.com/s-cod/rest-api/internal/app/model"
	"github.com/s-cod/rest-api/internal/app/store"
	"github.com/s-cod/rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

var (
	databaseURL string = "host=localhost dbname=restapi_test user=restapi password=restapi sslmode=disable"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)

}

func TestUserRepository_FingByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)

	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	_, err := s.User().FindByEmail(u.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u)
	user, err := s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
