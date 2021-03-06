package postgres

import (
	"bunko/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewPostStore(db *sqlx.DB) *PostStore {
	return &PostStore{
		db,
	}
}

type PostStore struct {
	*sqlx.DB
}

func (s PostStore) Post(id int) (entities.Post, error) {
	var p entities.Post

	if err := s.Get(&p, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return entities.Post{}, fmt.Errorf("error getting post: %w", err)
	}

	return p, nil
}

func (s PostStore) PostsByThread(threadID int) ([]entities.Post, error) {
	var p []entities.Post

	if err := s.Select(&p, `SELECT * FROM posts WHERE thread_id = $1`, threadID); err != nil {
		return []entities.Post{}, fmt.Errorf("error getting posts: %w", err)
	}

	return p, nil
}

func (s PostStore) CreatePost(p *entities.Post) error {
	if err := s.Get(p, "INSERT INTO  posts(title, content, thread_id) VALUES ($1, $2, $3) RETURNING *", p.Title, p.Content, p.ThreadId); err != nil {
		return fmt.Errorf("error create post: %w", err)
	}

	return nil
}

func (s PostStore) UpdatePost(p *entities.Post) error {
	if err := s.Get(p, "UPDATE posts SET title = $1, content = $2, votes = $3", p.Title, p.Content, p.Votes); err != nil {
		return fmt.Errorf("error update post: %w", err)
	}

	return nil
}

func (s PostStore) DeletePost(id int) error {
	if _, err := s.Exec("DELETE FROM  posts WHERE  id = $1", id); err != nil {
		return fmt.Errorf("error delete post: %w", err)
	}

	return nil
}
