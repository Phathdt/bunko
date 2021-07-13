-- +goose Up
-- +goose StatementBegin
CREATE TABLE comments(
	id serial PRIMARY KEY,
	content TEXT,
	post_id INT,
	votes INT,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	CONSTRAINT fk_comments_post_id FOREIGN KEY(post_id) REFERENCES posts(id)
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE comments
-- +goose StatementEnd
