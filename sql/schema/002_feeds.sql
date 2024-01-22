-- +goose Up
                          -- [2(1)]
Create Table feeds (      
    id TEXT PRIMARY KEY,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null,
    url TEXT unique not null,
    user_id text references users(id) ON  DELETE CASCADE

);

-- +goose Down
                        --[2(2)]
DROP TABLE feeds;       