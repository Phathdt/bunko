package entities

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
