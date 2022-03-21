BEGIN;
CREATE TABLE IF NOT EXISTS t_users (
    id uuid UNIQUE PRIMARY KEY,
    login text NOT NULL,
    password text NOT NULL,
    firstname text NOT NULL,
    lastname text NOT NULL
);
COMMIT;
