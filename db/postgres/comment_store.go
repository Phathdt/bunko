package postgres

import (
	"bunko"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewCommentStore(db *sqlx.DB) *CommentStore {
	return &CommentStore{
		db,
	}
}

type CommentStore struct {
	*sqlx.DB
}

func (s CommentStore) Comment(id int) (bunko.Comment, error) {
	var c bunko.Comment

	if err := s.Get(&c, `SELECT * FROM comments WHERE id = $1`, id); err != nil {
		return bunko.Comment{}, fmt.Errorf("error getting comment: %w", err)
	}

	return c, nil
}

func (s CommentStore) CommentsByPost(postID int) ([]bunko.Comment, error) {
	var c []bunko.Comment

	if err := s.Select(&c, `SELECT * FROM comments WHERE post_id= $1`, postID); err != nil {
		return []bunko.Comment{}, fmt.Errorf("error getting comments: %w", err)
	}

	return c, nil
}

func (s CommentStore) CreateComment(c *bunko.Comment) error {
	if err := s.Get(c, "INSERT INTO  comments(content, post_id) VALUES ($1, $2) RETURNING *", c.Content, c.PostId); err != nil {
		return fmt.Errorf("error create comment: %w", err)
	}

	return nil
}

func (s CommentStore) UpdateComment(c *bunko.Comment) error {
	if err := s.Get(c, "UPDATE comments SET content = $1, votes = $2 WHERE id = $3", c.Content, c.Votes, c.ID); err != nil {
		return fmt.Errorf("error update comment: %w", err)
	}

	return nil
}

func (s CommentStore) DeleteComment(id int) error {
	if _, err := s.Exec("DELETE FROM  comments WHERE  id = $1", id); err != nil {
		return fmt.Errorf("error delete comment: %w", err)
	}

	return nil
}
