package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"pokeapi/entities"
	"pokeapi/infrastructure/rest"
	"pokeapi/usecases/login"
	"pokeapi/usecases/pokemon"
	"pokeapi/usecases/users"
)

type Handler struct {
	app         *fiber.App
	createUsers users.CreateUsersExecutor
	updateUsers users.UpdateUsersExecutor
	login       login.LoginExecutor
	getUsers    users.GetUsersExecutor
	pokemos     pokemon.PokemonExecutor
}

func NewHandler(app *rest.Server,
	createUsers users.CreateUsersExecutor,
	login login.LoginExecutor,
	updateUsers users.UpdateUsersExecutor,
	getUsers users.GetUsersExecutor,
	pokemos pokemon.PokemonExecutor,
) Handler {
	return Handler{
		app:         app.App,
		createUsers: createUsers,
		updateUsers: updateUsers,
		login:       login,
		getUsers:    getUsers,
		pokemos:     pokemos,
	}
}

func (r *Handler) CreateUsers() {
	r.app.Post("Users", func(ctx *fiber.Ctx) error {
		users := entities.Users{}

		err := json.Unmarshal(ctx.Body(), &users)
		if err != nil {
			return err
		}

		return r.createUsers.CreateUsers(ctx.Context(), users)
	})
}

func (r *Handler) Login() {
	r.app.Post("Login", func(ctx *fiber.Ctx) error {
		users := entities.Users{}

		err := json.Unmarshal(ctx.Body(), &users)
		if err != nil {
			return err
		}

		id, err := r.login.Login(ctx.Context(), users)
		if err != nil {
			return ctx.SendStatus(http.StatusNoContent)
		}
		type response struct {
			ID string
		}

		b, _ := json.Marshal(response{
			ID: id,
		})

		return ctx.Status(http.StatusOK).Send(b)
	})
}

func (r *Handler) UpdateUsers() {
	r.app.Put("Users", func(ctx *fiber.Ctx) error {
		users := entities.Users{}

		err := json.Unmarshal(ctx.Body(), &users)
		if err != nil {
			return err
		}

		return r.updateUsers.UpdateUsers(ctx.Context(), users)
	})
}

func (r *Handler) GetUsers() {
	r.app.Get("User/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		user, err := r.getUsers.GetUsers(ctx.Context(), id)
		if err != nil {
			return err
		}
		b, err := json.Marshal(user)
		if err != nil {
			return err
		}
		return ctx.Status(200).Send(b)

	})
}

func (r *Handler) GetPokemons() {
	r.app.Get("pokemons", func(ctx *fiber.Ctx) error {
		p, err := r.pokemos.GetPokemons(ctx.Context())
		if err != nil {
			return err
		}
		b, err := json.Marshal(p)
		if err != nil {
			return err
		}
		return ctx.Status(200).Send(b)

	})
}

func (r *Handler) GetPokemon() {
	r.app.Post("pokemon", func(ctx *fiber.Ctx) error {

		req := entities.Req{}

		err := json.Unmarshal(ctx.Body(), &req)
		if err != nil {
			return err
		}

		p, err := r.pokemos.GetPokemon(ctx.Context(), req.Url)
		if err != nil {
			return err
		}
		b, err := json.Marshal(p)
		if err != nil {
			return err
		}
		return ctx.Status(200).Send(b)

	})
}

func (r *Handler) Start() error {
	r.CreateUsers()
	r.UpdateUsers()
	r.GetUsers()
	r.GetPokemons()
	r.GetPokemon()
	r.Login()
	return r.app.Listen(":3000")
}
