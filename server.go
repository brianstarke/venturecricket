package main

import (
	"github.com/brianstarke/venturecricket/cards"
	"github.com/brianstarke/venturecricket/domain"
	"github.com/brianstarke/venturecricket/users"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{IndentJSON: true}))
	m.Use(domain.DomainMiddleware())

	// authentication routes
	m.Post("/api/v1/auth/login", users.AuthenticateUser)

	// user routes
	m.Get("/api/v1/users", users.ListUsers)
	m.Get("/api/v1/users/:id", users.GetUser)
	m.Post("/api/v1/users", users.CreateUser)

	// card routes
	m.Get("/api/v1/cards/get/:gameId", cards.GetCards)
	m.Get("/api/v1/cards/create/:gameId", cards.CreateDeck)

	m.Run()
}
