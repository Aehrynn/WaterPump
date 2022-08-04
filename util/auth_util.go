package util

import (
	"fmt"
	"net/http"
	"strings"
)

func Authorize(w http.ResponseWriter, r *http.Request, apiKey string) error {

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Not Authenticated")
		return fmt.Errorf(`not authenticated`)
	}

	reqToken = strings.TrimSpace(splitToken[1])

	if reqToken != apiKey {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Not Authenticated")
		return fmt.Errorf("authentication failed")
	}

	return nil
}
