package handler

import (
	"context"
	"github.com/kutaui/htmx-go/view/pages"
	"net/http"
)

type HomeHandler struct{}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	err := pages.Home().Render(ctx, w)
	if err != nil {
		// Handle the error if needed
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
