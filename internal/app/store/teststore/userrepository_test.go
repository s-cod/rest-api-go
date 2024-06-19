package teststore_test

import (
	"testing"

	"github.com/s-cod/rest-api/internal/app/model"
	"github.com/s-cod/rest-api/internal/app/store"
	"github.com/s-cod/rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {

	s := teststore.New()

	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)

}

func TestUserRepository_FingByEmail(t *testing.T) {

	s := teststore.New()

	u := model.TestUser(t)

	_, err := s.User().FindByEmail(u.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u)
	user, err := s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
