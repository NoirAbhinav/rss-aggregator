package handler

import (
	"fmt"
	"net/http"

	models "github.com/NoirAbhinav/rss-aggregator/internal/db_handlers/models"
)

func (db *ApiConfig) UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("UserCreateHandler called")
	user_obj := models.User{
		Name: "Abhinav Nair",
	}
	user_obj.Create(db.DBPointer)
	respondWithError(w, 200, "User created")
}
