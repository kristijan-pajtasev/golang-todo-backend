package helpers

import "net/http"

func GetCurrentUserKey(r *http.Request) string {
	userId := r.URL.Query().Get("user")
	return userId
}
