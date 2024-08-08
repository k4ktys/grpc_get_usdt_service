CREATE TABLE IF NOT EXISTS trades
(
    id SERIAL PRIMARY KEY,
    trade_timestamp TIMESTAMP,
    data JSON
);