-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists subscriptions (
    subscription_id INT not null,
    user_id INT not null,
    PRIMARY KEY (subscription_id, user_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (subscription_id) REFERENCES notification_subscriptions(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists subscriptions;
-- +goose StatementEnd
