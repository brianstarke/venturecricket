package main

import (
	"github.com/brianstarke/venturecricket/domain"
	"github.com/brianstarke/venturecricket/users"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(domain.DomainMiddleware())

	// root route, replace with public app later
	m.Get("/", func() string {
		return "WORK IN PROGRESS"
	})

	// user routes
	m.Get("/users", users.ListUsers)
	m.Get("/users/:id", users.GetUser)
	m.Post("/users", users.CreateUser)

	m.Run()
}
