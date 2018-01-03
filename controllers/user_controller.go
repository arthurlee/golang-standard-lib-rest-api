package controllers

import (
	"../utils/caching"
	"database/sql"
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
