CREATE TABLE IF NOT EXISTS wallets(
    uuid UUID PRIMARY KEY,
    balance numeric(10, 2) NOT NULL
);

CREATE TYPE operation AS ENUM ("DEPOSIT", "WITHDRAW");
CREATE TABLE IF NOT EXISTS transactions (
    transaction_time timestamp,
    wallet UUID,
    operation_type operation,
);