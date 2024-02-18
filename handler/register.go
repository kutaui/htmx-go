package handler

import (
	"context"
	"github.com/kutaui/htmx-go/view/pages"
	"net/http"
)

type RegisterPageHandler struct{}

func (h *RegisterPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	err := pages.Register().Render(ctx, w)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
