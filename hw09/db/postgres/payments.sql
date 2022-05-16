-- initialize payments tables
CREATE TABLE IF NOT EXISTS t_accounts
(
   id uuid NOT NULL PRIMARY KEY,
   owner_id uuid NOT NULL
);

CREATE TABLE IF NOT EXISTS t_payments
(
    order_id uuid NOT NULL UNIQUE,
    amount int
);

INSERT INTO t_accounts (id, owner_id) VALUES ('3bbc3d90-001b-4e0f-82c4-13116ca31c59', '87389739-a7b3-4e71-abd7-585e4a7a8d26')