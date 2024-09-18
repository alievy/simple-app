package components

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// NavBar represents the navigation bar component
type NavBar struct {
	app.Compo
}

// Render renders the NavBar component
func (n *NavBar) Render() app.UI {
	return app.Nav().Class("navbar").Body(
		app.Ul().Body(
			app.Li().Body(
				app.A().Href("/").Text("Welcome"),
			),
			app.Li().Body(
				app.A().Href("/api-calls").Text("API Calls"),
			),
		),
	)
}
