package entities

import "time"

type Post struct {
	ID        int       `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Content   string    `db:"content" json:"content"`
	ThreadId  int       `db:"thread_id" json:"thread_id"`
	Votes     int       `db:"votes" json:"votes"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}
