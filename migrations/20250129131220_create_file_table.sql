-- +goose Up
-- +goose StatementBegin
create table if not exists file (
    id serial primary key,
    user_id int,
    file_url text,
    file_thumbnail_url text,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);

CREATE INDEX IF NOT EXISTS file_user_id_idx ON file USING HASH (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists file;
-- +goose StatementEnd
