-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
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
SELECT 'down SQL query';
drop table if exists user;
-- +goose StatementEnd
