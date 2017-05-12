package controllers

import (
	"github.com/jaax2707/ToDoGorm/access"
	"github.com/robfig/go-cache"
)

type DbController struct {
	cache *cache.Cache
	access *access.Db
}

func NewDbController(access *access.Db, cache *cache.Cache) *DbController {
	return &DbController{access:access, cache:cache}
}
