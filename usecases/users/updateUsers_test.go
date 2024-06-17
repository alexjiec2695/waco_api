package users_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"pokeapi/entities"
	"pokeapi/usecases/mock"
	"pokeapi/usecases/users"
	"testing"
)

func TestUpdateUsers_UpdateUsers_Success(t *testing.T) {
	storage := new(mock.ExecutorStorageMock)
	ctx := context.Background()
	user := entities.Users{}

	storage.On("UpdateUsers", ctx, user).Return(nil)
	update := users.NewUpdateUsers(storage)

	err := update.UpdateUsers(ctx, user)

	assert.NoError(t, err)
}

func TestUpdateUsers_UpdateUsers_Success_With_Error(t *testing.T) {
	storage := new(mock.ExecutorStorageMock)
	ctx := context.Background()
	user := entities.Users{}

	storage.On("UpdateUsers", ctx, user).Return(errors.New("error"))
	update := users.NewUpdateUsers(storage)

	err := update.UpdateUsers(ctx, user)

	assert.Error(t, err)
}
