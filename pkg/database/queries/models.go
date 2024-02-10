// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package queries

import (
	"database/sql"
)

type Address struct {
	ID          int64
	Address     int64
	Description sql.NullString
}

type Prefix struct {
	ID          int64
	Bitmask     int64
	Size        int64
	Description sql.NullString
	VlanID      sql.NullInt64
}

type User struct {
	ID           int64
	Username     string
	PasswordHash []byte
}

type Vlan struct {
	ID   int64
	Tag  int64
	Name sql.NullString
}
