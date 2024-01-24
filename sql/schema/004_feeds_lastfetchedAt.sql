-- +goose Up
                          -- [2(1)]
ALTER TABLE feeds ADD COLUMN lastfetched_at timestamp;

-- +goose Down
                        --[2(2)]
Alter Table feeds DROP Column lastfetched_at;
      