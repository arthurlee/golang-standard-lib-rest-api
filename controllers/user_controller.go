package controllers

import (
	"../utils/caching"
	"database/sql"
	"net/http"
	"encoding/json"
)

type UserController struct {
	DB    *sql.DB
	Cache caching.Cache
}

func NewUserController(db *sql.DB, c caching.Cache) *UserController {
	return &UserController{
		DB:    db,
		Cache: c,
	}
}

//
// Register
//

func (jc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)

}