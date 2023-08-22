CREATE TABLE trades (
    id SERIAL PRIMARY KEY,
    asset TEXT,
    token TEXT,
    price DECIMAL(10, 2),
    maker TEXT,
    taker TEXT,
    time TIMESTAMP
);


INSERT INTO trades (asset, token, price, maker, taker, time) VALUES ('NASDAQ: AMZN', 'AMZN', 134.51, 'Robert', 'Alissa', '2023-08-22 12:40:55');
