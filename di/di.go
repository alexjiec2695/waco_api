package di

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"pokeapi/infrastructure/handlers"
	"pokeapi/infrastructure/rest"
	resty2 "pokeapi/infrastructure/resty"
	"pokeapi/infrastructure/storage"
	"pokeapi/usecases/login"
	"pokeapi/usecases/pokemon"
	"pokeapi/usecases/users"
)

func Start() error {
	server := rest.NewServer()
	str := goDotEnvVariable("STR")
	persistence, err := storage.NewConnectionStorage(str)
	if err != nil {
		return err
	}
	createUsers := users.NewCreateUsers(persistence)
	login := login.NewLogin(persistence)
	updateUsers := users.NewUpdateUsers(persistence)
	getUsers := users.NewGetUsers(persistence)
	resty := resty2.NewPokemon()
	pokes := pokemon.NewPokemon(resty)

	routers := handlers.NewHandler(server, createUsers, login, updateUsers, getUsers, pokes)

	return routers.Start()
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}
