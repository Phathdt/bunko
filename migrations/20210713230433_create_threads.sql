-- +goose Up
-- +goose StatementBegin
CREATE TABLE threads(
	id serial PRIMARY KEY,
	title TEXT,
	description TEXT,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE threads
-- +goose StatementEnd
