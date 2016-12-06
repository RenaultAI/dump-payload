
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE providers (
  id uuid NOT NULL primary key,
  created_at timestamp with time zone NOT NULL,
  updated_at timestamp with time zone NOT NULL,
  friendly_name character varying(255) NOT NULL,
  nickname character varying(255) NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE providers;

