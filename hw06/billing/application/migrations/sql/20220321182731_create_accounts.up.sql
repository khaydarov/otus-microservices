BEGIN;
CREATE TABLE IF NOT EXISTS t_accounts (
    id uuid NOT NULL PRIMARY KEY,
    owner_id uuid NOT NULL,
    balance integer
);
CREATE UNIQUE INDEX idx_accounts_owner_unique ON t_accounts(id, owner_id);
COMMIT;