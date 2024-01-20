-- +goose Up
                          -- [2(1)]
Create Table users(      
    id TEXT PRIMARY KEY,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null

);

-- +goose Down
                        --[2(2)]
DROP TABLE users;       