package store_test

import (
	"testing"

	"github.com/s-cod/rest-api/internal/app/model"
	"github.com/s-cod/rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, "host=localhost dbname=restapi_test user=restapi password=restapi sslmode=disable")

	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email:             "user@example.org",
		EncryptedPassword: "test",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
