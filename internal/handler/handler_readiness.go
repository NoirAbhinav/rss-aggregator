package handler

import (
	"fmt"
	"net/http"
)

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("ReadinessHandler called")
	respondWithJSON(w, 200, map[string]string{"status": "ok"})
}
