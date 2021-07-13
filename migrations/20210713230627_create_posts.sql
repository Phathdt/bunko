-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts(
	id serial PRIMARY KEY,
	title TEXT,
	content TEXT,
	thread_id INT,
	votes INT DEFAULT 0,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	CONSTRAINT fk_posts_thread_id FOREIGN KEY(thread_id) REFERENCES threads(id)
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts
-- +goose StatementEnd
