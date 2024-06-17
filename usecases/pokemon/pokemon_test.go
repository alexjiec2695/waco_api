package pokemon_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"pokeapi/entities"
	"pokeapi/usecases/mock"
	"pokeapi/usecases/pokemon"
	"testing"
)

func TestPokemon_GetPokemon_Success(t *testing.T) {
	restyMock := new(mock.ExecutorRestyPokemon)
	ctx := context.Background()
	expect := entities.PokemonDetail{}

	restyMock.On("GetPokemonDetail", "url").Return(expect, nil)

	uc := pokemon.NewPokemon(restyMock)

	pokemons, err := uc.GetPokemon(ctx, "url")

	assert.NoError(t, err)
	assert.NotNil(t, pokemons)
	assert.Equal(t, expect, pokemons)
}
