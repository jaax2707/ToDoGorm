package controllers

import "github.com/jaax2707/ToDoGorm/access"

type DbController struct {
	access *access.Db
}

func NewDbController(access *access.Db) *DbController {
	return &DbController{access}
}
