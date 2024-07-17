-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE notification_subscription (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    notification_method VARCHAR(50) NOT NULL,
    --created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    --updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists notification_subscription;
-- +goose StatementEnd
