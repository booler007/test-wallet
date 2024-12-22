CREATE TABLE IF NOT EXISTS wallets(
    uuid UUID PRIMARY KEY,
    balance numeric(10, 2) NOT NULL
);

CREATE TYPE operation_type AS ENUM ('DEPOSIT', 'WITHDRAW');
CREATE TABLE IF NOT EXISTS transactions (
    uuid UUID PRIMARY KEY,
    transaction_time timestamp,
    wallet UUID references wallets (uuid),
    operation operation_type,
    amount numeric(10, 2)
);
