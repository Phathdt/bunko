package entities

import "time"

type Thread struct {
	ID          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"-"`
	UpdatedAt   time.Time `db:"updated_at" json:"-"`
}
