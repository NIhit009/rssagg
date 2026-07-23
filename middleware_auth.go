package main

import (
	"fmt"
	"net/http"

	"github.com/NIhit009/rssagg/internal/auth"
	"github.com/NIhit009/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			errResponse(w, 403, fmt.Sprintf("some error was found while logging: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKEY(r.Context(), apiKey)
		if err != nil {
			errResponse(w, 400, fmt.Sprintf("Error while retrieving the user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
