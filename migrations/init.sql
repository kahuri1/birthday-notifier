CREATE TABLE user (
    id SERIAL not null PRIMARY KEY,
    name VARCHAR(255) not null,
    birthday date not null,
    email varchar(255),
    tg_id INT,
);

CREATE TABLE subscription (
    subscription_id INT not null,
    user_id INT not null,
    PRIMARY KEY (subscription_id, user_id),
    #указать ключи
);

CREATE TABLE #для работы уведомления


