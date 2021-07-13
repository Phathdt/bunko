package bunko

import "time"

type Thread struct {
	ID          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"-"`
	UpdatedAt   time.Time `db:"updated_at" json:"-"`
}

type Post struct {
	ID        int       `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Content   string    `db:"content" json:"content"`
	ThreadId  int       `db:"thread_id" json:"thread_id"`
	Votes     int       `db:"votes" json:"votes"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}

type Comment struct {
	ID        int       `db:"id" json:"id"`
	Content   string    `db:"content" json:"content"`
	Votes     int       `db:"votes" json:"votes"`
	PostId    int       `db:"post_id" json:"post_id"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}

type ThreadStore interface {
	Thread(id int) (Thread, error)
	Threads() ([]Thread, error)
	CreateThread(t *Thread) error
	UpdateThread(t *Thread) error
	DeleteThread(id int) error
}

type PostStore interface {
	Post(id int) (Post, error)
	PostsByThread(threadID int) ([]Post, error)
	CreatePost(p *Post) error
	UpdatePost(p *Post) error
	DeletePost(id int) error
}

type CommentStore interface {
	Comment(id int) (Comment, error)
	CommentsByPost(postID int) ([]Comment, error)
	CreateComment(c *Comment) error
	UpdateComment(c *Comment) error
	DeleteComment(id int) error
}

type Store interface {
	ThreadStore
	PostStore
	CommentStore
}
