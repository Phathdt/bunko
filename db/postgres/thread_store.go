package postgres

import (
	"bunko/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewThreadStore(db *sqlx.DB) *ThreadStore {
	return &ThreadStore{
		db,
	}
}

type ThreadStore struct {
	*sqlx.DB
}

func (s ThreadStore) Thread(id int) (entities.Thread, error) {
	var t entities.Thread

	if err := s.Get(&t, `SELECT * FROM threads WHERE id = $1`, id); err != nil {
		return entities.Thread{}, fmt.Errorf("error getting thread: %w", err)
	}

	return t, nil
}

func (s ThreadStore) Threads() ([]entities.Thread, error) {
	var t []entities.Thread

	if err := s.Select(&t, `SELECT * FROM threads`); err != nil {
		return []entities.Thread{}, fmt.Errorf("error getting threads: %w", err)
	}

	return t, nil
}

func (s ThreadStore) CreateThread(t *entities.Thread) error {
	if err := s.Get(t, "INSERT INTO  threads(title, description) VALUES ($1, $2) RETURNING *", t.Title, t.Description); err != nil {
		return fmt.Errorf("error create thread: %w", err)
	}

	return nil
}

func (s ThreadStore) UpdateThread(t *entities.Thread) error {
	if err := s.Get(t, "UPDATE threads SET title = $1, description = $2 WHERE  id = $3", t.Title, t.Description, t.ID); err != nil {
		return fmt.Errorf("error update thread: %w", err)
	}

	return nil
}

func (s ThreadStore) DeleteThread(id int) error {
	if _, err := s.Exec("DELETE FROM  threads WHERE  id = $1", id); err != nil {
		return fmt.Errorf("error delete thread: %w", err)
	}

	return nil
}
