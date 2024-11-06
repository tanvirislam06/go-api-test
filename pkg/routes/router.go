package routes

import (
	"fmt"
	"net/http"

	"github.com/tanvirislam06/go-api-test/internal/user"
)

func SetupRouter(userHandler *user.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	// Root handler
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the User API")
	})
	mux.HandleFunc("/user", userHandler.GetUser)
	return mux
}
