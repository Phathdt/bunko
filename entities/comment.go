package entities

import "time"

type Comment struct {
	ID        int       `db:"id" json:"id"`
	Content   string    `db:"content" json:"content"`
	Votes     int       `db:"votes" json:"votes"`
	PostId    int       `db:"post_id" json:"post_id"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}
