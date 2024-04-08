package helpers

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func GenerateSlug(ctx context.Context, db *sql.DB, tableName string, title string) string{
	slug := strings.ToLower(strings.ReplaceAll(title, " ", "-"))

	query := fmt.Sprintf("SELECT slug FROM %s WHERE slug LIKE ? ORDER BY slug DESC", tableName)
	rows, err := db.QueryContext(ctx, query, slug+"%")
	PanicIfError(err)
	defer rows.Close()

	existingSlugs := make(map[string]bool)

	for rows.Next() {
		var existingSlug string
		err := rows.Scan(&existingSlug)
		PanicIfError(err)
		existingSlugs[existingSlug] = true
	}

	counter := 1
	newSlug := slug
	for existingSlugs[newSlug] {
		newSlug = fmt.Sprintf("%s-%d", slug, counter)
		counter++
	}

	return newSlug
}