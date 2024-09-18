package components

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// Welcome represents the welcome page
type Welcome struct {
	app.Compo
}

// Render renders the Welcome component
func (w *Welcome) Render() app.UI {
	return app.Div().Body(
		&NavBar{}, // Include the NavBar here
		app.H1().Text("Welcome to My App"),
		app.P().Text("This is the welcome screen. Use the navbar to navigate."),
	)
}
