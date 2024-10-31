package api

import (
	"context"
	"github.com/DjordjeVuckovic/lets-go/apps/identity-server/internal/db"
)

func RegisterUser(ctx context.Context, database *db.Database) {
	query := "INSERT INTO users (name, age) VALUES ($1, $2)"
	_, err := database.DB.ExecContext(ctx, query, "Djordje", 25)
	if err != nil {
		return
	}

}
