CREATE TABLE IF NOT EXISTS t_accounts
(
   id   uuid NOT NULL PRIMARY KEY,
   type int
);

CREATE TABLE IF NOT EXISTS t_user_accounts
(
    user_id     uuid NOT NULL,
    account_id  uuid NOT NULL
)

CREATE TABLE IF NOT EXISTS t_transactions
(
    id              uuid NOT NULL PRIMARY KEY,
    description     text
)

CREATE TABLE IF NOT EXISTS t_postings
(
    id              serial PRIMARY KEY,
    type            int,
    transaction_id  uuid,
    account_id      uuid,
    amount          int
)