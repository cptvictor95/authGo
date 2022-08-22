package api

import (
	"fmt"
	"net/http"
	"strings"
)

func UsersRoute() {
	http.HandleFunc("/api/users/", func(w http.ResponseWriter, r *http.Request) {
		userId := strings.TrimPrefix(r.URL.Path, "/api/users/")

		if len(userId) == 0 {
			fmt.Fprintf(w, "Users route is working, %s", r.URL.Path)

			return
		}
		fmt.Printf("Single user route is working. Current user ID: %s",userId)
		fmt.Fprintf(w, "Single user route is working. Current user ID: %s", userId)
	})
}