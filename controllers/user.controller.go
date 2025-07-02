package controllers

import (
	"context"

	db "personal-portfolio-backend/db/query/sqlc"
)

type UserController struct {
	db  *db.Queries
	ctx context.Context
}
