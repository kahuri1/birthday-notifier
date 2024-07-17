-- +goose Up
-- +goose StatementBegin
create type method_notification as enum ('email', 'telegram');

CREATE TABLE notification_subscriptions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    notification_method method_notification NOT NULL,
    --created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    --updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists notification_subscriptions;
DROP TYPE method_notification;
-- +goose StatementEnd
