-- +goose Up
                          -- [2(1)]
Create Table posts (      
    id TEXT PRIMARY KEY,
    created_at timestamp not null,
    updated_at timestamp not null,
    title text not null,
    description text,
    published_at timestamp not null,
    url text not null unique,
    feed_id text not null references feeds(id) on delete cascade
);

-- +goose Down
                        --[2(2)]
DROP TABLE posts;    