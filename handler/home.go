package handler

import (
	"context"
	"github.com/kutaui/htmx-go/view/home"
	"net/http"
)

type HomeHandler struct{}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// You can create a context if needed
	ctx := context.Background()

	// Pass the context and the response writer to the Render function
	err := home.Home().Render(ctx, w)
	if err != nil {
		// Handle the error if needed
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
