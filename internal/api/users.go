package api

import (
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	name string
	age int
}

func UsersRoute() {
	http.HandleFunc("/api/users/", func(w http.ResponseWriter, r *http.Request) {
		userId := strings.TrimPrefix(r.URL.Path, "/api/users/")

		if len(userId) == 0 {
			var users map[string]User

			users = make(map[string]User)
			users["qwe123"] = User{
				"Victor", 27,
			}
			users["asd456"] = User{
				"Thales", 28,
			}

			fmt.Printf("Users map: %q", users)
			fmt.Fprintf(w, "Users route is working on path: %s \n Users map: %q", r.URL.Path, users)

			return
		}
		fmt.Printf("Single user route is working. Current user ID: %s",userId)
		fmt.Fprintf(w, "Single user route is working. Current user ID: %s", userId)
	})
}