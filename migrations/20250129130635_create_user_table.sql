-- +goose Up
-- +goose StatementBegin
create table if not exists users (
    id serial primary key,
    email varchar(60) unique,
    phone varchar(20) unique,
    password text not null,
    file_id int,
    file_url text,
    file_thumbnail_url text,
    bank_account_name varchar(80),
    bank_account_holder varchar(80),
    bank_account_number varchar(80),
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);

CREATE INDEX IF NOT EXISTS users_email_idx ON users USING HASH (email);
CREATE INDEX IF NOT EXISTS users_phone_idx ON users USING HASH (phone);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
