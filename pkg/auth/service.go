package auth

import (
	"net/http"

	"gopkg.in/oauth2.v3/server"
)

/*
ValidateToken checks if user token is valid and authorises user to access route
*/
func ValidateToken(f http.HandlerFunc, srv *server.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := srv.ValidationBearerToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		f.ServeHTTP(w, r)
	})
}
