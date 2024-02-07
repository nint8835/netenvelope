package database

import (
	"database/sql"

	_ "modernc.org/sqlite"

	"github.com/nint8835/netenvelope/pkg/database/queries"
)

func New(path string) (*queries.Queries, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return queries.New(db), nil
}
