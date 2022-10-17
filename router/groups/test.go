package groups

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func TestDefault(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test default"))
}

func TestGroup() http.Handler {
	testRouter := chi.NewRouter()
	testRouter.Group(func(testRouter chi.Router) {
		testRouter.Get("/", TestDefault)
	})
	return testRouter
}
