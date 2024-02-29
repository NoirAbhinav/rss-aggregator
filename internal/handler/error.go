package handler

import (
	"fmt"
	"net/http"
)

func ErrHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("ErrHandler called")
	RespondWithError(w, 400, "Something went wrong ")
}
