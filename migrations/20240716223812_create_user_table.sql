-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists users (
    id SERIAL not null PRIMARY KEY,
    name VARCHAR(255) not null,
    birthday date not null,
    email varchar(255),
    tg_id INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
