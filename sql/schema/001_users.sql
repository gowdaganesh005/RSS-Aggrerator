-- +goose Up
Create Table users( 
    id TEXT PRIMARY KEY,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null

);

-- +goose Down
DROP TABLE users;