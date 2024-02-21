package main

import (
	"net/http"

	models "github.com/NoirAbhinav/rss-aggregator/internal/db_handlers/models"
	utility "github.com/NoirAbhinav/rss-aggregator/internal/db_handlers/utils"
)

func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	db, err := utility.GetDB()
	if err != nil {
		respondWithError(w, 500, "Database connection failed")
		return
	}
	user_obj := models.User{
		Name: "Abhinav Nair",
	}
	user_obj.Create(db)
	respondWithError(w, 400, "Something went wrong ")
}
