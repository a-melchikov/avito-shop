DROP TABLE IF EXISTS coin_transactions CASCADE;
DROP TABLE IF EXISTS user_purchases CASCADE;
DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS wallets CASCADE;
DROP TABLE IF EXISTS users CASCADE;

DROP INDEX IF EXISTS idx_users_username;
DROP INDEX IF EXISTS idx_wallets_user_id;
DROP INDEX IF EXISTS idx_products_name;
DROP INDEX IF EXISTS idx_user_purchases_user_id;
DROP INDEX IF EXISTS idx_user_purchases_product_id;
DROP INDEX IF EXISTS idx_user_purchases_purchase_date;
DROP INDEX IF EXISTS idx_coin_transactions_from_user_id;
DROP INDEX IF EXISTS idx_coin_transactions_to_user_id;
DROP INDEX IF EXISTS idx_coin_transactions_transaction_date;
