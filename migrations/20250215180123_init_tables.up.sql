-- Таблица пользователей
CREATE TABLE users
(
    -- Идентификатор пользователя
    id                SERIAL PRIMARY KEY,
    -- Логин пользователя, сделал ограничение в 16 символов
    username          VARCHAR(16) UNIQUE NOT NULL,
    -- Хэш пароля (SHA-256 в HEX, ровно 64 символа)
    password_hash     CHAR(64)           NOT NULL,
    -- Дата регистрации пользователя
    registration_date TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Таблица кошельков пользователей
CREATE TABLE wallets
(
    -- Идентификатор кошелька совпадает с id пользователя
    user_id INT PRIMARY KEY,
    -- Баланс пользователя, изначально баланс пользователя 1000, также делается проверка, чтобы пользователь не ушел в минус
    balance INTEGER NOT NULL DEFAULT 1000 CHECK (balance >= 0),
    CONSTRAINT fk_wallet_user
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE
);

-- Таблица товаров (продуктов)
CREATE TABLE products
(
    -- Идентификатор товара
    id    SERIAL PRIMARY KEY,
    -- Название товара
    name  VARCHAR(255) UNIQUE NOT NULL,
    -- Цена товара, проверка чтобы цена была больше 0
    price INTEGER             NOT NULL CHECK (price > 0)
);

-- Таблица покупок пользователей
CREATE TABLE user_purchases
(
    -- Идентификатор покупки
    id            SERIAL PRIMARY KEY,
    -- Ссылка на пользователя, который совершил покупку
    user_id       INT       NOT NULL,
    -- Ссылка на товар, используется id товара
    product_id    INT       NOT NULL,
    -- Количество купленных товаров, проверка больше 0
    quantity      INT       NOT NULL CHECK (quantity > 0),
    -- Время покупки
    purchase_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_purchase_user
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE,
    CONSTRAINT fk_purchase_product
        FOREIGN KEY (product_id)
            REFERENCES products (id)
            ON DELETE CASCADE
);

-- Таблица истории переводов/пополнений
CREATE TABLE coin_transactions
(
    -- Идентификатор транзакции
    id               SERIAL PRIMARY KEY,
    -- Идентификатор отправителя монет. NULL, если это пополнение
    from_user_id     INT,
    -- Идентификатор получателя монет
    to_user_id       INT       NOT NULL,
    -- Сумма перевода, проверка больше 0
    amount           INTEGER   NOT NULL CHECK (amount > 0),
    -- Время транзакции
    transaction_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_transaction_from
        FOREIGN KEY (from_user_id)
            REFERENCES users (id)
            ON DELETE SET NULL,
    CONSTRAINT fk_transaction_to
        FOREIGN KEY (to_user_id)
            REFERENCES users (id)
            ON DELETE CASCADE
);

-- Индексы для таблицы users
CREATE INDEX idx_users_username ON users (username);

-- Индексы для таблицы wallets
CREATE INDEX idx_wallets_user_id ON wallets (user_id);

-- Индексы для таблицы products
CREATE INDEX idx_products_name ON products (name);

-- Индексы для таблицы user_purchases
CREATE INDEX idx_user_purchases_user_id ON user_purchases (user_id);
CREATE INDEX idx_user_purchases_product_id ON user_purchases (product_id); -- Индекс на id товара
CREATE INDEX idx_user_purchases_purchase_date ON user_purchases (purchase_date);

-- Индексы для таблицы coin_transactions
CREATE INDEX idx_coin_transactions_from_user_id ON coin_transactions (from_user_id);
CREATE INDEX idx_coin_transactions_to_user_id ON coin_transactions (to_user_id);
CREATE INDEX idx_coin_transactions_transaction_date ON coin_transactions (transaction_date);

INSERT INTO products (name, price)
VALUES ('t-shirt', 80),
       ('cup', 20),
       ('book', 50),
       ('pen', 10),
       ('powerbank', 200),
       ('hoody', 300),
       ('umbrella', 200),
       ('socks', 10),
       ('wallet', 50),
       ('pink-hoody', 500);
