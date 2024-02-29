package api

import (
	"net/http"

	"crypto/sha256"
	"encoding/json"
	"fmt"

	handler "github.com/NoirAbhinav/rss-aggregator/internal/handler"
	"gorm.io/gorm"

	auth "github.com/NoirAbhinav/rss-aggregator/internal/auth"
	models "github.com/NoirAbhinav/rss-aggregator/internal/db_handlers/models"
	"github.com/google/uuid"
)

type ApiConfig struct {
	DBPointer *gorm.DB
}

type ApiMethod interface {
	UserCreateHandler()
	UserSelectHandler()
}

type ApiMiddleware interface {
	MiddlewareAuth() http.HandlerFunc
}

type authedhandler func(http.ResponseWriter, *http.Request, *models.User)

func (db *ApiConfig) MiddlewareAuth(authhandler authedhandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			handler.RespondWithError(w, 401, "Unauthorized")
			return
		}
		user_obj := models.User{
			Apikey: apikey,
		}
		user, err := user_obj.Select(db.DBPointer)
		if err != nil {
			handler.RespondWithError(w, 404, "User does not exist")
			return
		}
		authhandler(w, r, user)
	}
}

func (db *ApiConfig) UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	format := struct {
		Name string `json:"name"`
	}{}
	fmt.Print("UserCreateHandler called")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&format)
	if err != nil {
		handler.RespondWithJSON(w, 400, "Invalid request: wrong params sent")
	}
	user_obj := models.User{
		ID:     uuid.New(),
		Name:   format.Name,
		Apikey: fmt.Sprintf("%x", sha256.Sum256([]byte(format.Name))),
	}
	user_obj.Create(db.DBPointer)
	handler.RespondWithJSON(w, 201, "User created")
}

func (db *ApiConfig) UserSelectHandler(w http.ResponseWriter, r *http.Request, user_obj *models.User) {
	user, err := user_obj.Select(db.DBPointer)
	if err != nil {
		handler.RespondWithError(w, 404, "User does not exist")
		return
	}
	handler.RespondWithJSON(w, 200, user)
}

func (db *ApiConfig) UserFeedCreateHandler(w http.ResponseWriter, r *http.Request, user *models.User) {
	defer r.Body.Close()
	format := struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}{}
	fmt.Print("UserCreateHandler called")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&format)
	if err != nil {
		handler.RespondWithJSON(w, 400, "Invalid request: wrong params sent")
	}
	feed_obj := models.UserFeed{
		ID:        uuid.New(),
		Name:      format.Name,
		Url:       format.Url,
		UserRefer: user.ID,
	}
	feed_obj.Create(db.DBPointer)
	handler.RespondWithJSON(w, 201, "Feed created")
}
