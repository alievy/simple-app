package components

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// FileManager represents the API Calls view
type FileManager struct {
	app.Compo
	HealthStatus string
	loading      bool // Indicator to show if API call is in progress
}

// OnMount is triggered when the component is mounted
func (f *FileManager) OnMount(ctx app.Context) {
	f.fetchHealthStatusAsync(ctx)
}

// fetchHealthStatusAsync makes the API call asynchronously using ctx.Async
func (f *FileManager) fetchHealthStatusAsync(ctx app.Context) {
	f.loading = true
	ctx.Dispatch(func(ctx app.Context) {})
	ctx.Async(func() {
		resp, err := http.Get("http://localhost:8080/v1/health")
		if err != nil {
			ctx.Dispatch(func(ctx app.Context) {
				f.HealthStatus = "Error fetching API"
				f.loading = false
			})
			return
		}
		defer resp.Body.Close()
		var result map[string]string
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			ctx.Dispatch(func(ctx app.Context) {
				f.HealthStatus = "Error decoding response"
				f.loading = false
			})
			return
		}
		ctx.Dispatch(func(ctx app.Context) {
			f.HealthStatus = result["status"]
			f.loading = false
		})
	})
}

// Render renders the FileManager component UI
func (f *FileManager) Render() app.UI {
	return app.Div().Body(
		&NavBar{}, // Include the NavBar here
		app.H1().Text("API Calls View"),
		app.If(f.loading,
			func() app.UI {
				return app.Div().Text("Loading...")
			},
		).Else(
			func() app.UI {
				return app.Div().Text(fmt.Sprintf("API Health Status: %s", f.HealthStatus))
			},
		),
		app.Button().
			Text("Refresh API Call").
			OnClick(func(ctx app.Context, e app.Event) {
				f.fetchHealthStatusAsync(ctx)
			}),
	)
}
