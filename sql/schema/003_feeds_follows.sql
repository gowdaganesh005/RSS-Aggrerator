-- +goose Up
                          -- [2(1)]
Create Table feeds_follows (      
    id TEXT PRIMARY KEY,
    created_at timestamp not null,
    updated_at timestamp not null,
    
    user_id  text not null references users(id) ON  DELETE CASCADE,
    feed_id  text not null references feeds(id) ON  DELETE CASCADE

);

-- +goose Down
                        --[2(2)]
DROP TABLE feeds_follows;       