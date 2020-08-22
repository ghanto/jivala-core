package user

import (
	"fmt"
	"net/http"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("/users")
	w.WriteHeader(http.StatusOK)
}
