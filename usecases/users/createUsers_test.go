package users_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"pokeapi/entities"
	ucMock "pokeapi/usecases/mock"
	"pokeapi/usecases/users"
	"testing"
)

func TestCreateUsers_CreateUsers_Successful(t *testing.T) {
	store := new(ucMock.ExecutorStorageMock)
	ctx := context.Background()
	user := entities.Users{}

	store.On("CreateUsers", ctx, mock.Anything).Return(nil)

	uc := users.NewCreateUsers(store)

	err := uc.CreateUsers(ctx, user)

	assert.NoError(t, err)
}
