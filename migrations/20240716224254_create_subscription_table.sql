-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE if not exists subscription (
    subscription_id INT not null,
    user_id INT not null,
    PRIMARY KEY (subscription_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists subscription;
-- +goose StatementEnd
