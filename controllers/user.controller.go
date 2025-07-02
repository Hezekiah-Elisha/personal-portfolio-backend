package controllers

import (
	"context"

	db "github.com/hezekiah-elisha/personal-portfolio-backend/db/query/sqlc"
)

type UserController struct {
	db  *db.Queries
	ctx context.Context
}
