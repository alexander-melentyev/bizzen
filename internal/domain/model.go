package domain

import (
	"database/sql"
	"time"
)

// Model a basic Golang struct which includes the following fields:
//  ID
//  Creator
//  CreatedAt
//  Updater
//  UpdatedAt
//  Deleter
//  DeletedAt
// It may be embedded into your model or you may build your own model without it
//
// Example:
//  type Users struct {
//      Model
//      Email string `db:"email" json:"email"`
//  }
type Model struct {
	ID        uint64         `db:"id" json:"id"`                // E.g., 1.
	Creator   string         `db:"creator" json:"creator"`      // E.g., "Alexander Melentyev".
	Updater   string         `db:"updater" json:"updater"`      // E.g., "Alexander Melentyev".
	CreatedAt time.Time      `db:"created_at" json:"createdAt"` // E.g., "2021-02-27 00:00:00".
	UpdatedAt time.Time      `db:"updated_at" json:"updatedAt"` // E.g., "2021-02-27 00:00:00".
	Deleter   sql.NullString `db:"deleter" json:"deleter"`      // E.g., nil, "Alexander Melentyev".
	DeletedAt sql.NullTime   `db:"deleted_at" json:"deletedAt"` // E.g., nil, "2021-02-27 00:00:00".
}
