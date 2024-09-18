package main

import (
	"log"
	"net/http"
	"simple-app/components"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func main() {
	// Define routes for components
	app.Route("/", func() app.Composer { return &components.Welcome{} })
	app.Route("/api-calls", func() app.Composer { return &components.FileManager{} })

	// Mount the NavBar component (as a layout or included globally)
	app.Route("/", func() app.Composer { return &components.NavBar{} })

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "My App",
		Description: "An example Go PWA with navigation",
		Styles: []string{
			"/web/styles.css",
		},
	})

	// Serve the application on localhost
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
